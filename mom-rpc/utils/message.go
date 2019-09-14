package utils

// Message is a structure to encode to json.
//
// Members:
//  text - the data it contains.
//
type Message struct {
	Text string
}

// DecodeMessage is a function to return a LoremIpsum string.
//
// Parameters:
//  option - the key to the answer.
//
// Returns:
//  the corresponding string.
//
func DecodeMessage(option *Message) string {
	switch option.Text {
	case "Bom dia":
		return "Boa noite!\n"
	case "Obrigado":
		return "De nada!\n"
	}
	return "Mensagem desconhecida.\n"
}
