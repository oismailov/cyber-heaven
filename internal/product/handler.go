package product

import (
	"cyber-heaven-api/config"
	"cyber-heaven-api/pkg/queue/rabbitmq"
)

type Handler struct {
	Config          *config.Config
	Repository      Repository
	RabbitMQManager *rabbitmq.Manager
}

func NewHandler(
	config *config.Config,
	repository Repository,
	rabbitMQManager *rabbitmq.Manager,
) *Handler {
	return &Handler{
		Config:          config,
		Repository:      repository,
		RabbitMQManager: rabbitMQManager,
	}
}
