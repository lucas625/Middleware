package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	//"github.com/lucas625/Middleware/utils"
	//"github.com/lucas625/Middleware/mom-rpc/rpc/multiplicador/impl"
	//"github.com/lucas625/Middleware/LLgRPC/server/service/database"
	//"github.com/lucas625/Middleware/LLgRPC/server/service/manager"
	"github.com/lucas625/Middleware/database-rpc/service/manager"
)

func main() {

	// create new instance of database
	manager := new(manager.Manager)

	// create new rpc server
	server := rpc.NewServer()
	server.RegisterName("Database", manager)

	// associate a http handler to servidor
	server.HandleHTTP("/", "/debug")

	// create tcp listen
	l, err := net.Listen("tcp", ":"+strconv.Itoa(8080))
	utils.PrintError(err, "Servidor não inicializado")

	// wait for calls
	fmt.Println("Servidor pronto (RPC-HTTP) ...\n")
	http.Serve(l, nil)
}
