package messaging

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const colaPedidosConfirmados = "pedidos-confirmados"

// RabbitMQPublisher publica eventos en la cola "pedidos-confirmados".
// Si RabbitMQ no está disponible al arrancar, el publisher queda inactivo
// y el microservicio sigue funcionando igual (solo loguea que no pudo publicar).
type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQPublisher(url string) *RabbitMQPublisher {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Printf("no se pudo conectar a RabbitMQ (%s): %v", url, err)
		return &RabbitMQPublisher{}
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Printf("no se pudo abrir el canal de RabbitMQ: %v", err)
		return &RabbitMQPublisher{}
	}

	if _, err := channel.QueueDeclare(
		colaPedidosConfirmados,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,
	); err != nil {
		log.Printf("no se pudo declarar la cola %s: %v", colaPedidosConfirmados, err)
		return &RabbitMQPublisher{}
	}

	log.Println("conectado a RabbitMQ, publicando en la cola", colaPedidosConfirmados)
	return &RabbitMQPublisher{conn: conn, channel: channel}
}

func (p *RabbitMQPublisher) PublicarPedidoConfirmado(evento EventoPedidoConfirmado) error {
	if p.channel == nil {
		log.Printf("RabbitMQ no disponible, se omite la publicación del evento %s", evento.PedidoID)
		return nil
	}

	cuerpo, err := json.Marshal(evento)
	if err != nil {
		return err
	}

	return p.channel.PublishWithContext(
		context.Background(),
		"",                     // exchange
		colaPedidosConfirmados, // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        cuerpo,
		},
	)
}

func (p *RabbitMQPublisher) Cerrar() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}
