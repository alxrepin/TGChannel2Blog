package main

import (
	"context"
	"encoding/json"
	"log"

	"app/internal/application"
	"app/internal/application/usecase/sync"
	"app/internal/domain"
	"app/internal/infrastructure/bus"
	"app/internal/infrastructure/postgres"

	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/jetstream"
	nc "github.com/nats-io/nats.go"
)

func main() {
	config, err := application.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	conn, err := nc.Connect(config.NATSURL)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer conn.Close()

	js, err := conn.JetStream()
	if err != nil {
		log.Fatalf("failed to create JetStream context: %v", err)
	}

	_, err = js.AddStream(&nc.StreamConfig{
		Name:     string(domain.RawMessageReceived),
		Subjects: []string{string(domain.RawMessageReceived)},
	})
	if err != nil && err.Error() != "stream name already in use" {
		log.Fatalf("failed to create stream: %v", err)
	}

	_, err = js.AddStream(&nc.StreamConfig{
		Name:     string(domain.MediaReceived),
		Subjects: []string{string(domain.MediaReceived)},
	})
	if err != nil && err.Error() != "stream name already in use" {
		log.Fatalf("failed to create media stream: %v", err)
	}

	subscriber, err := jetstream.NewSubscriber(jetstream.SubscriberConfig{
		URL:                 config.NATSURL,
		ResourceInitializer: jetstream.GroupedConsumer("subscriber"),
	})

	if err != nil {
		log.Fatalf("failed to create subscriber: %v", err)
	}
	defer subscriber.Close()

	publisher, err := jetstream.NewPublisher(
		jetstream.PublisherConfig{
			URL: config.NATSURL,
		},
	)
	if err != nil {
		log.Fatalf("failed to create publisher: %v", err)
	}
	defer publisher.Close()

	bus := bus.NewWatermillBus(publisher)

	repository := postgres.NewPostRepository(config.DB)
	uc := sync.NewSyncRawMessageUseCase(repository)

	messages, err := subscriber.Subscribe(context.Background(), string(domain.RawMessageReceived))
	if err != nil {
		log.Fatalf("failed to subscribe: %v", err)
	}

	log.Println("Subscriber started, waiting for messages...")

	for msg := range messages {
		var rawMessage domain.RawMessage
		err := json.Unmarshal(msg.Payload, &rawMessage)
		if err != nil {
			log.Printf("failed to unmarshal message: %v", err)
			msg.Ack()
			continue
		}

		if (rawMessage.Text == nil || *rawMessage.Text == "") && rawMessage.GroupID > 0 {
			data, err := json.Marshal(rawMessage)
			if err != nil {
				log.Printf("failed to marshal raw message for media: %v", err)
				msg.Ack()
				continue
			}
			err = bus.Dispatch(domain.MediaReceived, data)
			if err != nil {
				log.Printf("failed to dispatch media message: %v", err)
			}
			msg.Ack()
			continue
		}

		err = uc.Execute(context.Background(), rawMessage)
		if err != nil {
			log.Printf("failed to process message %d: %v", rawMessage.ID, err)
			msg.Nack()
		} else {
			log.Printf("Successfully processed message %d", rawMessage.ID)
			msg.Ack()
		}
	}
}
