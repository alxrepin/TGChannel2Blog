package telegram

import (
	"context"

	"app/internal/domain"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)

type RawMessageRepository struct {
	client  *Client
	factory *RawMessageFactory
}

func NewRawMessageRepository(client *Client) *RawMessageRepository {
	return &RawMessageRepository{
		client:  client,
		factory: NewRawMessageFactory(),
	}
}

func (r *RawMessageRepository) GetAll(ctx context.Context, channelUsername string) ([]domain.RawMessage, error) {
	var messages []domain.RawMessage

	err := r.client.Run(ctx, func(ctx context.Context, client *telegram.Client) error {
		tgMessages, err := r.client.FetchMessages(ctx, client, channelUsername)
		if err != nil {
			return err
		}

		for _, msg := range tgMessages {
			if tgMsg, ok := msg.(*tg.Message); ok {
				rawMsg := r.factory.Create(tgMsg)
				messages = append(messages, rawMsg)
			}
		}

		return nil
	})

	return messages, err
}
