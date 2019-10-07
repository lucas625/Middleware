package calculatorProxy

// CalculatorProxy is a structure for being the proxy of a calculator.
//
// Members:
//  none
//
type CalculatorProxy struct{}

// Mul is a function to multiply a number by 2.
//
// Parameters:
//  p1 - a number.
//
// Returns:
//  th result.
//
func (proxy CalculatorProxy) Mul(p1 int) int {
	return 2
}

// NewCalculatorProxy is a function to instantiate a new calculator.
//
// Parameters:
//  none
//
// Returns:
//  a CalculatorProxy.
func NewCalculatorProxy() CalculatorProxy {
	return CalculatorProxy{}
}
