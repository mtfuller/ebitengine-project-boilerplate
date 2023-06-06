package components

type Collision struct {
	Enabled bool
}

func (c Collision) GetName() string {
	return "collision"
}