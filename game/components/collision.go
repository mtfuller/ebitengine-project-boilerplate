package components

type Collision struct {
	Enabled bool
	Solid   bool
}

func (c Collision) GetName() string {
	return "collision"
}
