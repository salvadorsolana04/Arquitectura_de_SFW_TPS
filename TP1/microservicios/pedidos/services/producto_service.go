package services

import (
	"sync"
	"time"

	"pedidos/models"
	"pedidos/repositories"
)

const cacheTTL = 30 * time.Second

// ProductoService expone el catálogo de productos con una caché simple
// en memoria para evitar golpear el repositorio en cada consulta.
type ProductoService struct {
	repo *repositories.ProductoRepository

	mu           sync.Mutex
	cache        []models.Producto
	cacheExpira  time.Time
}

func NewProductoService(repo *repositories.ProductoRepository) *ProductoService {
	return &ProductoService{repo: repo}
}

func (s *ProductoService) ListarProductos() []models.Producto {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cache != nil && time.Now().Before(s.cacheExpira) {
		return s.cache
	}

	productos := s.repo.ListarTodos()
	s.cache = productos
	s.cacheExpira = time.Now().Add(cacheTTL)
	return productos
}
