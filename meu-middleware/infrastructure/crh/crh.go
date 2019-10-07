package crh

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
)

// CRH is a structure for Client to Server setups.
//
// Members:
//  ServerHost - server host IP.
//  ServerPort - server host port.
//
type CRH struct {
	ServerHost string
	ServerPort int
}

// SendReceive is a funcion that receives a byte package and sends it to a server
//
// Parameters:
// msgToServer - Package to be sent
//
// Returns:
// Message received from server
//
func (crh CRH) SendReceive(msgToServer []byte) []byte {

	// connect to server
	var conn net.Conn
	var err error
	for {
		fmt.Println("aaaaaaaaaaa")
		conn, err = net.Dial("tcp", "localhost:"+strconv.Itoa(crh.ServerPort))
		if err == nil {
			//log.Fatalf("CRH:: %s", err)
			break
		}

	}

	defer conn.Close()

	// send message size
	sizeMsgToServer := make([]byte, 4)
	l := uint32(len(msgToServer))
	binary.LittleEndian.PutUint32(sizeMsgToServer, l)
	conn.Write(sizeMsgToServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// send message
	_, err = conn.Write(msgToServer)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// receive message size
	sizeMsgFromServer := make([]byte, 4)
	_, err = conn.Read(sizeMsgFromServer)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

	// receive reply
	msgFromServer := make([]byte, sizeFromServerInt)
	_, err = conn.Read(msgFromServer)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	return msgFromServer
}
