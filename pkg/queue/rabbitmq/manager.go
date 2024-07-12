package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type Manager struct {
	Channel *amqp091.Channel
}

func NewManager(
	channel *amqp091.Channel,
) *Manager {
	return &Manager{Channel: channel}
}

func (m *Manager) Publish(message []byte) error {
	q := m.declareQueue()
	if err := m.Channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	); err != nil {
		log.Fatalf("unable to publish the message: %s", err)
		return err
	}

	return nil
}

func (m *Manager) Consume() (<-chan amqp091.Delivery, error) {
	q := m.declareQueue()
	messages, err := m.Channel.Consume(
		q.Name,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args,
	)
	if err != nil {
		log.Fatalf("unable to publish the message: %s", err)
		return nil, err
	}
	return messages, nil
}

func (m *Manager) declareQueue() amqp091.Queue {
	q, err := m.Channel.QueueDeclare(
		"product-log", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatalf("unable to declare queue: %s", err)
	}

	return q
}
