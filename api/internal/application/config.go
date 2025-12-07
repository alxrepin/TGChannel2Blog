package application

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Config struct {
	APIID           int
	APIHash         string
	Phone           string
	ChannelUsername string
	DatabaseURL     string
	NATSURL         string
	DB              *pgxpool.Pool
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	apiIDStr := os.Getenv("TELEGRAM_API_ID")
	if apiIDStr == "" {
		return nil, errors.New("TELEGRAM_API_ID environment variable is required")
	}

	apiID, err := strconv.Atoi(apiIDStr)
	if err != nil {
		return nil, err
	}

	apiHash := os.Getenv("TELEGRAM_API_HASH")
	if apiHash == "" {
		return nil, errors.New("TELEGRAM_API_HASH environment variable is required")
	}

	phone := os.Getenv("TELEGRAM_PHONE")
	if phone == "" {
		return nil, errors.New("TELEGRAM_PHONE environment variable is required")
	}

	channel := os.Getenv("TELEGRAM_CHANNEL")
	if channel == "" {
		return nil, errors.New("TELEGRAM_CHANNEL environment variable is required")
	}

	databaseUser := os.Getenv("DB_USERNAME")
	if databaseUser == "" {
		return nil, errors.New("DB_USERNAME environment variable is required")
	}

	databasePassword := os.Getenv("DB_PASSWORD")
	if databasePassword == "" {
		return nil, errors.New("DB_PASSWORD environment variable is required")
	}

	databaseBase := os.Getenv("DB_DATABASE")
	if databaseBase == "" {
		return nil, errors.New("DB_DATABASE environment variable is required")
	}

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		databaseUser,
		databasePassword,
		databaseBase,
	)

	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		return nil, errors.New("NATS_URL environment variable is required")
	}

	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Config{
		APIID:           apiID,
		APIHash:         apiHash,
		Phone:           phone,
		ChannelUsername: channel,
		DatabaseURL:     databaseURL,
		NATSURL:         natsURL,
		DB:              db,
	}, nil
}
