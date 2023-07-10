package brokers

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

// NewRabbitMQ creates a new RabbitMQ instance
func NewRabbitMQ() *RabbitMQ {
	log.Printf("dialing %q", "amqp://guest:guest@localhost:5672/")
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(fmt.Errorf("dial: %s", err))
		return nil
	}

	return &RabbitMQ{
		conn: connection,
	}
}

func (rabbit *RabbitMQ) Publish(body []byte) {

	log.Printf("got Connection, getting Channel")
	channel, err := rabbit.conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := channel.ExchangeDeclare(
		"gotr-city-exchange", // name
		"fanout",             // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // noWait
		nil,                  // arguments
	); err != nil {
		fmt.Println(err)
	}

	if err = channel.Publish(
		"gotr-city-exchange", // publish to an exchange
		"",                   // routing to 0 or more queues
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		fmt.Println(err)
	}
}
