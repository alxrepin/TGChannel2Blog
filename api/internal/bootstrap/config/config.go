package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type TelegramConfig struct {
	BotToken string
	AppID    int
	AppHash  string
	Phone    string
	Channel  string
}

type DatabaseConfig struct {
	UserName string
	Password string
	Database string
	URL      string
}

type BusConfig struct {
	URL string
}

type StorageConfig struct {
	MinioURL       string
	MinioAccessKey string
	MinioSecretKey string
	MinioBucket    string
}

type Config struct {
	Bus      BusConfig
	Database DatabaseConfig
	Telegram TelegramConfig
	Storage  StorageConfig
}

func MustLoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Load telegram config
	telegramAppIDStr := os.Getenv("TELEGRAM_APP_ID")
	if telegramAppIDStr == "" {
		panic("TELEGRAM_APP_ID environment variable is required")
	}

	telegramAppID, err := strconv.Atoi(telegramAppIDStr)
	if err != nil {
		panic(err)
	}

	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramBotToken == "" {
		panic(errors.New("TELEGRAM_BOT_TOKEN environment variable is required"))
	}

	telegramAppHash := os.Getenv("TELEGRAM_APP_HASH")
	if telegramAppHash == "" {
		panic(errors.New("TELEGRAM_APP_HASH environment variable is required"))
	}

	telegramPhone := os.Getenv("TELEGRAM_PHONE")
	if telegramPhone == "" {
		panic(errors.New("TELEGRAM_PHONE environment variable is required"))
	}

	telegramChannel := os.Getenv("TELEGRAM_CHANNEL")
	if telegramChannel == "" {
		panic(errors.New("TELEGRAM_CHANNEL environment variable is required"))
	}

	// Load database config
	databaseUser := os.Getenv("DB_USERNAME")
	if databaseUser == "" {
		panic(errors.New("DB_USERNAME environment variable is required"))
	}

	databasePassword := os.Getenv("DB_PASSWORD")
	if databasePassword == "" {
		panic(errors.New("DB_PASSWORD environment variable is required"))
	}

	databaseBase := os.Getenv("DB_DATABASE")
	if databaseBase == "" {
		panic(errors.New("DB_DATABASE environment variable is required"))
	}

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		databaseUser,
		databasePassword,
		databaseBase,
	)

	// Load bus config
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		panic(errors.New("NATS_URL environment variable is required"))
	}

	// Load storage config
	minioURL := os.Getenv("MINIO_URL")
	if minioURL == "" {
		panic(errors.New("MINIO_URL environment variable is required"))
	}

	minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	if minioAccessKey == "" {
		panic(errors.New("MINIO_ACCESS_KEY environment variable is required"))
	}

	minioSecretKey := os.Getenv("MINIO_SECRET_KEY")
	if minioSecretKey == "" {
		panic(errors.New("MINIO_SECRET_KEY environment variable is required"))
	}

	minioBucket := os.Getenv("MINIO_BUCKET")
	if minioBucket == "" {
		panic("MINIO_BUCKET environment variable is required")
	}

	return &Config{
		Telegram: TelegramConfig{
			BotToken: telegramBotToken,
			AppID:    telegramAppID,
			AppHash:  telegramAppHash,
			Phone:    telegramPhone,
			Channel:  telegramChannel,
		},
		Database: DatabaseConfig{
			UserName: databaseUser,
			Password: databasePassword,
			Database: databaseBase,
			URL:      databaseURL,
		},
		Bus: BusConfig{
			URL: natsURL,
		},
	}
}
