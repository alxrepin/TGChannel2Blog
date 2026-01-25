package minio

import (
	"app/internal/context/domain"
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	client   *minio.Client
	bucket   string
	endpoint string
}

func NewClient(endpoint, accessKey, secretKey, bucket string) (*Client, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // Set to true if using HTTPS
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
	}, nil
}

func (c *Client) Upload(ctx context.Context, objectName string, data []byte, contentType string) (string, error) {
	_, err := c.client.PutObject(
		ctx,
		c.bucket,
		objectName,
		bytes.NewReader(data),
		int64(len(data)),
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return "", err
	}

	// Return the URL to the uploaded object
	url := fmt.Sprintf("http://%s/%s/%s", c.endpoint, c.bucket, objectName)
	return url, nil
}

// Ensure Client implements service.Storage
var _ domain.Storage = (*Client)(nil)
