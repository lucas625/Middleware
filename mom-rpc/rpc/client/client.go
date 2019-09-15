package main

import (
	"fmt"
	"net/rpc"
	"shared"
	"strconv"
)

func main() {
	SAMPLE_SIZE := 10000
	var reply int
	// connect to servidor
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(8080))
	shared.ChecaErro(err, "O Servidor não está pronto")

	// make requests
	//fmt.Println("Client started execution... ")
	//start := time.Now()
	for i := 0; i < SAMPLE_SIZE; i++ {

		//t1 := time.Now()

		// prepara request
		args := i

		// envia request e recebe resposta
		client.Call("Multiplicador.Mul", args, &reply)
		fmt.Printf("%v\n", *&reply)
		//t2 := time.Now()
		//x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		//fmt.Println(x)
	}
	//elapsed := time.Since(start)
	//fmt.Printf("Tempo: %s \n", elapsed)
}
