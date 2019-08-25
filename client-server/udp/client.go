package udp

import (
	"fmt"
	"log"
	"net"
	"time"
)

// udp client
func udpClient(address string) (final_time int) {
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

	fmt.Println("connected to server: %s", connection.RemoteAddr().String())

	// closes the connection once the function ends
	defer connection.Close()

	initialTime := time.Now()
	// write a message to server
	message := []byte("Hello UDP server!")

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
			log.Fatalln(err)
			return
		}
		fmt.Println("Received from UDP server : ", string(buffer[:n]))
	}()
	// finding the rtt
	endTime := int(time.Now().Sub(initialTime))
	fmt.Printf("The RTT took: %d.\n", endTime)

	return endTime
}
