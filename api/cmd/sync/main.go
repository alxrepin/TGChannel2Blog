package main

import (
	"app/internal/domain"
	"context"
	"log"

	"app/internal/application"
	"app/internal/application/usecase/sync"
	"app/internal/infrastructure/bus"
	"app/internal/infrastructure/postgres"
	"app/internal/infrastructure/telegram"

	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/jetstream"
	"github.com/nats-io/nats.go"
)

func main() {
	config, err := application.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	conn, err := nats.Connect(config.NATSURL)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer conn.Close()

	js, err := conn.JetStream()
	if err != nil {
		log.Fatalf("failed to create JetStream context: %v", err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     string(domain.RawMessageReceived),
		Subjects: []string{string(domain.RawMessageReceived)},
	})
	if err != nil && err.Error() != "stream name already in use" {
		log.Fatalf("failed to create stream: %v", err)
	}

	publisher, err := jetstream.NewPublisher(
		jetstream.PublisherConfig{
			URL: config.NATSURL,
		},
	)
	if err != nil {
		log.Fatalf("failed to create publisher: %v", err)
	}
	defer publisher.Close()

	client := telegram.NewClient(config.APIID, config.APIHash, config.Phone)
	repository := telegram.NewRawMessageRepository(client)
	bus := bus.NewWatermillBus(publisher)
	uc := sync.NewLoadHistoryRawMessagesUseCase(repository, bus)

	err = uc.Execute(context.Background(), config.ChannelUsername)
	if err != nil {
		log.Fatalf("failed to sync messages: %v", err)
	}

	channelRepo := postgres.NewChannelRepository(config.DB)
	channelInfo, err := repository.GetChannelInfo(context.Background(), config.ChannelUsername)
	if err != nil {
		log.Fatalf("failed to get channel info: %v", err)
	}

	err = channelRepo.CreateOrUpdate(context.Background(), channelInfo)
	if err != nil {
		log.Fatalf("failed to save channel info: %v", err)
	}
}
