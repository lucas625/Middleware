package main

import (
	"log"
	"mom-rpc/utils"
	"github.com/streadway/amqp"// go rabbitMQ client library
  )



func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")// Standard
	utils.PrintError()
}