package rabbitmq

import (
	"cyber-heaven-api/config"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func NewRabbitMQConnection(
	conf *config.Config,
) *amqp091.Connection {
	authUrl := fmt.Sprintf(
		"amqp://%s:%s@%s:%d",
		conf.Queue.User,
		conf.Queue.Password,
		conf.Queue.Host,
		conf.Queue.Port,
	)
	conn, err := amqp091.Dial(authUrl)
	if err != nil {
		log.Fatalf("unable to connect tot rabbit mq: %s", err)
		return nil
	}
	return conn
}

func NewRabbitMQChannel(
	mq *amqp091.Connection,
) *amqp091.Channel {
	ch, err := mq.Channel()
	if err != nil {
		return nil
	}

	return ch
}
