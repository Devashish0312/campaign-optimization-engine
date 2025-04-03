package messagequeue

import (
	"testing"
	"log"
)

func TestSendMessage(t *testing.T) {
	// This test ensures that sending a message to the queue works without errors
	err := SendMessage("campaign_updates", "Test message")
	if err != nil {
		t.Fatalf("Failed to send message to RabbitMQ: %v", err)
	}
}
