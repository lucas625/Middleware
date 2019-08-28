package main

import (
	"bufio"
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
func TcpClient(address string, wg *sync.WaitGroup, numberOfCalls int, calc *CalcTimes, count bool) {
	defer wg.Done()
	// gets the ip and port in the format: ip:port.
	ipPort, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	// connecting to server
	connection, err := net.Dial("tcp", ipPort.String())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("connected to server: %s.\n", connection.RemoteAddr().String())

	// closes the connection once the function ends
	defer connection.Close()

	for i := 0; i < numberOfCalls; i++ {

		initialTime := time.Now()
		fmt.Fprintf(connection, strconv.Itoa(i)+"\n")
		// listen for reply
		message, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		aux := strings.Split(string(message[:len(message)-1]), " ")
		bol := aux[0]
		val := aux[1]
		fmt.Printf("Received from TCP server: %s is an even number = %s.\n", val, bol)
		if count { // the question only asks for one client to be calculated
			endTime := float64(time.Now().Nanosecond()-initialTime.Nanosecond()) / 1000000
			if endTime > 0 {
				addTime(calc, endTime)
				fmt.Printf("The RTT took: %0.2fms.\n", endTime)
			}
		}
	}
}

func main() {
	server := "localhost"
	numberOfClients := 5
	services := make([]string, numberOfClients)

	numberOfCalls := 10000

	tCalc := CalcTimes{Used: 0, Values: make([]float64, numberOfCalls*numberOfClients)}
	var wg sync.WaitGroup

	for i := 0; i < numberOfClients; i++ {
		services[i] = server + ":" + strconv.Itoa(8080+i)
		wg.Add(1)
		// only the first client will count
		if i == 0 {
			go TcpClient(services[i], &wg, numberOfCalls, &tCalc, true)
		} else {
			go TcpClient(services[i], &wg, numberOfCalls, &tCalc, false)
		}
	}

	wg.Wait()
	average := calcAverage(&tCalc)
	standardDev := standardDeviation(&tCalc, average)
	fmt.Printf("The Average RTT was: %0.4fms.\n", average)
	fmt.Printf("The Standard Deviation on the RTT was: %0.4fms.\n", standardDev)
	fmt.Println("Successful operations: ", tCalc.Used)
}
