package components

type Velocity struct {
	VX, VY float64
}

func (v Velocity) GetName() string {
	return "velocity"
}