package main

import (
	"fmt"
	"log"
	"os"

	"github.com/diogocarasco/golang-amqp-consumer/config"
	handler_negocio1 "github.com/diogocarasco/golang-amqp-consumer/internal/handlers/negocio1"
	amqp "github.com/streadway/amqp"
)

func main() {

	config.Load()

	queue := os.Args[1]

	amqpServerURL := "amqp://guest:guest@message-broker:5672/"

	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Subscribing to queue from parameter for getting messages.
	messages, err := channelRabbitMQ.Consume(
		queue, // queue name
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // arguments
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	forever := make(chan bool)

	go func() {
		for message := range messages {

			switch message.Headers["messageEvent"] {
			case "PaymentCreated":
				handler_negocio1.Handle(message.Headers, string(message.Body))
				continue
			case "PaymentFailed":
				fmt.Println("falhou")
				continue
			default:
				fmt.Println("Message received with unknown event")
			}

		}
	}()

	<-forever
}
