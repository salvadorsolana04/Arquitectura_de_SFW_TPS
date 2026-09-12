package services

import (
	"errors"

	"clientes/models"
	"clientes/repositories"
)

var ErrClienteNoEncontrado = errors.New("cliente no encontrado")

// ClienteService contiene las reglas de negocio del dominio "clientes".
type ClienteService struct {
	repo *repositories.ClienteRepository
}

func NewClienteService(repo *repositories.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) CrearCliente(nombre, email string) (models.Cliente, error) {
	if nombre == "" {
		return models.Cliente{}, errors.New("el nombre es obligatorio")
	}
	return s.repo.Guardar(nombre, email), nil
}

func (s *ClienteService) ObtenerCliente(id string) (models.Cliente, error) {
	cliente, existe := s.repo.BuscarPorID(id)
	if !existe {
		return models.Cliente{}, ErrClienteNoEncontrado
	}
	return cliente, nil
}
