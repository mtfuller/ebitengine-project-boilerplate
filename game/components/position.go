package components

type Position struct {
	X, Y float64
}

func (p Position) GetName() string {
	return "position"
}