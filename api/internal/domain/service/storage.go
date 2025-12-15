package service

import "context"

type Storage interface {
	Upload(ctx context.Context, objectName string, data []byte, contentType string) (string, error)
}
