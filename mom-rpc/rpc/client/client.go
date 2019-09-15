package main

import (
	"fmt"
	"net/rpc"
	"strconv"
	"time"

	"github.com/lucas625/Middleware/utils"
)

func main() {
	numberOfCalls := 10000 // the number of server calls

	calc := utils.InitCalcValues(make([]float64, numberOfCalls, numberOfCalls)) // object to store the rtts

	var reply int
	// connect to servidor
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(8080))
	utils.PrintError(err, "O Servidor não está pronto")

	// make requests
	//fmt.Println("Client started execution... ")
	//start := time.Now()
	for i := 0; i < numberOfCalls; i++ {
		initialTime := time.Now()

		// prepara request
		args := i

		// envia request e recebe resposta
		client.Call("Multiplicador.Mul", args, &reply)

		endTime := float64(time.Now().Sub(initialTime).Milliseconds()) // RTT
		utils.AddValue(&calc, endTime)

		fmt.Printf("%v\n", *&reply)

		//time.Sleep(25 * time.Millisecond)
	}
	// evaluating
	avrg := utils.CalcAverage(&calc)
	stdv := utils.CalcStandardDeviation(&calc, avrg)

	utils.PrintEvaluation(avrg, stdv)
}
