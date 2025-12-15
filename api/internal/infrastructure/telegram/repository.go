package telegram

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"app/internal/domain"
	"app/internal/domain/service"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/tg"
)

type RawMessageRepository struct {
	client  *Client
	factory *RawMessageFactory
	storage service.Storage
}

func NewRawMessageRepository(client *Client, storage service.Storage) *RawMessageRepository {
	return &RawMessageRepository{
		client:  client,
		factory: NewRawMessageFactory(),
		storage: storage,
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

		switch media.Type {
		case domain.MediaTypePhoto:
			location = &tg.InputPhotoFileLocation{
				ID:            media.ID,
				AccessHash:    media.AccessHash,
				FileReference: media.FileReference,
				ThumbSize:     media.PhotoSizeType,
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

		var avatarURL *string
		if tgChannel.Photo != nil {
			if photo, ok := tgChannel.Photo.(*tg.ChatPhoto); ok {
				// Download the photo using InputPeerPhotoFileLocation
				location := &tg.InputPeerPhotoFileLocation{
					Peer:    &tg.InputPeerChannel{ChannelID: tgChannel.ID, AccessHash: tgChannel.AccessHash},
					PhotoID: photo.PhotoID,
				}

				d := downloader.NewDownloader()
				builder := d.Download(client.API(), location)
				buf := &bytes.Buffer{}
				_, err := builder.Stream(ctx, buf)
				if err != nil {
					return fmt.Errorf("failed to download channel photo: %w", err)
				}

				// Upload to storage
				objectName := fmt.Sprintf("channel_avatars/%d.jpg", tgChannel.ID)
				url, err := r.storage.Upload(ctx, objectName, buf.Bytes(), "image/jpeg")
				if err != nil {
					return fmt.Errorf("failed to upload channel photo: %w", err)
				}

				avatarURL = &url
			}
		}

		channel = &domain.Channel{
			ID:            tgChannel.ID,
			Name:          channelUsername,
			Title:         tgChannel.Title,
			Description:   about,
			Avatar:        avatarURL,
			Subscriptions: subscriptions,
			CreatedAt:     time.Unix(int64(tgChannel.Date), 0),
		}

		return nil
	})

	return channel, err
}

func (r *RawMessageRepository) GetByID(ctx context.Context, channelUsername string, messageID int) (*domain.RawMessage, error) {
	var message *domain.RawMessage

	err := r.client.Run(ctx, func(ctx context.Context, client *telegram.Client) error {
		tgMessage, err := r.client.FetchMessage(ctx, client, channelUsername, messageID)
		if err != nil {
			return err
		}

		if tgMessage == nil {
			return nil
		}

		rawMsg := r.factory.Create(tgMessage)
		message = &rawMsg

		return nil
	})

	return message, err
}
