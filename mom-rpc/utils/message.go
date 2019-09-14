package utils

// Message is a structure to encode to json.
//
// Members:
//  Client - the index of the client.
//  Value  - the data it contains.
//
type Message struct {
	Client int
	Value  int
}

// DecodeMessage is a function to return a LoremIpsum string.
//
// Parameters:
//  msg - the Message.
//
// Returns:
//  the corresponding string.
//
func DecodeMessage(msg *Message) int {
	return 2 * msg.Value
}
