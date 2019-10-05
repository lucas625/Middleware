package invoker

import (
	"github.com/lucas625/Middleware/meu-middleware/distribution/marshaller"
	"github.com/lucas625/Middleware/meu-middleware/infrastructure/srh"
	"github.com/lucas625/Middleware/meu-middleware/distribution/miop"
	"github.com/lucas625/Middleware/multiplicador/impl"
	"github.com/lucas625/Middleware/utils"
)

type MultiplicadorInvoker struct {}

func NewMultiplicadorInvoker() MultiplicadorInvoker {
	p := new(MultiplicadorInvoker)
	return p
}

func (MultiplicadorInvoker) Invoke (){
	srhImpl := srh.SRH{ServerHost:"localhost",ServerPort:8080}
	marshallerImpl := marshaller.marshaller{}
	miopPacketReply := miop.Packet{}
	replParams := make([]interface(),1)

	multiplicadorImpl := impl.Multiplicador{}

	for {
		// Receive data
		rcvMsgBytes := srhImpl.Receive()

		// 	unmarshall
		miopPacketRequest := marshallerImpl.unmarshall(rcvMsgBytes)
		operation := miopPacketRequest.Bd.ReqHeader.Operation
		
		// setup request
		_p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
		replParams[0] = multiplicadorImpl.Mul(_p1)

		// assembly packet
		repHeader := miop.ReplyHeader{Context:"", RequestID: miopPacketRequest.Bd.ReqHeader.RequestID, Status:1}
		repBody   := miop.ReplyBody{OperationResult:replParams}
		header    := miop.Header{Magic:"MIOP", Version:"1.0", ByteOrder:true, MessageType:1} // MessageType 1 = request
		body      := miop.Body{RepHeader: repHeader, RepBody: repbody}
		miopPacketReply = miopPacket{Hdr:header, Bd:body}

		// marshall reply
		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		// send Reply
		srhImpl.Send(msgToClientBytes)
	}
}