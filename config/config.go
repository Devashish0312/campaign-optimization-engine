package config

import (
	"log"
	"os"
)

var Config struct {
	DatabaseURL string
	RedisURL    string
	MessageQueueURL string
}

func LoadConfig() {
	Config.DatabaseURL = os.Getenv("DATABASE_URL")
	Config.RedisURL = os.Getenv("REDIS_URL")
	Config.MessageQueueURL = os.Getenv("MESSAGE_QUEUE_URL")

	if Config.DatabaseURL == "" || Config.RedisURL == "" || Config.MessageQueueURL == "" {
		log.Fatal("Missing required environment variables")
	}
}
