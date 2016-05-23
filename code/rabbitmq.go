package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func rabbitsend(jsondata string) {
	conn, err := amqp.Dial("amqp://test:test@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	err = ch.ExchangeDeclare(
		"mse",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	q, err := ch.QueueDeclare(
		"mse", // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	err = ch.QueueBind(
		q.Name,
		"mse",
		"mse",
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	body := jsondata
	err = ch.Publish(
		"mse", // exchange
		"mse", // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
