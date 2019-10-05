package impl

type Multiplicador struct{}

func (Multiplicador) Mul(x int) int {
	return x * 2
}
