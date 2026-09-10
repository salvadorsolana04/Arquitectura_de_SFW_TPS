package repositories

import (
	"time"

	"pedidos/models"
)

// ProductoRepository simula el acceso a una fuente de datos de productos.
// El sleep representa el costo de una consulta "real" (DB u otro servicio),
// que es justamente lo que la caché de ProductoService evita repetir.
type ProductoRepository struct{}

func NewProductoRepository() *ProductoRepository {
	return &ProductoRepository{}
}

func (r *ProductoRepository) ListarTodos() []models.Producto {
	time.Sleep(50 * time.Millisecond)

	return []models.Producto{
		{ID: "P-1", Nombre: "Auriculares", Stock: 10},
		{ID: "P-2", Nombre: "Teclado", Stock: 8},
	}
}
