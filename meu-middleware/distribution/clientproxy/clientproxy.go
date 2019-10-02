package clientproxy

// ClientProxy is a struct that holds the data need to contact the server
//
// Objects:
// Host - Holds a ip address
// Port - Stores the port usd
// Id - Identifies the client
// TypeName - Declares the type used

type ClientProxy struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}