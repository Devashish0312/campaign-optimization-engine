package utils

import (
	"testing"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func TestRedisCache(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Assuming Redis is running locally
	})

	ctx := context.Background()

	// Set a value in Redis
	err := rdb.Set(ctx, "campaign:123", "bid:0.5", 0).Err()
	if err != nil {
		t.Fatalf("Failed to set value in Redis: %v", err)
	}

	// Get the value from Redis
	val, err := rdb.Get(ctx, "campaign:123").Result()
	if err != nil {
		t.Fatalf("Failed to get value from Redis: %v", err)
	}

	if val != "bid:0.5" {
		t.Errorf("Expected bid:0.5, got %v", val)
	}
}
