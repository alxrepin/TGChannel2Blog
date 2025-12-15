package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"app/internal/application"
	"app/internal/application/usecase/sync"
	"app/internal/domain"
	"app/internal/infrastructure/bus"
	"app/internal/infrastructure/postgres"

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

	bus := bus.NewNatsBus(js)

	postRepository := postgres.NewPostRepository(db)
	uc := sync.NewSyncRawMessageUseCase(postRepository)

	sub, err := js.PullSubscribe(string(domain.RawMessageReceived), "subscriber")
	if err != nil {
		log.Fatalf("failed to subscribe: %v", err)
	}

	log.Println("Subscriber started, waiting for messages...")

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
			var rawMessage domain.RawMessage
			err := json.Unmarshal(msg.Data, &rawMessage)
			if err != nil {
				log.Printf("failed to unmarshal message: %v", err)
				msg.Ack()
				continue
			}

			if rawMessage.ID != 61 {
				continue
			}

			if (rawMessage.Text == nil || *rawMessage.Text == "") && rawMessage.GroupID > 0 {
				mediaMsg := domain.MediaMessage{
					ID:      rawMessage.ID,
					GroupID: rawMessage.GroupID,
				}
				data, err := json.Marshal(mediaMsg)
				if err != nil {
					log.Printf("failed to marshal media message: %v", err)
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

			if rawMessage.Media != nil {
				mediaMsg := domain.MediaMessage{
					ID:      rawMessage.ID,
					GroupID: rawMessage.GroupID,
				}
				data, err := json.Marshal(mediaMsg)
				if err != nil {
					log.Printf("failed to marshal media message: %v", err)
				} else {
					log.Printf("dispatching media message: %v", string(data))
					err = bus.Dispatch(domain.MediaReceived, data)
					if err != nil {
						log.Printf("failed to dispatch media message: %v", err)
					}
				}
			}

			err = uc.Execute(context.Background(), rawMessage)
			if err != nil {
				log.Printf("failed to process message %d: %v", rawMessage.ID, err)
				msg.Nak()
			} else {
				log.Printf("Successfully processed message %d", rawMessage.ID)
				msg.Ack()
			}
		}
	}
}
