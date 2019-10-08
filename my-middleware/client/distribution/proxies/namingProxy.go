package proxies

import (
	"errors"

	"github.com/lucas625/Middleware/my-middleware/client/distribution/requestor"
	"github.com/lucas625/Middleware/my-middleware/common/distribution/clientproxy"
	"github.com/lucas625/Middleware/my-middleware/common/utils"
)

// Server is a structure for managing a naming service.
//
// Members:
//  NS   - the naming service.
//  IP   - the ip of the server.
//  Port - port to the service.
//
type Server struct {
	IP   string
	Port int
}

// Lookup is a function to find the server of a an object.
//
// Parameters:
//  name - the name of the object.
//
// Returns:
//  the proxy of the object.
//
func (sv Server) Lookup(name string) interface{} {
	param := make([]interface{}, 1)
	param[0] = name
	rq := utils.Request{Op: "Lookup", Params: param}
	inv := utils.Invocation{Host: sv.IP, Port: sv.Port, Request: rq}
	reqtor := requestor.Requestor{}
	cp := reqtor.Invoke(inv).(clientproxy.ClientProxy)
	var result interface{}
	switch cp.TypeName {
	case "Calculator":
		result = CalculatorProxy{Host: cp.Host, Port: cp.Port, ID: cp.ID}
	default:
		utils.PrintError(errors.New("unrecognized clientproxy type"), "type of the clientproxy: "+cp.TypeName)
	}
	return result.(CalculatorProxy)
}

// InitServer is a function to locate a server.
//
// parameters:
//  none.
//
// Returns:
//  the location of the server.
//
func InitServer(ip string) Server {
	sv := Server{IP: ip, Port: 8090}
	return sv
}
