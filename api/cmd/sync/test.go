package main

import (
	"app/internal/domain"
	"context"
	"log"

	"app/internal/application"
	"app/internal/infrastructure/minio"
	"app/internal/infrastructure/telegram"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

func main() {
	config, err := application.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

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

	client := telegram.NewClient(config.APIID, config.APIHash, config.Phone)
	minioClient, err := minio.NewClient(config.MinioURL, config.MinioAccessKey, config.MinioSecretKey, config.MinioBucket)
	if err != nil {
		log.Fatalf("failed to create minio client: %v", err)
	}

	repository := telegram.NewRawMessageRepository(client, minioClient)

	messages, err := repository.GetByID(context.Background(), "allrpn", 61)

	log.Printf("%v", messages.Media)
}
