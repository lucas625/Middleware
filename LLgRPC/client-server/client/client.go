package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/lucas625/Middleware/LLgRPC/client/distribution/proxies"
	"github.com/lucas625/Middleware/LLgRPC/common/distribution/namingproxy"
	"github.com/lucas625/Middleware/LLgRPC/common/service/person"
	"github.com/lucas625/Middleware/LLgRPC/common/utils"
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
	namingServer := namingproxy.InitServer("localhost")
	manager := namingServer.Lookup("Manager").(proxies.ManagerProxy)
	// executing
	for i := 0; i < numberOfCalls; i++ {
		initialTime := time.Now() //calculating time
		p := person.InitPerson("lucas", 10, "M", 1)
		result := manager.AddPerson(*p)
		endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
		fmt.Println(result)                                            // making the request
		utils.AddValue(calc, endTime)                                  // pushing to the stored values
		time.Sleep(10 * time.Millisecond)                              // setting the sleep time
	}

	if start >= 40 {
		manager.Write("files")
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
	// creating the calcvalues object
	calc := utils.InitCalcValues(make([]float64, numberOfCalls, numberOfCalls))
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
