package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"clientes/services"
)

type ClienteController struct {
	service *services.ClienteService
}

func NewClienteController(service *services.ClienteService) *ClienteController {
	return &ClienteController{service: service}
}

type crearClienteRequest struct {
	Nombre string `json:"nombre" binding:"required"`
	Email  string `json:"email"`
}

func (cc *ClienteController) CrearCliente(c *gin.Context) {
	var req crearClienteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cliente, err := cc.service.CrearCliente(req.Nombre, req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cliente)
}

func (cc *ClienteController) ObtenerCliente(c *gin.Context) {
	id := c.Param("id")

	cliente, err := cc.service.ObtenerCliente(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cliente)
}
