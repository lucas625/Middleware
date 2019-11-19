package absoluteobjectreference

// AOR is a struct that holds the absolute object reference.
//
// Members:
//  IP        - Holds an ip address.
//  Port      - Stores the used port.
//  InvokerID - ID of the invoker.
//  Protocol  - Communication protocol (In this case TCP).
//  ObjectID  - ID of the object.
//
type AOR struct {
	IP        string
	Port      int
	InvokerID int
	Protocol  string
	ObjectID  int
}

// InitAOR is a function to initialize a AOR.
//
// Parameters:
//  ip        - Holds an ip address.
//  port      - Stores the used port.
//  invokerID - ID of the invoker.
//  protocol  - Communication protocol (In this case TCP).
//  objectID  - ID of the object.
//
// Returns:
//  a AOR.
//
func InitAOR(ip string, port, invokerID int, protocol string, objectID int) AOR {
	var aor AOR
	aor.IP = ip
	aor.Port = port
	aor.InvokerID = invokerID
	aor.Protocol = protocol
	aor.ObjectID = objectID
	return aor
}
