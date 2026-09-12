package services

import (
	"log"

	"pedidos/messaging"
	"pedidos/models"
	"pedidos/repositories"
)

// Publicador desacopla al servicio de la implementación concreta de mensajería.
type Publicador interface {
	PublicarPedidoConfirmado(evento messaging.EventoPedidoConfirmado) error
}

// PedidoService confirma pedidos y publica el evento correspondiente.
type PedidoService struct {
	repo       *repositories.PedidoRepository
	publicador Publicador
}

func NewPedidoService(repo *repositories.PedidoRepository, publicador Publicador) *PedidoService {
	return &PedidoService{repo: repo, publicador: publicador}
}

func (s *PedidoService) ConfirmarPedido(clienteID, productoID string) (models.Pedido, error) {
	pedido := s.repo.Guardar(clienteID, productoID)

	evento := messaging.EventoPedidoConfirmado{
		Tipo:       "pedido.confirmado",
		PedidoID:   pedido.ID,
		ClienteID:  pedido.ClienteID,
		ProductoID: pedido.ProductoID,
	}

	if err := s.publicador.PublicarPedidoConfirmado(evento); err != nil {
		log.Printf("error al publicar el evento del pedido %s: %v", pedido.ID, err)
	}

	return pedido, nil
}
