package clientproxy

import (
	"github.com/lucas625/Middleware/meu-middleware/distribution/requestor"
	"github.com/lucas625/Middleware/utils"
	//"github.com/lucas625/Middleware/meu-middleware/identification/reference"
	"reflect"
)
// ClientProxy is a struct that holds the data need to contact the server
//
// Members:
//  Host     - Holds a ip address.
//  Port     - Stores the used port.
//  ID       - Identifies the client.
//  TypeName - Declares the type used.
//
type ClientProxy struct {
	Host     string
	Port     int
	ID       int
	TypeName string
	//AOR      reference.AbsoluteObjectReference
}

// NewClientProxy is a constructor for ClientProxy
//
func NewClientProxy() ClientProxy {
	p := new(ClientProxy)

	p.TypeName = reflect.TypeOf(ClientProxy{}).String()
	p.Host = "localhost"
	p.Port = 8080
	return *p
}

// Mul is a funcion that receives a number and returns its double
//
// Parameters:
// p1 - Number to get multiplied
//
// Returns:
// The result obtained
//
func (proxy ClientProxy) Mul (p1 int) int {

	// Sets up the necessary structs for the requestor
	params := make([]interface{},1)
	params[0] = p1
	request := utils.Request{"Mul", params}
	inv := utils.Invocation {Host:proxy.Host, Port:proxy.Port, Request:request}

	// Invokes requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return int(ter[0].(float64))
}