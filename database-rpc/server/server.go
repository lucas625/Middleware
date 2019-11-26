package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"

	"github.com/lucas625/Middleware/LLgRPC/common/utils"
	"github.com/lucas625/Middleware/database-rpc/service/database"
	"github.com/lucas625/Middleware/database-rpc/service/manager"
)

func main() {
	db := database.InitDatabase()
	// create new instance of database
	manager1 := manager.InitManager(db)
	
	// create new rpc server
	server := rpc.NewServer()
	server.RegisterName("Manager", manager1)
	

	// create tcp listen
	l, err := net.Listen("tcp", ":"+strconv.Itoa(8080))
	server.Accept(l)
	utils.PrintError(err, "Servidor não inicializado")

	// wait for calls
	fmt.Println("Servidor pronto (RPC-HTTP)")
	
}
