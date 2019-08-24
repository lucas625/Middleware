package udp

import (
	"fmt"
	"net"
	"time"
)

// udp client
func udpClient(address string) (err error) {
	// gets the ip and port in the format: ip:port.
	ipPort, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	// connecting to server
	connection, err := net.DialUDP("udp", nil, ipPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("connected to server: %s", connection.RemoteAddr().String())

	// closes the connection once the function ends
	defer connection.Close()

	initialTime := time.Now()
	// write a message to server
	message := []byte("Hello UDP server!")

	_, err = connection.Write(message)

	if err != nil {
		fmt.Println(err)
	}
	go func() {
		// time.Now().Add uses nanoseconds
		deadline := time.Now().Add(2 * 10000000000)
		err = connection.SetReadDeadline(deadline)
		// receive message from server
		buffer := make([]byte, 1024)
		n, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received from UDP server : ", string(buffer[:n]))
	}()
	// finding the rtt
	endTime := time.Now()
	fmt.Printf("The RTT took: %v.\n", endTime.Sub(initialTime))

	return
}
