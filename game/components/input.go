package components

type Input struct {
	Enabled bool
}

func (i Input) GetName() string {
	return "input"
}