package main

import (
	"fmt"
	"log"
	"net"
	"time"
	"strconv"
	"sync"
)

// udp client
func UdpClient(address string, wg *sync.WaitGroup) (final_time int) {
	defer wg.Done()
	// gets the ip and port in the format: ip:port.
	ipPort, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalln(err)
		return -1
	}

	// connecting to server
	connection, err := net.DialUDP("udp", nil, ipPort)
	if err != nil {
		log.Fatalln(err)
		return -1
	}

	fmt.Printf("connected to server: %s.\n", connection.RemoteAddr().String())

	// closes the connection once the function ends
	defer connection.Close()

	initialTime := time.Now()
	// write a message to server
	message := []byte(strconv.Itoa(13))

	_, err = connection.Write(message)

	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		// time.Now().Add uses nanoseconds
		deadline := time.Now().Add(2 * 10000000000)
		err = connection.SetReadDeadline(deadline)
		// receive message from server
		buffer := make([]byte, 1024)
		n, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Received from UDP server : ", string(buffer[:n]))
	}()
	// finding the rtt
	endTime := int(time.Now().Sub(initialTime))
	fmt.Printf("The RTT took: %0.4fms.\n", float32(endTime)/1000000)

	return endTime
}

func main() {
	server := "localhost"
	numberOfClients := 5
	services := make([]string, numberOfClients)

	var wg sync.WaitGroup

	for i := 0; i < numberOfClients; i++ {
		services[i] = server + ":" + strconv.Itoa(8080+i)
		wg.Add(1)
		go UdpClient(services[i], &wg)
	}

	wg.Wait()

}