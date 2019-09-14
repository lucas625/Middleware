package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
)

func handleUDPConnection(conn *net.UDPConn) {
	// here is where you want to do stuff like read or write to client

	buffer := make([]byte, 4096)

	// cleaning buffers
	err := conn.SetWriteBuffer(0)
	if err != nil {
		log.Print(err)
	}
	err = conn.SetReadBuffer(0)
	if err != nil {
		log.Print(err)
	}
	// setting buffers
	err = conn.SetWriteBuffer(64 * 1024 * 100000)
	if err != nil {
		log.Print(err)
	}
	err = conn.SetReadBuffer(64 * 1024 * 100000)
	if err != nil {
		log.Print(err)
	}

	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatalln(err)
	}

	val, err := strconv.Atoi(string(buffer[:n]))

	fmt.Println("UDP client : ", addr)
	fmt.Printf("Received from UDP client : %d.\n", val)

	// write message back to client
	message := []byte(strconv.FormatBool((val%2 == 0)) + " " + strconv.Itoa(val))
	_, err = conn.WriteToUDP(message, addr)

	if err != nil {
		log.Println(err)
	}
	// while true
	buffer = nil
	handleUDPConnection(conn)

}

func waitHandleConn(conn *net.UDPConn, wg *sync.WaitGroup) {
	defer wg.Done() // tells one of the goroutines ended after this func ends.
	defer conn.Close()
	handleUDPConnection(conn)
}

func main() {
	hostName := "localhost"
	// 5 clients allowed at the same time
	numberOfClients := 5
	portNums := make([]string, numberOfClients)
	services := make([]string, len(portNums))
	udpAddresses := make([]*net.UDPAddr, len(portNums))
	listeners := make([]*net.UDPConn, len(portNums))

	// filling the arrays
	for i, _ := range portNums {
		portNums[i] = strconv.Itoa(8080 + i)
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
			log.Fatalln(err)
		}
		listeners[i] = listener

		fmt.Println("UDP server up and listening on port", portNums[i])

	}

	// sync
	var wg sync.WaitGroup

	for _, listener := range listeners {
		wg.Add(1) // tells the wait group to wait for one more goroutine
		go waitHandleConn(listener, &wg)
	}

	// waiting for all goroutines
	wg.Wait()
	fmt.Println("Server terminated!")

}
