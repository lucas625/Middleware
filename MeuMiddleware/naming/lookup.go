package naming

import (
	"github.com/lucas625/Middleware/MeuMiddleware/clientproxy"
	"fmt"
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
//  a boolean checking if the name was added to the NamingService.
//
func (naming *NamingService) Bind(name string, proxy clientproxy.ClientProxy) (bool) {
	_, present := naming.Repository[name]
	if present {
		return false
	}
	naming.Repository[name] = proxy
	return true
}

// Lookup is a function to get a ClientProxy from the repository.
//
// Parameters:
//  name - the name key to the ClientProxy.
//
// Returns:
//  the ClientProxy.
//
func (naming *NamingService) Lookup(name string) (clientproxy.ClientProxy) {
	cp, present := naming.Repository[name]
	if !present {
		return nil
	}
	return cp
}

// List is a function to print all names in the naming service.
//
// Parameters:
//  none
//
// Returns:
//  none
//
func (naming *NamingService) List() {
	nameList := ""
	for key := range naming.Repository {
		nameList += key + "\n"
	}
	fmt.Print(nameList)
}