package impl

type MultiplicadorRPC struct{}

func (t *MultiplicadorRPC) Mul(req int, reply *int) error {
	*reply = req * 2
	return nil
}
