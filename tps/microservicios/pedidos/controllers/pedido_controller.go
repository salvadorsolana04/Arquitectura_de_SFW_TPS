package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pedidos/services"
)

type PedidoController struct {
	service *services.PedidoService
}

func NewPedidoController(service *services.PedidoService) *PedidoController {
	return &PedidoController{service: service}
}

type confirmarPedidoRequest struct {
	ClienteID  string `json:"cliente_id" binding:"required"`
	ProductoID string `json:"producto_id" binding:"required"`
}

func (pc *PedidoController) ConfirmarPedido(c *gin.Context) {
	var req confirmarPedidoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pedido, err := pc.service.ConfirmarPedido(req.ClienteID, req.ProductoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje":     "pedido confirmado",
		"pedido_id":   pedido.ID,
		"cliente_id":  pedido.ClienteID,
		"producto_id": pedido.ProductoID,
	})
}
