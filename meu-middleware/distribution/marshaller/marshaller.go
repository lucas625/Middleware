package marshaller

import (
	"encoding/json"
	"log"
	"github.com/lucas625/Middleware/meu-middleware/distribution/miop"
)

// Marshaller is a structure to enable Marshaller funcions
//
// Members:
//  none
//
type Marshaller struct{}

// Marshall is a funcion that receives a MIOP packet and transforms it to a bytes package
//
// Parameters:
//  msg - Target packet
//
// Returns:
//  packet transformed to bytes
//
func (Marshaller) Marshall(msg miop.Packet) []byte {

	r, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Marshaller:: Marshall:: %s", err)
	}

	return r
}

// Unmarshall is a funcion that receives a MIOP packet and transforms it to a bytes package
//
// Parameters:
//  msg - Target bytes package
//
// Returns:
//  MIOT packet from the bytes
//
func (Marshaller) Unmarshall(msg []byte) miop.Packet {

	r := miop.Packet{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		log.Fatalf("Marshaller:: Unmarshall:: %s", err)
	}
	return r
}


