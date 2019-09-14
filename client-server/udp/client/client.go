package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// class for calc time
// remember, only one client will take care of this.
type CalcTimes struct {
	Used   int
	Values []float64
}

func standardDeviation(calc *CalcTimes, average float64) float64 {
	var sd float64
	for i := 0; i < calc.Used; i++ {
		sd += math.Pow(calc.Values[i]-average, 2)
	}
	sd = math.Sqrt(sd / float64(calc.Used))
	return sd
}

func calcAverage(calc *CalcTimes) float64 {
	// calculates the average value of a CalcTimes struct
	var total float64
	for i := 0; i < calc.Used; i++ {
		total += calc.Values[i]
	}
	return total / float64(calc.Used)
}

func addTime(calc *CalcTimes, value float64) {
	calc.Values[calc.Used] = value
	calc.Used++
}

// udp client
func UdpClient(address string, wg *sync.WaitGroup, numberOfCalls int, calc *CalcTimes, count bool) {
	defer wg.Done()
	// gets the ip and port in the format: ip:port.
	ipPort, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalln(err)
	}
	// connecting to server
	connection, err := net.DialUDP("udp", nil, ipPort)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("connected to server: %s.\n", connection.RemoteAddr().String())

	// cleaning buffers
	err = connection.SetWriteBuffer(0)
	if err != nil {
		log.Print(err)
	}
	err = connection.SetReadBuffer(0)
	if err != nil {
		log.Print(err)
	}

	// setting buffers
	err = connection.SetWriteBuffer(64 * 1024 * numberOfCalls)
	if err != nil {
		log.Print(err)
	}
	err = connection.SetReadBuffer(64 * 1024 * numberOfCalls)
	if err != nil {
		log.Print(err)
	}
	// closes the connection once the function ends
	defer connection.Close()

	var wgCalls sync.WaitGroup

	for i := 0; i < numberOfCalls; i++ {

		// write a message to server
		message := []byte(strconv.Itoa(i)) //sending time to get the right rtt
		initialTime := time.Now()
		_, err = connection.Write(message)

		if err != nil {
			log.Fatalln(err)
		}
		wgCalls.Add(1)
		go func() {
			defer wgCalls.Done()
			// time.Now().Add uses nanoseconds
			deadline := time.Now().Add(35000000000) // 35s
			err = connection.SetReadDeadline(deadline)
			// receive message from server
			buffer := make([]byte, 4096)
			n, _, err := connection.ReadFromUDP(buffer)
			if err == nil {
				endTime := float64(time.Now().Nanosecond()-initialTime.Nanosecond()) / 1000000
				aux := strings.Split(string(buffer[:n]), " ")
				bol := aux[0]
				val := aux[1]
				// finding the rtt
				if count { // the question only asks for one client to be calculated

					if endTime >= 0 {
						addTime(calc, endTime)
						fmt.Printf("The RTT took: %0.2fms.\n", endTime)
					}
				}

				fmt.Printf("Received from UDP server: %s is an even number = %s.\n", val, bol)
			}

			buffer = nil

		}()

	}
	wgCalls.Wait()
}

func main() {
	server := "localhost"
	numberOfClients := 1
	services := make([]string, numberOfClients)

	numberOfCalls := 10000

	tCalc := CalcTimes{Used: 0, Values: make([]float64, numberOfCalls*numberOfClients)}
	var wg sync.WaitGroup

	for i := 0; i < numberOfClients; i++ {
		services[i] = server + ":" + strconv.Itoa(8080+i)
		wg.Add(1)
		// only the first client will count
		if i == 0 {
			go UdpClient(services[i], &wg, numberOfCalls, &tCalc, true)
		} else {
			go UdpClient(services[i], &wg, numberOfCalls, &tCalc, false)
		}
	}

	wg.Wait()
	average := calcAverage(&tCalc)
	standardDev := standardDeviation(&tCalc, average)
	fmt.Printf("The Average RTT was: %0.4fms.\n", average)
	fmt.Printf("The Standard Deviation on the RTT was: %0.4fms.\n", standardDev)
	fmt.Println("Successful operations: ", tCalc.Used)
}
