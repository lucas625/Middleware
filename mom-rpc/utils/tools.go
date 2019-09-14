package utils

import(
	"log"
)

// PrintError is a function to print an error message.
//
// Parameters:
//  err     - the error.
//	message - the string with extra information.
// 
// Returns:
//	none
func PrintError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}