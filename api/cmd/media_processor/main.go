package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"app/internal/application"
	"app/internal/domain"
	"app/internal/infrastructure/minio"
	"app/internal/infrastructure/postgres"
	"app/internal/infrastructure/telegram"

	"github.com/jackc/pgx/v5/pgxpool"
	nc "github.com/nats-io/nats.go"
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
	telegramRepo := telegram.NewRawMessageRepository(telegramClient, minioClient)
	channelRepo := postgres.NewChannelRepository(db)

	sub, err := js.PullSubscribe(string(domain.MediaReceived), "media_processor")
	if err != nil {
		log.Fatalf("failed to subscribe: %v", err)
	}

	log.Println("Media processor started, waiting for messages...")

	for {
		msgs, err := sub.Fetch(1, nc.MaxWait(5*time.Second))
		if err != nil {
			if err == nc.ErrTimeout {
				continue
			}
			log.Printf("failed to fetch messages: %v", err)
			continue
		}

		for _, msg := range msgs {
			var mediaMsg domain.MediaMessage
			err := json.Unmarshal(msg.Data, &mediaMsg)
			if err != nil {
				log.Printf("failed to unmarshal media message: %v", err)
				msg.Ack()
				continue
			}

			channel, err := channelRepo.Get(context.Background())
			if err != nil {
				log.Printf("failed to get channel %d: %v", mediaMsg.GroupID, err)
				msg.Ack()
				continue
			}

			rawMessage, err := telegramRepo.GetByID(context.Background(), channel.Name, mediaMsg.ID)
			if err != nil {
				log.Printf("failed to get raw message %d: %v", mediaMsg.ID, err)
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
				msg.Ack()
				continue
			}

			objectName := fmt.Sprintf("%d-%s", rawMessage.ID, rawMessage.Media.Type)
			contentType := getContentType(rawMessage.Media.Type)

			_, err = minioClient.Upload(context.Background(), objectName, data, contentType)
			if err != nil {
				log.Printf("failed to upload media for message %d: %v", rawMessage.ID, err)
				msg.Ack()
				continue
			}

			log.Printf("Successfully processed media for message %d", rawMessage.ID)
			msg.Ack()
		}
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
