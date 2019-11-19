package main

import (
	"github.com/lucas625/Middleware/LLgRPC/naming-service/proxies"
)

func main() {
	// setting the naming server on
	namingServer := proxies.InitServer()
	go namingServer.Run()

	c := make(chan int)
	<-c
}
