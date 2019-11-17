package clientproxy

import (
	"github.com/lucas625/Middleware/LLgRPC/common/distribution/absoluteobjectreference"
)

// ClientProxy is a struct that holds the data need to contact the server.
//
// Members:
//  AOR      - Holds the absolute object reference.
//  TypeName - Declares the type used.
//
type ClientProxy struct {
	AOR      absoluteobjectreference.AOR
	TypeName string
}

// InitClientProxy is a function to initialize a client proxy.
//
// Parameters:
//  aor      - Absolute object reference.
//  typename - Declares the type used.
//
// Returns:
//  a client proxy
//
func InitClientProxy(aor absoluteobjectreference.AOR, typename string) ClientProxy {
	var cp ClientProxy
	cp.AOR = aor
	cp.TypeName = typename
	return cp
}
