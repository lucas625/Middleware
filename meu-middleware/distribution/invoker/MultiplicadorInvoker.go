package invoker

import (
	"github.com/lucas625/Middleware/meu-middleware/distribution/marshaller"
	"github.com/lucas625/Middleware/meu-middleware/infrastructure/srh"
	"github.com/lucas625/Middleware/meu-middleware/distribution/miop"
	"github.com/lucas625/Middleware/meu-middleware/distribution/lifecycle-management/pooling"
	"github.com/lucas625/Middleware/meu-middleware/multiplicador/impl"
)

// MultiplicadorInvoker is a structure to run the invoker.
//
// Members:
//  none
//
type MultiplicadorInvoker struct {}

// NewMultiplicadorInvoker is a funcion to initialize a MultiplicadorInvoker.
//
// Parameters:
//  none
//
// Returns:
//  the MultiplicadorInvoker
//
func NewMultiplicadorInvoker() MultiplicadorInvoker {
	p := new(MultiplicadorInvoker)
	return *p
}

// Invoke is a funcion to set the server running.
//
// Parameters:
//  none
//
// Returns:
//  none
//
func (MultiplicadorInvoker) Invoke (){
	marshallerImpl := marshaller.Marshaller{}
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)

	// creating the pool
	multiplicadorList := make([]interface{}, 11)
	for i := 0; i < len(multiplicadorList); i++ {
		multiplicadorList[i] = impl.Multiplicador{}
	}
	multPool := pooling.InitPool(multiplicadorList)
	defer pooling.EndPool(multPool)

	for {
		print("Server invoking")
		srhImpl := srh.SRH{ServerHost:"localhost",ServerPort:8080}
		
		// Receive data
		rcvMsgBytes := srhImpl.Receive()

		// 	unmarshall
		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)
		
		
		// setup request
		_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
		var multA *impl.Multiplicador
		multA = multPool.GetFromPool().(*impl.Multiplicador)

		// finding the operation
		operation := miopPacketRequest.Bd.ReqHeader.Operation
		switch operation {
		case "Mul":
			replParams[0] = multA.Mul(_p1)
		}

		// assembly packet
		repHeader := miop.ReplyHeader{Context:"", RequestID: miopPacketRequest.Bd.ReqHeader.RequestID, Status:1}
		repBody   := miop.ReplyBody{OperationResult: replParams}
		header    := miop.Header{Magic:"MIOP", Version:"1.0", ByteOrder:true, MessageType:1} // MessageType 1 = request
		body      := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		// marshall reply
		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		// send Reply
		srhImpl.Send(msgToClientBytes)
	}
}