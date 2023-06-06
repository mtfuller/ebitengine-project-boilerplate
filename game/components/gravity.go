package components

type Gravity struct {
	Enabled bool
}

func (g Gravity) GetName() string {
	return "gravity"
}