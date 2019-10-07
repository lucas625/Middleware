package calculator

// Calculator is a structure for multiplying
//
// Members:
//  none
//
type calculator struct{}

// Mul is a function for multiplying a number by 2.
//
// Parameters:
//  x - the number to be multiplied.
//
// Returns:
//  the result.
//
func (calculator) Mul(x int) int {
	return x * 2
}
