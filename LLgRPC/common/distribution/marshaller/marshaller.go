package marshaller

import (
	"encoding/json"

	"github.com/lucas625/Middleware/LLgRPC/common/distribution/packet"
	"github.com/lucas625/Middleware/LLgRPC/common/utils"
)

// Marshaller is a structure to enable Marshaller funcions
//
// Members:
//  none
//
type Marshaller struct{}

// Marshall is a funcion that receives a packet and transforms it to a bytes package
//
// Parameters:
//  msg - Target packet
//
// Returns:
//  packet transformed to bytes
//
func (Marshaller) Marshall(msg packet.Packet) []byte {

	r, err := json.Marshal(msg)
	utils.PrintError(err, "Failed to marshal.")

	return r
}

// Unmarshall is a funcion that receives a packet and transforms it to a bytes package
//
// Parameters:
//  msg - Target bytes package
//
// Returns:
//   a packet
//
func (Marshaller) Unmarshall(msg []byte) packet.Packet {

	r := packet.Packet{}
	err := json.Unmarshal(msg, &r)
	utils.PrintError(err, "Failed to unmarshal.")
	return r
}
