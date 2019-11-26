package main

import (
	"fmt"
	"net/rpc"
	"strconv"
	"sync"
	"time"

	"github.com/lucas625/Middleware/LLgRPC/common/utils"
	"github.com/lucas625/Middleware/database-rpc/common/service/person"
)

func runExperiment(numberOfCalls int, wg *sync.WaitGroup, calc *utils.CalcValues, start int) {
	defer wg.Done()
	// connect to servidor
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(8080))
	utils.PrintError(err, "O Servidor não está pronto")

	// make requests
	for i := 0; i < numberOfCalls; i++ {
		var reply bool
		// prepara request
		args := person.InitPerson("lucas", 22, "M", 1)

		initialTime := time.Now()
		// envia request e recebe resposta
		client.Call("Manager.AddPerson", args, &reply)

		endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
		utils.AddValue(calc, endTime)

		fmt.Printf("%v\n", *&reply)

		time.Sleep(10 * time.Millisecond)
	}
	if start >= 40 {
		var reply bool
		// prepara request
		args := "outF/"
		client.Call("Manager.Write", args, &reply)
		var rep interface{}
		client.Call("Manager.List", "", &rep)
	}
}

// doSomething is a function to do some random stuff while the client is making requests.
//
// Parameters:
//  none
//
// Returns:
//  none
//
func doSomething() {
	for j := 0; j < 10; j++ {
		time.Sleep(50 * time.Millisecond)
		j--
	}
}

func main() {
	numberOfCalls := 50
	perCall := 10
	aux := numberOfCalls / perCall

	calc := utils.InitCalcValues(make([]float64, numberOfCalls, numberOfCalls)) // object to store the rtts
	var wg sync.WaitGroup
	go doSomething()

	for i := 0; i < aux; i++ {
		wg.Add(1)
		go runExperiment(perCall, &wg, &calc, (i * perCall))
		wg.Wait()
	}

	// evaluating
	avrg := utils.CalcAverage(&calc)
	stdv := utils.CalcStandardDeviation(&calc, avrg)

	utils.PrintEvaluation(avrg, stdv, 8)
}
