package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/lucas625/Middleware/LLgRPC/client/distribution/proxies"
	"github.com/lucas625/Middleware/LLgRPC/common/distribution/namingproxy"
	"github.com/lucas625/Middleware/LLgRPC/common/service/person"
)

// runExperiment is a function to run the experiment.
//
// Parameters:
//  numberOfCalls - the total number of calls.
//  wg            - the wait group.
//
// Returns:
//  none
//
func runExample(numberOfCalls int, wg *sync.WaitGroup) {
	defer wg.Done()
	// getting the clientproxy
	namingServer := namingproxy.InitServer("localhost")
	manager := namingServer.Lookup("Manager").(proxies.ManagerProxy)
	manager.Write("files") // making the request
	// executing
	for i := 0; i < numberOfCalls; i++ {

		p := person.InitPerson("lucas", 10, "M", 1)
		result := manager.AddPerson(*p) // making the request
		fmt.Println(i, result)

		time.Sleep(100 * time.Millisecond) // setting the sleep time
	}

	manager.Write("files") // making the request

}

func main() {
	numberOfCalls := 10
	var wg sync.WaitGroup

	wg.Add(1)
	go runExample(numberOfCalls, &wg)
	wg.Wait()

}
