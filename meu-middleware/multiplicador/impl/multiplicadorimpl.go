package impl

// Multiplicador is a structure for multiplying
//
// Members:
//  none
//
type Multiplicador struct{}

// Mul is a function for multiplying a number by 2.
//
// Parameters:
//  x - the number to be multiplied.
//
// Returns:
//  the result.
//
func (Multiplicador) Mul(x int) int {
	return x * 2
}
