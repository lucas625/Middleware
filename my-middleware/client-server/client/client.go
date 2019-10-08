package main

import (
	"fmt"

	"github.com/lucas625/Middleware/my-middleware/client/distribution/proxies"
)

func main() {
	// getting the clientproxy
	namingServer := proxies.InitServer("localhost")
	calculator := namingServer.Lookup("Calculator").(proxies.CalculatorProxy)
	// executing
	for i := 0; i < 100; i++ {
		fmt.Println(calculator.Mul(i))
	}
}
