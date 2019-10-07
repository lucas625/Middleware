package calculatorProxy

import (
	"github.com/lucas625/Middleware/my-middleware/client/distribution/requestor"
	"github.com/lucas625/Middleware/my-middleware/common/distribution/clientproxy"
	"github.com/lucas625/Middleware/my-middleware/common/utils"
)

// CalculatorProxy is a struct that holds the data need to contact the server
//
// Members:
//  Host - Holds an ip address.
//  Port - Stores the used port.
//  ID   - Identifies the process.
//
type CalculatorProxy struct {
	Host string
	Port int
	ID   int
}

// Mul is a function to multiply a number by 2.
//
// Parameters:
//  p1 - a number.
//
// Returns:
//  the result.
//
func (proxy CalculatorProxy) Mul(p1 int) int {
	rq := utils.Request{}
	inv := utils.Invocation{Host: proxy.Host, Port: proxy.Port, Request: rq}
	reqtor := requestor.Requestor{}
	result := reqtor.Invoke(inv).(int)
	return result
}

// NewCalculatorProxy is a function to instantiate a new calculator based on clientproxy.
//
// Parameters:
//  cp - the clientproxy.
//
// Returns:
//  a CalculatorProxy.
//
func NewCalculatorProxy(cp clientproxy.ClientProxy) CalculatorProxy {
	return CalculatorProxy{Host: cp.Host, Port: cp.Port, ID: cp.ID}
}
