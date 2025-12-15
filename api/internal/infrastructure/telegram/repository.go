package telegram

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"app/internal/domain"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/downloader"
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

func (r *RawMessageRepository) DownloadMedia(ctx context.Context, media domain.Media) ([]byte, error) {
	var data []byte

	err := r.client.Run(ctx, func(ctx context.Context, client *telegram.Client) error {
		var location tg.InputFileLocationClass

		fmt.Printf(string(media.FileReference))
		switch media.Type {
		case domain.MediaTypePhoto:
			location = &tg.InputPhotoFileLocation{
				ID:            media.ID,
				AccessHash:    media.AccessHash,
				FileReference: media.FileReference,
			}
		case domain.MediaTypeVideo, domain.MediaTypeAudio, domain.MediaTypeDocument:
			location = &tg.InputDocumentFileLocation{
				ID:            media.ID,
				AccessHash:    media.AccessHash,
				FileReference: media.FileReference,
			}
		default:
			return fmt.Errorf("unsupported media type: %s", media.Type)
		}

		d := downloader.NewDownloader()
		builder := d.Download(client.API(), location)
		buf := &bytes.Buffer{}
		_, err := builder.Stream(ctx, buf)
		if err != nil {
			return fmt.Errorf("failed to stream media: %w", err)
		}
		data = buf.Bytes()

		return nil
	})

	return data, err
}

func (r *RawMessageRepository) GetChannelInfo(ctx context.Context, channelUsername string) (*domain.Channel, error) {
	var channel *domain.Channel

	err := r.client.Run(ctx, func(ctx context.Context, client *telegram.Client) error {
		tgChannel, about, subscriptions, err := r.client.FetchChannelInfo(ctx, client, channelUsername)
		if err != nil {
			return err
		}

		// Print avatar link to console
		if tgChannel.Photo != nil {
			//if photo, ok := tgChannel.Photo.(*tg.ChatPhoto); ok {
			//	//fmt.Printf("Avatar link: https://t.me/%s\n", channelUsername, photo)
			//}
		}

		channel = &domain.Channel{
			ID:            tgChannel.ID,
			Name:          channelUsername,
			Title:         tgChannel.Title,
			Description:   about,
			Avatar:        nil,
			CreatedAt:     time.Unix(int64(tgChannel.Date), 0),
			Subscriptions: subscriptions,
		}

		return nil
	})

	return channel, err
}
