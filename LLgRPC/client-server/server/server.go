package main

import (
	"fmt"

	"github.com/lucas625/Middleware/LLgRPC/common/distribution/absoluteobjectreference"
	"github.com/lucas625/Middleware/LLgRPC/common/distribution/clientproxy"
	"github.com/lucas625/Middleware/LLgRPC/common/distribution/namingproxy"
	"github.com/lucas625/Middleware/LLgRPC/server/distribution/invoker"
)

func main() {
	// registering the calculator
	var cp clientproxy.ClientProxy
	cp = clientproxy.InitClientProxy(absoluteobjectreference.InitAOR("localhost", 8080, 1, "tcp", 1), "Manager")
	nclient := namingproxy.InitServer(cp.AOR.IP)
	nclient.Bind("Manager", cp)
	fmt.Println("Manager registered!")
	// control loop passed to middleware
	fmt.Println("Database Server running!!")
	manInvoker := invoker.ManagerInvoker{}

	go manInvoker.Invoke()
	c := make(chan int)
	<-c
}
