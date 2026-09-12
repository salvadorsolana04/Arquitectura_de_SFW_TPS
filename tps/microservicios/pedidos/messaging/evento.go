package messaging

// EventoPedidoConfirmado es el mensaje publicado cuando un pedido se confirma.
type EventoPedidoConfirmado struct {
	Tipo       string `json:"tipo"`
	PedidoID   string `json:"pedido_id"`
	ClienteID  string `json:"cliente_id"`
	ProductoID string `json:"producto_id"`
}
