package product

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := h.Repository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	//go func(p *Product) {
	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Error marshalling struct to JSON: %s", err)
	}

	if err = h.RabbitMQManager.Publish(productJSON); err != nil {
		log.Fatalf("error publishing to rabbit mq: %s", err)
	}
	//}(product)

	c.JSON(200, product)
}
