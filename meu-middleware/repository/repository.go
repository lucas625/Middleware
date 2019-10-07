package repository

import (
	"reflect"

	"github.com/lucas625/Middleware/meu-middleware/distribution/clientproxy"
)

// CheckRepository compares proxies
func CheckRepository(proxy clientproxy.ClientProxy) interface{} {
	var clientProxy interface{}

	switch proxy.TypeName {
	case reflect.TypeOf(clientproxy.ClientProxy{}).String():
		calculatorProxy := clientproxy.NewClientProxy()
		calculatorProxy.TypeName = proxy.TypeName
		calculatorProxy.Host = proxy.Host
		calculatorProxy.Port = proxy.Port
		clientProxy = calculatorProxy
	}

	return clientProxy
}
