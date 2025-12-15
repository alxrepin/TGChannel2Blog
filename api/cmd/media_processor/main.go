package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"app/internal/application"
	"app/internal/domain"
	"app/internal/infrastructure/minio"
	"app/internal/infrastructure/telegram"

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
		Name:     string(domain.MediaReceived),
		Subjects: []string{string(domain.MediaReceived)},
	})
	if err != nil && err.Error() != "stream name already in use" {
		log.Fatalf("failed to create media stream: %v", err)
	}

	subscriber, err := jetstream.NewSubscriber(jetstream.SubscriberConfig{
		URL:                 config.NATSURL,
		ResourceInitializer: jetstream.GroupedConsumer("media_processor"),
	})

	if err != nil {
		log.Fatalf("failed to create subscriber: %v", err)
	}
	defer subscriber.Close()

	minioClient, err := minio.NewClient(
		config.MinioURL,
		config.MinioAccessKey,
		config.MinioSecretKey,
		config.MinioBucket,
	)
	if err != nil {
		log.Fatalf("failed to create Minio client: %v", err)
	}

	telegramClient := telegram.NewClient(config.APIID, config.APIHash, config.Phone)
	telegramRepo := telegram.NewRawMessageRepository(telegramClient)

	messages, err := subscriber.Subscribe(context.Background(), string(domain.MediaReceived))
	if err != nil {
		log.Fatalf("failed to subscribe: %v", err)
	}

	log.Println("Media processor started, waiting for messages...")

	for msg := range messages {
		var rawMessage domain.RawMessage
		err := json.Unmarshal(msg.Payload, &rawMessage)
		if err != nil {
			log.Printf("failed to unmarshal message: %v", err)
			msg.Ack()
			continue
		}

		if rawMessage.Media == nil {
			log.Printf("no media in message %d", rawMessage.ID)
			msg.Ack()
			continue
		}

		data, err := telegramRepo.DownloadMedia(context.Background(), *rawMessage.Media)
		if err != nil {
			log.Printf("failed to download media for message %d: %v", rawMessage.ID, err)
			msg.Nack()
			continue
		}

		objectName := fmt.Sprintf("%d-%s", rawMessage.ID, rawMessage.Media.Type)
		contentType := getContentType(rawMessage.Media.Type)

		err = minioClient.Upload(objectName, data, contentType)
		if err != nil {
			log.Printf("failed to upload media for message %d: %v", rawMessage.ID, err)
			msg.Nack()
			continue
		}

		log.Printf("Successfully processed media for message %d", rawMessage.ID)
		msg.Ack()
	}
}

func getContentType(mediaType domain.MediaType) string {
	switch mediaType {
	case domain.MediaTypePhoto:
		return "image/jpeg"
	case domain.MediaTypeVideo:
		return "video/mp4"
	case domain.MediaTypeAudio:
		return "audio/mpeg"
	case domain.MediaTypeDocument:
		return "application/octet-stream"
	default:
		return "application/octet-stream"
	}
}
