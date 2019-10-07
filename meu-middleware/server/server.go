package main

import (
	"fmt"
	"github.com/lucas625/Middleware/meu-middleware/distribution/clientproxy"
	//"calculadora/mymiddleware/client/proxies"
	"github.com/lucas625/Middleware/meu-middleware/naming/proxy"
	//"mymiddleware/services/naming/proxy"
	"github.com/lucas625/Middleware/meu-middleware/distribution/invoker"
	//"calculadora/mymiddleware/server/invoker"
)

func main() {

	// create a built-in proxy of naming service
	namingProxy := proxy.NamingProxy{}

	// create a proxy of calculator service
	multiplicador := clientproxy.NewClientProxy()
	//converter := proxies.NewConverterProxy()
	fmt.Println(3)
	// register service in the naming service
	namingProxy.Register("Multiplicador", multiplicador)
	//namingProxy.Register("Converter", converter)
	fmt.Println(4)
	// control loop passed to middleware
	fmt.Println("Multiplicator Server running!!")
	multiplicadorInvoker := invoker.NewMultiplicadorInvoker()
	//converterInvoker := invoker.NewConverter()

	go multiplicadorInvoker.Invoke()
	//go converterInvoker.Invoke()

	fmt.Scanln()
}

