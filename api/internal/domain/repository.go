package domain

import "context"

type RawMessageRepository interface {
	GetAll(ctx context.Context, channelUsername string) ([]RawMessage, error)
}

type PostRepository interface {
	CreateOrUpdate(ctx context.Context, post *Post) error
	GetList(ctx context.Context, page, limit int) ([]Post, int, error)
}
