package utils

// PrintCall is a function to return a LoremIpsum string.
//
// Parameters:
//  option - the key to the answer.
//
// Returns:
//  the corresponding string.
//
func PrintCall(option int) string {
	switch option {
	case 0:
		return "Bom dia!\n"
	case 1:
		return "Boa noite!\n"
	case 2:
		return "Obrigado!\n"
	case 3:
		return "De nada!\n"

	}
}

}