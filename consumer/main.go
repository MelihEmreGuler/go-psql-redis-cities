package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}

	//NotifyClose returns a chan *Error which will be notified when the connection is closed.
	go func() {
		fmt.Printf("closing: %s", <-conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("declared Exchange, declaring Queue %q", "long-running-task-queue")
	queue, err := channel.QueueDeclare(
		"long-running-task-queue", // name of the queue
		true,                      // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // noWait
		nil,                       // arguments
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, "")

	if err = channel.QueueBind(
		queue.Name,           // name of the queue
		"",                   // bindingKey
		"gotr-city-exchange", // sourceExchange
		false,                // noWait
		nil,                  // arguments
	); err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", "city-consumer")
	deliveries, err := channel.Consume(
		queue.Name,      // name
		"city-consumer", // consumerTag,
		false,           // noAck
		false,           // exclusive
		false,           // noLocal
		false,           // noWait
		nil,             // arguments
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		message := <-deliveries
		fmt.Println("read data")

		//print the published message to the console
		fmt.Println(string(message.Body))

		//async or long-running task here ...

		//acknowledge the message
		message.Ack(false)
	}
}
