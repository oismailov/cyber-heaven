package product

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListProduct(c *gin.Context) {
	c.JSON(200, []Product{{
		ID:   1,
		Name: "Apple",
	}, {
		ID:   2,
		Name: "Samsung",
	}})
}
