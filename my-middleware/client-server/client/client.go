package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/lucas625/Middleware/my-middleware/client/distribution/proxies"
	"github.com/lucas625/Middleware/my-middleware/common/utils"
)

// runExperiment is a function to run the experiment.
//
// Parameters:
//  numberOfCalls - the total number of calls.
//
// Returns:
//  none
//
func runExperiment(numberOfCalls int, wg *sync.WaitGroup, start int) {
	defer wg.Done()
	// getting the clientproxy
	namingServer := proxies.InitServer("localhost")
	calculator := namingServer.Lookup("Calculator").(proxies.CalculatorProxy)
	// creating the calcvalues object
	calc := utils.InitCalcValues(make([]float64, numberOfCalls, numberOfCalls))
	// executing
	for i := start; i < numberOfCalls; i++ {
		initialTime := time.Now()                                      //calculating time
		fmt.Println(calculator.Mul(i))                                 // making the request
		endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
		utils.AddValue(&calc, endTime)                                 // pushing to the stored values
		//time.Sleep(10 * time.Millisecond)                              // setting the sleep time
	}
	// evaluating
	avrg := utils.CalcAverage(&calc)
	stdv := utils.CalcStandardDeviation(&calc, avrg)

	utils.PrintEvaluation(avrg, stdv, 8)
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
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		i--
	}
}

func main() {
	numberOfCalls := 50
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go runExperiment(((i + 1) * numberOfCalls), &wg, (i * numberOfCalls))
		wg.Wait()
	}
	go doSomething()
}
