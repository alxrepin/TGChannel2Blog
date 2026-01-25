package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	Pool *pgxpool.Pool
}

func MustNewClient(ctx context.Context, url string) *Client {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		panic(err)
	}

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	return &Client{
		Pool: pool,
	}
}

func (c *Client) Close() {
	c.Pool.Close()
}
