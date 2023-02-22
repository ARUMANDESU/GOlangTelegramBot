package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"telegrambot/bot"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	token, exists := os.LookupEnv("TELEGRAM_BOT_TOKEN")
	if !exists {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable not set.")
	}

	// Create new bot instance
	b, err := bot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	// Add command and message handlers
	b.AddHandlers()

	// Start bot
	b.Start()
}
