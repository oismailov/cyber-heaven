package product

import (
	"github.com/gin-gonic/gin"
)

func LoadProductRoutes(
	r *gin.RouterGroup,
	h *Handler,
) {
	r.GET("products/:id", h.GetProduct)
	//r.GET("products", h.ListProduct)
}
