package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"pedidos/controllers"
	"pedidos/messaging"
	"pedidos/repositories"
	"pedidos/services"
)

const rabbitMQURL = "amqp://guest:guest@localhost:5672/"

func main() {
	publicador := messaging.NewRabbitMQPublisher(rabbitMQURL)
	defer publicador.Cerrar()

	productoRepo := repositories.NewProductoRepository()
	productoService := services.NewProductoService(productoRepo)
	productoController := controllers.NewProductoController(productoService)

	pedidoRepo := repositories.NewPedidoRepository()
	pedidoService := services.NewPedidoService(pedidoRepo, publicador)
	pedidoController := controllers.NewPedidoController(pedidoService)

	router := gin.Default()
	router.GET("/productos", productoController.ListarProductos)
	router.POST("/pedidos", pedidoController.ConfirmarPedido)

	log.Println("Microservicio pedidos escuchando en http://localhost:8082")
	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
