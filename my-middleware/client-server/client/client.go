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
func runExperiment(numberOfCalls int, wg *sync.WaitGroup, calc *utils.CalcValues, start int) {
	defer wg.Done()
	// getting the clientproxy
	namingServer := proxies.InitServer("localhost")
	calculator := namingServer.Lookup("Calculator").(proxies.CalculatorProxy)

	// executing
	for i := start; i < numberOfCalls; i++ {
		initialTime := time.Now() //calculating time
		result := calculator.Mul(i + start)
		endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
		fmt.Println(result)                                            // making the request
		utils.AddValue(calc, endTime)                                  // pushing to the stored values
		time.Sleep(10 * time.Millisecond)                              // setting the sleep time
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
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)
		i--
	}
}

func main() {
	numberOfCalls := 10000
	// creating the calcvalues object
	calc := utils.InitCalcValues(make([]float64, numberOfCalls, numberOfCalls))
	var wg sync.WaitGroup
	go doSomething()
	aux := numberOfCalls / 500
	for i := 0; i < aux; i++ {
		wg.Add(1)
		go runExperiment(500, &wg, &calc, (i * aux))
		wg.Wait()
	}

	// evaluating
	avrg := utils.CalcAverage(&calc)
	stdv := utils.CalcStandardDeviation(&calc, avrg)

	utils.PrintEvaluation(avrg, stdv, 8)

}
