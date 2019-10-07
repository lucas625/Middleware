package clientproxy

// ClientProxy is a struct that holds the data need to contact the server
//
// Members:
//  Host     - Holds an ip address.
//  Port     - Stores the used port.
//  ID       - Identifies the process.
//  TypeName - Declares the type used.
//
type ClientProxy struct {
	Host     string
	Port     int
	ID       int
	TypeName string
}
