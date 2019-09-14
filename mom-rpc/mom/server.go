package main

import (
	"github.com/lucas625/Projeto-CG/mom-rpc/utils"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.PrintError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.PrintError(err, "Failed to open a channel")
	defer ch.Close()

}