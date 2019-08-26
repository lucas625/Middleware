package main

import (
	"strconv"
	client "Go\src\Middleware\client-server\src\udp\client"
)

func main() {
	server := "localhost"
	numberOfClients := 5
	services := make([]string, numberOfClients)

	for i := 0; i < numberOfClients; i++ {
		services[i] = server + ":" + strconv.Itoa(8080+i)
		go client.UdpClien(services[i])
	}

}
