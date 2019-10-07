package main

import (
	"github.com/lucas625/Middleware/meu-middleware/naming/proxy"
	"github.com/lucas625/Middleware/meu-middleware/distribution/clientproxy"
	"fmt"
	"time"
)

// ExecuteExperiment starts the client operation
//
func ExecuteExperiment() {
	// create a built-in proxy of naming service
	namingService := proxy.NamingProxy{}

	// look for a service in naming service
	multiplicador := namingService.Lookup("Multiplicador").(clientproxy.ClientProxy)

	// invoke remote operation
	for i := 0; i < 5000; i++ {
		t1 := time.Now()
		//fmt.Print(multiplicador.Mul(i))
		multiplicador.Mul(i)
		fmt.Println(time.Now().Sub(t1))
	}

}

func main() {
	go ExecuteExperiment()

	fmt.Scanln()
}
