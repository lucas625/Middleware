package calculator

// Calculator is a struct for calculate values.
//
// Members:
//  none
//
type Calculator struct{}

// Mul is a function to calculate two times a number.
//
// Parameters:
//  req   - value to be multiplied.
//  reply - int pointer to hold the return
//
// Returns:
//  An error, if there are any.
//
func (t *Calculator) Mul(req int, reply *int) error {
	*reply = req * 2
	return nil
}
