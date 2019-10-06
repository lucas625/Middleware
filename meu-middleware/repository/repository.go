package repository

import (
	"reflect"
	"github.com/lucas625/Middleware/meu-middleware/distribution/clientproxy"
)
// CheckRepository compares proxies
func CheckRepository(proxy clientproxy.ClientProxy) interface{}{
	var clientProxy interface{}

	switch proxy.TypeName{
	case reflect.TypeOf(proxies.ClientProxy{}).String():
		calculatorProxy := proxies.NewClientProxy()
		calculatorProxy.Proxy.TypeName = proxy.TypeName
		calculatorProxy.Proxy.Host = proxy.Host
		calculatorProxy.Proxy.Port = proxy.Port
		clientProxy = calculatorProxy
	}

	return clientProxy
}

