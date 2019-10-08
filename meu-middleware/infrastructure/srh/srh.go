package srh

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

// SRH is a structure for Server to Client setups.
//
// Members:
//  ServerHost - server host IP.
//  ServerPort - server host port.
//
type SRH struct {
	ServerHost string
	ServerPort int
}

var ln net.Listener
var conn net.Conn
var err error

// Receive is a funcion that listens for messages
//
// Parameters:
// none
//
// Returns:
// Message received from client
//
func (srh SRH) Receive() []byte {

	// create listener
	ln, err = net.Listen("tcp", srh.ServerHost+":"+strconv.Itoa(srh.ServerPort))
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// accept connections
	conn, err = ln.Accept()
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// receive message size
	size := make([]byte, 4)
	_, err = conn.Read(size)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	// receive message
	msg := make([]byte, sizeInt)
	_, err = conn.Read(msg)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	return msg
}

// Send is a funcion that sends a message to a client
//
// Parameters:
// msgToClient - A byte package to be sent to a client
//
// Returns:
// none
//
func (SRH) Send(msgToClient []byte) {

	// send message size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	conn.Write(size)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// send message
	_, err = conn.Write(msgToClient)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// close connection
	conn.Close()
	ln.Close()
}
