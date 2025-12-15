package telegram

import (
	"time"

	"app/internal/domain"

	"github.com/gotd/td/tg"
)

type RawMessageFactory struct{}

func NewRawMessageFactory() *RawMessageFactory {
	return &RawMessageFactory{}
}

func (f *RawMessageFactory) Create(msg *tg.Message) domain.RawMessage {
	rawMsg := domain.RawMessage{
		ID:      msg.ID,
		Date:    time.Unix(int64(msg.Date), 0),
		GroupID: msg.GroupedID,
	}

	if msg.Message != "" {
		rawMsg.Text = &msg.Message
	}

	if msg.Media != nil {
		media := f.extractMedia(msg.Media)

		if media != nil {
			rawMsg.Media = media
		}
	}

	if msg.Entities != nil {
		rawMsg.Entities = f.convertEntities(msg.Entities)
	}

	return rawMsg
}

func (f *RawMessageFactory) extractMedia(media tg.MessageMediaClass) *domain.Media {
	switch m := media.(type) {
	case *tg.MessageMediaPhoto:
		if m.Photo != nil {
			if photo, ok := m.Photo.(*tg.Photo); ok {
				return &domain.Media{
					Type:          domain.MediaTypePhoto,
					ID:            photo.ID,
					AccessHash:    photo.AccessHash,
					FileReference: photo.FileReference,
				}
			}
		}
	case *tg.MessageMediaDocument:
		doc := m.Document

		if doc, ok := doc.(*tg.Document); ok {
			mediaType := domain.MediaTypeDocument

			if doc.MimeType != "" {
				if len(doc.MimeType) >= 5 && doc.MimeType[:5] == "video" {
					mediaType = domain.MediaTypeVideo
				} else if len(doc.MimeType) >= 5 && doc.MimeType[:5] == "audio" {
					mediaType = domain.MediaTypeAudio
				}
			}

			return &domain.Media{
				Type:          mediaType,
				ID:            doc.ID,
				AccessHash:    doc.AccessHash,
				FileReference: doc.FileReference,
			}
		}
	}

	return nil
}

func (f *RawMessageFactory) convertEntities(entities []tg.MessageEntityClass) []domain.RawMessageEntity {
	var result []domain.RawMessageEntity

	for _, e := range entities {
		var entity domain.RawMessageEntity

		switch ent := e.(type) {
		case *tg.MessageEntityTextURL:
			entity = domain.RawMessageEntity{
				Type:   domain.EntityTypeURL,
				Offset: ent.Offset,
				Length: ent.Length,
				URL:    &ent.URL,
			}
		case *tg.MessageEntityURL:
			entity = domain.RawMessageEntity{
				Type:   domain.EntityTypeURL,
				Offset: ent.Offset,
				Length: ent.Length,
			}
		case *tg.MessageEntityBold:
			entity = domain.RawMessageEntity{
				Type:   domain.EntityTypeBold,
				Offset: ent.Offset,
				Length: ent.Length,
			}
		case *tg.MessageEntityItalic:
			entity = domain.RawMessageEntity{
				Type:   domain.EntityTypeItalic,
				Offset: ent.Offset,
				Length: ent.Length,
			}
		case *tg.MessageEntityCode:
			entity = domain.RawMessageEntity{
				Type:   domain.EntityTypeCode,
				Offset: ent.Offset,
				Length: ent.Length,
			}
		case *tg.MessageEntityPre:
			entity = domain.RawMessageEntity{
				Type:   domain.EntityTypePre,
				Offset: ent.Offset,
				Length: ent.Length,
			}
		case *tg.MessageEntityCustomEmoji:
			entity = domain.RawMessageEntity{
				Type:          domain.EntityTypeCustomEmoji,
				Offset:        ent.Offset,
				Length:        ent.Length,
				CustomEmojiID: &ent.DocumentID,
			}
		default:
			continue
		}

		result = append(result, entity)
	}

	return result
}
