# eventos

Esta carpeta representa conceptualmente al consumidor de logística.

Según la consigna del TP1, logística **no debe implementarse**: el evento
`pedido.confirmado` se publica desde `pedidos/messaging/rabbitmq_publisher.go`
hacia la cola `pedidos-confirmados`, y este espacio queda como referencia de
dónde, en un caso real, se agregaría el microservicio consumidor.
