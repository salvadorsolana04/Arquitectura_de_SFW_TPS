package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pedidos/services"
)

type ProductoController struct {
	service *services.ProductoService
}

func NewProductoController(service *services.ProductoService) *ProductoController {
	return &ProductoController{service: service}
}

func (pc *ProductoController) ListarProductos(c *gin.Context) {
	productos := pc.service.ListarProductos()
	c.JSON(http.StatusOK, gin.H{"productos": productos})
}
