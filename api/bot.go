package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Bot token from environment
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN not set in .env")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Chat ID to send notifications (replace with real chat ID)
	chatID := int64(123456789)

	for update := range updates {
		var change map[string]interface{}

		if update.ChannelPost != nil {
			// New post
			change = map[string]interface{}{
				"type":    "new_post",
				"channel": update.ChannelPost.Chat.UserName,
				"message": update.ChannelPost,
			}
		} else if update.EditedChannelPost != nil {
			// Edited post
			change = map[string]interface{}{
				"type":    "edited_post",
				"channel": update.EditedChannelPost.Chat.UserName,
				"message": update.EditedChannelPost,
			}
		// Note: Deleted messages handling may require additional setup or different library
		} else {
			continue
		}

		// Check if from @allrpn
		if channel, ok := change["channel"].(string); ok && channel == "allrpn" {
			jsonBytes, err := json.MarshalIndent(change, "", "  ")
			if err != nil {
				log.Printf("JSON marshal error: %v", err)
				continue
			}

			msg := tgbotapi.NewMessage(chatID, string(jsonBytes))
			_, err = bot.Send(msg)
			if err != nil {
				log.Printf("Send error: %v", err)
			}
		}
	}
}