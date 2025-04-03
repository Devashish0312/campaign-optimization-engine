package messagequeue

import (
	"log"
	"github.com/streadway/amqp"
)

type MessageQueueService struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewMessageQueueService() *MessageQueueService {
	return &MessageQueueService{}
}

func (m *MessageQueueService) Start() {
	log.Println("Message Queue Service Started")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	m.conn = conn
	defer m.conn.Close()

	ch, err := m.conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	m.ch = ch
	defer m.ch.Close()

	// Create a queue
	_, err = m.ch.QueueDeclare(
		"campaign_updates", // Queue name
		false,              // Durable
		false,              // Delete when unused
		false,              // Exclusive
		false,              // No-wait
		nil,                // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	// Send a test message
	err = m.ch.Publish(
		"",                 // Default exchange
		"campaign_updates", // Queue name
		false,              // Mandatory
		false,              // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Campaign update message"),
		},
	)
	if err != nil {
		log.Fatal("Failed to publish a message:", err)
	}
}
