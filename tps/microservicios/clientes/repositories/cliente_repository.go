package repositories

import (
	"fmt"
	"sync"

	"clientes/models"
)

// ClienteRepository guarda los clientes en memoria.
type ClienteRepository struct {
	mu        sync.RWMutex
	clientes  map[string]models.Cliente
	contador  int
}

func NewClienteRepository() *ClienteRepository {
	return &ClienteRepository{
		clientes: make(map[string]models.Cliente),
	}
}

func (r *ClienteRepository) Guardar(nombre, email string) models.Cliente {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.contador++
	cliente := models.Cliente{
		ID:     fmt.Sprintf("C-%d", r.contador),
		Nombre: nombre,
		Email:  email,
	}
	r.clientes[cliente.ID] = cliente
	return cliente
}

func (r *ClienteRepository) BuscarPorID(id string) (models.Cliente, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cliente, existe := r.clientes[id]
	return cliente, existe
}
