package repositories

import (
	"fmt"
	"sync"

	"pedidos/models"
)

// PedidoRepository guarda los pedidos confirmados en memoria.
type PedidoRepository struct {
	mu       sync.Mutex
	pedidos  map[string]models.Pedido
	contador int
}

func NewPedidoRepository() *PedidoRepository {
	return &PedidoRepository{
		pedidos: make(map[string]models.Pedido),
	}
}

func (r *PedidoRepository) Guardar(clienteID, productoID string) models.Pedido {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.contador++
	pedido := models.Pedido{
		ID:         fmt.Sprintf("PED-%d", r.contador),
		ClienteID:  clienteID,
		ProductoID: productoID,
		Estado:     "confirmado",
	}
	r.pedidos[pedido.ID] = pedido
	return pedido
}
