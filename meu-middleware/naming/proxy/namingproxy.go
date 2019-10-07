package proxy

import (
	"github.com/lucas625/Middleware/meu-middleware/distribution/requestor"
	"github.com/lucas625/Middleware/utils"
	"github.com/lucas625/Middleware/meu-middleware/repository"
	"github.com/lucas625/Middleware/meu-middleware/distribution/clientproxy"

	"fmt"
)

// NamingProxy is a struct for naming
//
// Members:
//  none
//
type NamingProxy struct{}

// Register is a function to register on the naming proxy.
//
// Parameters:
//  p1    - the key.
//  proxy - the proxy.
//
// Returns:
//  a boolean checking if was ok.
//
func (NamingProxy) Register(p1 string, proxy interface{}) bool {

	// prepare invocation
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = proxy
	namingproxy := clientproxy.ClientProxy{Host:"",Port:8081,ID:0}
	request := utils.Request{Op: "Register", Params: params}
	inv := utils.Invocation{Host: namingproxy.Host, Port: namingproxy.Port,Request: request}
	fmt.Println("hue")
	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})
	fmt.Println("whay")
	return ter[0].(bool)
}

// Lookup is a function to search the name.
//
// Parameters:
//  p1 - the key.
//
// Returns:
//  the service.
//
func (NamingProxy) Lookup(p1 string) interface{} {
	// prepare invocation
	params := make([]interface{}, 1)
	params[0] = p1
	namingproxy := clientproxy.ClientProxy{Host:"",Port:8081,ID:0}
	request := utils.Request{Op: "Lookup", Params: params}
	inv := utils.Invocation{Host:namingproxy.Host,Port:namingproxy.Port,Request:request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	// process reply
	proxyTemp := ter[0].(map[string]interface{})
	clientProxyTemp := clientproxy.ClientProxy{TypeName:proxyTemp["TypeName"].(string),Host:proxyTemp["Host"].(string),Port:int(proxyTemp["Port"].(float64))}
	clientProxy := repository.CheckRepository(clientProxyTemp)

	return clientProxy
}
