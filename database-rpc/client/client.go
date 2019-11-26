package main

import (
	"fmt"
	"net/rpc"
	"strconv"
	"time"

	//"sync"

	"github.com/lucas625/Middleware/LLgRPC/common/service/person"
	"github.com/lucas625/Middleware/utils"
	//"github.com/lucas625/Middleware/LLgRPC/common/utils"
)

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

/*
func runExperiment(numberOfCalls int, wg *sync.WaitGroup, calc *utils.CalcValues, start int) {
	defer wg.Done()
	// getting the clientproxy
	//namingServer := namingproxy.InitServer("localhost")
	//manager := namingServer.Lookup("Manager").(proxies.ManagerProxy)
	// executing
	for i := 0; i < numberOfCalls; i++ {
		initialTime := time.Now() //calculating time
		p := person.InitPerson("lucas", 10, "M", 1)
		client.Call("manager.AddPerson", *p) //, &reply
		//result := manager.AddPerson(*p)
		endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
		fmt.Println(i+start, result)                                   // making the request
		utils.AddValue(calc, endTime)                                  // pushing to the stored values
		time.Sleep(10 * time.Millisecond)                              // setting the sleep time
	}

	if start >= 40 {
		manager.Write("files")
		manager.List()
	}

}
*/
func main() {
	//numberOfCalls := 10000 // the number of server calls
	fmt.Println("1")
	numberOfCalls := 50
	perCall := 10
	aux := numberOfCalls / perCall

	calc := utils.InitCalcValues(make([]float64, numberOfCalls, numberOfCalls)) // object to store the rtts

	var reply int
	// connect to servidor
	fmt.Println("1")
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(8080))
	fmt.Println("1")
	utils.PrintError(err, "O Servidor não está pronto")
	//var wg sync.WaitGroup
	go doSomething()
	// make requests
	/*
		for i := 0; i < numberOfCalls; i++ {
			// prepara request
			args := i

			initialTime := time.Now()
			// envia request e recebe resposta
			client.Call("Multiplicador.Mul", args, &reply)

			endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
			utils.AddValue(&calc, endTime)

			fmt.Printf("%v\n", *&reply)

			time.Sleep(10 * time.Millisecond)
		}
	*/

	for i := 0; i < aux; i++ {
		initialTime := time.Now()
		p := person.InitPerson("lucas", 10, "M", 1)
		client.Call("manager.AddPerson", p, &reply)                    //
		endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
		utils.AddValue(&calc, endTime)
		//fmt.Printf("%v\n", *&reply)
		fmt.Println(i, *&reply)
		//wg.Add(1)
		//go runExperiment(perCall, &wg, &calc, (i * perCall))
		//wg.Wait()
		time.Sleep(10 * time.Millisecond)
	}

	// evaluating
	avrg := utils.CalcAverage(&calc)
	stdv := utils.CalcStandardDeviation(&calc, avrg)

	utils.PrintEvaluation(avrg, stdv, 8)
}
