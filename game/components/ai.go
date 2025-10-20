package components

type AI struct {
	Enabled       bool
	Direction     float64
	Speed         float64
	PatrolDistance float64
	StartX        float64
}

func (a AI) GetName() string {
	return "ai"
}
