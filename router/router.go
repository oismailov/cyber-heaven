package router

import (
	"cyber-heaven-api/internal/product"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"log"
	"net/http"
)

func LoadRoutes(
	di *dig.Container,
) *gin.Engine {
	r := gin.Default()

	routeGroup := r.Group("v1/")

	if err := di.Invoke(func(
		productHandler *product.Handler,
	) {
		product.LoadProductRoutes(routeGroup, productHandler)
	}); err != nil {
		log.Fatalf("Unable to invoke handlers: %v", err)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "page not found",
		})
	})

	return r
}
