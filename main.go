package main

import (
	"cyber-heaven-api/config"
	di "cyber-heaven-api/container"
	"cyber-heaven-api/router"
	"fmt"
	"log"
)

func main() {
	container := di.Bootstrap("config/config.json")

	r := router.LoadRoutes(container)

	if err := container.Invoke(func(config *config.Config) {
		if e := r.Run(fmt.Sprintf(":%d", config.Server.Port)); e != nil {
			log.Fatalf("Failed to start server: %v", e)
		}
	}); err != nil {
		log.Fatalf("Failed to invoke config: %v", err)
	}
}
