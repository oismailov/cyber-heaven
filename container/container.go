package container

import (
	"cyber-heaven-api/config"
	"cyber-heaven-api/internal/product"
	"cyber-heaven-api/pkg/database"
	"cyber-heaven-api/pkg/queue/rabbitmq"
	"go.uber.org/dig"
)

func Bootstrap(configFile string) *dig.Container {
	container := dig.New()

	bootstrapConfig(configFile, container)
	loadDependencies(container)

	return container
}

func loadDependencies(container *dig.Container) {
	dependencies := []interface{}{
		database.NewMySqlConnection,
		product.NewProductRepository,
		product.NewHandler,
		rabbitmq.NewRabbitMQConnection,
		rabbitmq.NewRabbitMQChannel,
		rabbitmq.NewManager,
	}

	for _, dependency := range dependencies {
		if err := container.Provide(dependency); err != nil {
			panic(err)
		}
	}
}

func bootstrapConfig(configFile string, container *dig.Container) {
	if err := container.Provide(func() (*config.Config, error) {
		return config.LoadConfig(configFile)
	}); err != nil {
		panic(err)
	}
}
