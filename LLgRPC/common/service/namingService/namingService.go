package namingservice

import (
	"github.com/lucas625/Middleware/LLgRPC/common/distribution/clientproxy"
)

// NamingService is a structure for holding all the names.
//
// Members:
//  Repository - a map with the name as key and the client proxy as value.
//
type NamingService struct {
	Repository map[string]clientproxy.ClientProxy
}

// Bind is a function to add a name to the repository.
//
// Parameters:
//  name  - the name to be added.
//  proxy - the ClientProxy to with the name key.
//
// Returns:
//  a flag alerting for an error.
//
func (naming *NamingService) Bind(name string, proxy clientproxy.ClientProxy) bool {
	_, present := naming.Repository[name]
	if present {
		return true
	}
	naming.Repository[name] = proxy
	return false
}

// Lookup is a function to get a ClientProxy from the repository.
//
// Parameters:
//  name - the name key to the ClientProxy.
//
// Returns:
//  the ClientProxy.
//  a flag alerting for an error.
//
func (naming *NamingService) Lookup(name string) (clientproxy.ClientProxy, bool) {
	cp, present := naming.Repository[name]
	if !present {
		var nilClientProxy clientproxy.ClientProxy // cannot return nil for struct
		return nilClientProxy, true
	}
	return cp, false
}

// Unbind is a function to remove a ClientProxy from the repository.
//
// Parameters:
//  name - the name key to the ClientProxy.
//
// Returns:
//  a flag alerting for an error.
//
func (naming *NamingService) Unbind(name string) bool {
	_, present := naming.Repository[name]
	if present {
		delete(naming.Repository, name)
		return false
	}
	return true
}

// List is a function to return all data in the naming service.
//
// Parameters:
//  none
//
// Returns:
//  all data in the naming service.
//
func (naming *NamingService) List() map[string]clientproxy.ClientProxy {
	return naming.Repository
}
