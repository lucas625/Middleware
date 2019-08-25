package main

import (
	"fmt"
	"log"
	"net"
)

func handleUDPConnection(conn *net.UDPConn) {

	// here is where you want to do stuff like read or write to client

	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP client : ", addr)
	fmt.Println("Received from UDP client :  ", string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}

	// NOTE : Need to specify client address in WriteToUDP() function
	//        otherwise, you will get this error message
	//        write udp : write: destination address required if you use Write() function instead of WriteToUDP()

	// write message back to client
	message := []byte("Hello UDP client!")
	_, err = conn.WriteToUDP(message, addr)

	if err != nil {
		log.Println(err)
	}

}

func main() {
	hostName := "localhost"
	// 5 clients allowed at the same time
	portNums := []string{"8080", "8081", "8082", "8083", "8084"}
	services := make([]string, len(portNums))
	udpAddresses := make([]*net.UDPAddr, len(portNums))
	listeners := make([]*net.UDPConn, len(portNums))

	// filling the arrays
	for i, _ := range portNums {
		services[i] = hostName + ":" + portNums[i]

		// resolving the address
		udpAdd, err := net.ResolveUDPAddr("udp", services[i])
		if err != nil {
			log.Fatal(err)
		}
		udpAddresses[i] = udpAdd

		// listener for UDP connection
		listener, err := net.ListenUDP("udp", udpAddresses[i])
		if err != nil {
			log.Fatal(err)
		}
		listeners[i] = listener

		fmt.Println("UDP server up and listening on port", portNums[i])

	}

	// while true
	for {
		for _, listener := range listeners {
			// wait for UDP client to connect
			handleUDPConnection(listener)
		}
	}

	for _, listener := range listeners {
		listener.Close()
	}

}
