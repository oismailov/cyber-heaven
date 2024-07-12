package main

import (
	di "cyber-heaven-api/container"
	"cyber-heaven-api/internal/product"
	"cyber-heaven-api/pkg/queue/rabbitmq"
	"encoding/json"
	"log"
)

func main() {
	container := di.Bootstrap("config/config.json")
	if err := container.Invoke(func(rmqManager *rabbitmq.Manager) {
		messages, err := rmqManager.Consume()
		if err != nil {
			log.Fatalf("Failed to consume message: %v", err)
		}

		forever := make(chan bool)
		go func() {
			for message := range messages {
				var p product.Product
				if err = json.Unmarshal(message.Body, &p); err != nil {
					log.Fatalf("Failed to convert bytes to json config: %v", err)
				}
				log.Printf("message: %+v\n", p)
			}
		}()

		<-forever

	}); err != nil {
		log.Fatalf("Failed to invoke config: %v", err)
	}
}
