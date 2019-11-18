package main

import (
	"github.com/lucas625/Middleware/LLgRPC/naming-server/distribution/proxies"
)

func main() {
	// setting the naming server on
	namingServer := proxies.InitServer()
	namingServer.Run()
}
