package utils

import (
	"log"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	return client
}
