package main

import (
	"github.com/lucas625/Middleware/mom-rpc/utils"
	"github.com/streadway/amqp"
)

func main() {
	// Connecting to rabbitmq server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.PrintError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Creating a channel
	ch, err := conn.Channel()
	utils.PrintError(err, "Failed to open a channel.")
	defer ch.Close()

	// Creating queues
	requestQueue, err := ch.QueueDeclare(
		"request", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	utils.PrintError(err, "Failed to declare a queue.")

	replyQueue, err := ch.QueueDeclare(
		"response", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	utils.PrintError(err, "Failed to declare a queue.")

	// Preparing to read messages from client
	msgfromClient, err := ch.Consume(
		requestQueue.Name, // queue
		"",                // consumer
		true,              // autoAck
		false,             // exclusive
		false,             // noLocal
		false,             // noWait
		nil,               // args
	)
	utils.PrintError(err, "Failed to consume from client.")

}
