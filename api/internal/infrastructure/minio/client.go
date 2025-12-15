package minio

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	client *minio.Client
	bucket string
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
		client: client,
		bucket: bucket,
	}, nil
}

func (c *Client) Upload(objectName string, data []byte, contentType string) error {
	_, err := c.client.PutObject(
		context.Background(),
		c.bucket,
		objectName,
		bytes.NewReader(data),
		int64(len(data)),
		minio.PutObjectOptions{ContentType: contentType},
	)
	return err
}
