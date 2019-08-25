package main

import (
	"strconv"
)

func main() {
	server := "localhost"
	numberOfClients := 5
	services := make([]string, numberOfClients)

	for i := 0; i < numberOfClients; i++ {
		services[i] = server + ":" + strconv.Itoa(8080+i)
	}

}
