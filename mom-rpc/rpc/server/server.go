package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/lucas625/Middleware/utils"
	"github.com/lucas625/Middleware/mom-rpc/rpc/multiplicador/impl"
)

func main() {

	// create new instance of multiplicator
	multiplicador := new(impl.MultiplicadorRPC)

	// create new rpc server
	server := rpc.NewServer()
	server.RegisterName("Multiplicador", multiplicador)

	// associate a http handler to servidor
	server.HandleHTTP("/", "/debug")

	// create tcp listen
	l, err := net.Listen("tcp", ":"+strconv.Itoa(8080))
	utils.PrintError(err, "Servidor não inicializado")

	// wait for calls
	fmt.Println("Servidor pronto (RPC-HTTP) ...\n")
	http.Serve(l, nil)
}
