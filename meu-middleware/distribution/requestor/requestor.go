package requestor

import (
	"github.com/lucas625/Middleware/meu-middleware/distribution/marshaller"
	"github.com/lucas625/Middleware/meu-middleware/infrastructure/crh"
	"github.com/lucas625/Middleware/meu-middleware/distribution/miop"
	"github.com/lucas625/Middleware/utils"
	"fmt"
)

// Requestor is a structure to enable Requestor funcions
//
// Members:
//  none
//
type Requestor struct{}

// Invoke is a funcion that receives a Invocation and returns a Interface based on the Invocation parameters
//
// Parameters:
// inv - Received invocation
//
// Returns:
// interface
//
func (Requestor) Invoke(inv utils.Invocation) interface{} {
	marshallerInst := marshaller.Marshaller{}
	crhInst := crh.CRH{ServerHost:inv.Host,ServerPort:inv.Port}

	// create request packet
	reqHeader := miop.RequestHeader{Context:"Context",RequestID:1000,ResponseExpected:true,ObjectKey:2000,Operation:inv.Request.Op}
	reqBody := miop.RequestBody{Body:inv.Request.Params}
	header := miop.Header{Magic:"MIOP",Version:"1.0",ByteOrder:true,MessageType:1} // MessageType = 1 == Request
	body := miop.Body{ReqHeader:reqHeader,ReqBody:reqBody}
	miopPacketRequest := miop.Packet{Hdr:header,Bd:body}

	// serialise request packet
	msgToClientBytes := marshallerInst.Marshall(miopPacketRequest)
	
	// send request packet and receive reply packet
	msgFromServerBytes := crhInst.SendReceive(msgToClientBytes)
	fmt.Println("oi")
	miopPacketReply := marshallerInst.Unmarshall(msgFromServerBytes)
	
	// extract result from reply packet
	r := miopPacketReply.Bd.RepBody.OperationResult

	return r
}




