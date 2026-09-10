package models

// Pedido representa un pedido confirmado.
type Pedido struct {
	ID         string `json:"pedido_id"`
	ClienteID  string `json:"cliente_id"`
	ProductoID string `json:"producto_id"`
	Estado     string `json:"estado"`
}
