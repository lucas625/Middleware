package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
)

func handleTCPConnection(listener *net.TCPListener) {
	// here is where you want to do stuff like read or write to client
	conn, err := listener.Accept()
	if err != nil {
		log.Print(err)
	} else {
		defer conn.Close()
		fmt.Println("TCP client : ", conn.RemoteAddr().String())
		for {
			// will listen for message to process ending in newline (\n)
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println(err)
				break
			}
			// output message received
			fmt.Println("Received from TCP client:", string(message))
			// sample process for string received
			val, err := strconv.Atoi(string(message[:len(message)-1]))
			newmessage := strconv.Itoa(0)
			if err != nil {
				log.Println(err)
			} else {
				newmessage = strconv.FormatBool((val%2 == 0)) + " " + strconv.Itoa(val)
			}

			// send new string back to client
			conn.Write([]byte(newmessage + "\n"))
		}
	}
	handleTCPConnection(listener)

}

func waitHandleConn(listener *net.TCPListener, wg *sync.WaitGroup) {
	defer wg.Done() // tells one of the goroutines ended after this func ends.
	defer listener.Close()
	handleTCPConnection(listener)
}

func main() {
	hostName := "localhost"
	// 5 clients allowed at the same time
	numberOfClients := 5
	portNums := make([]string, numberOfClients)
	services := make([]string, len(portNums))
	tcpAddresses := make([]*net.TCPAddr, len(portNums))
	listeners := make([]*net.TCPListener, len(portNums))

	// filling the arrays
	for i, _ := range portNums {
		portNums[i] = strconv.Itoa(8080 + i)
		services[i] = hostName + ":" + portNums[i]

		// resolving the address
		tcpAdd, err := net.ResolveTCPAddr("tcp", services[i])
		if err != nil {
			log.Fatal(err)
		}
		tcpAddresses[i] = tcpAdd

		// listener for UDP connection
		listener, err := net.ListenTCP("tcp", tcpAddresses[i])
		if err != nil {
			log.Fatalln(err)
		}
		listeners[i] = listener

		fmt.Println("TCP server up and listening on port", portNums[i])

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
