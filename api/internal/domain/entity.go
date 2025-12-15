package domain

import "time"

type MediaType string

const (
	MediaTypePhoto    MediaType = "photo"
	MediaTypeVideo    MediaType = "video"
	MediaTypeAudio    MediaType = "audio"
	MediaTypeDocument MediaType = "document"
)

type Media struct {
	Type          MediaType `json:"type"`
	ID            int64     `json:"id"`
	AccessHash    int64     `json:"access_hash"`
	FileReference []byte    `json:"file_reference"`
}

type RawMessageEntityType string

const (
	EntityTypeURL         RawMessageEntityType = "url"
	EntityTypeBold        RawMessageEntityType = "bold"
	EntityTypeItalic      RawMessageEntityType = "italic"
	EntityTypeCode        RawMessageEntityType = "code"
	EntityTypePre         RawMessageEntityType = "pre"
	EntityTypeTextLink    RawMessageEntityType = "text_link"
	EntityTypeCustomEmoji RawMessageEntityType = "custom_emoji"
)

type RawMessageEntity struct {
	Type          RawMessageEntityType `json:"type"`
	Offset        int                  `json:"offset"`
	Length        int                  `json:"length"`
	URL           *string              `json:"url,omitempty"`
	User          *int64               `json:"user,omitempty"`
	CustomEmojiID *int64               `json:"custom_emoji_id,omitempty"`
}

type RawMessage struct {
	ID       int                `json:"id"`
	Text     *string            `json:"text,omitempty"`
	Date     time.Time          `json:"date"`
	GroupID  int64              `json:"group_id"`
	Media    *Media             `json:"media,omitempty"`
	Entities []RawMessageEntity `json:"entities,omitempty"`
}

type Post struct {
	ID             int64      `json:"id" db:"id"`
	GroupID        int64      `json:"group_id" db:"group_id"`
	Title          *string    `json:"title,omitempty" db:"title"`
	URL            *string    `json:"url,omitempty" db:"url"`
	Text           *string    `json:"text,omitempty" db:"text"`
	SEOTitle       *string    `json:"seo_title,omitempty" db:"seo_title"`
	SEODescription *string    `json:"seo_description,omitempty" db:"seo_description"`
	SEOKeywords    *string    `json:"seo_keywords,omitempty" db:"seo_keywords"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type Channel struct {
	ID            int64   `json:"id" db:"id"`
	Name          string  `json:"name" db:"name"`
	Title         string  `json:"title" db:"title"`
	Description   *string `json:"description,omitempty" db:"description"`
	Avatar        *string `json:"avatar,omitempty" db:"avatar"`
	Subscriptions int64   `json:"subscriptions" db:"subscriptions"`
}
