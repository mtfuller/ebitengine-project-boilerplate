package components

import (
	"github.com/yourname/yourgame/engine"
)

type Render struct {
	Spritesheet *engine.Spritesheet
	CurrentSprite string
	CurrentFrame int
	MaxCountdown int
	Countdown int
	Z int
}

func (r Render) GetName() string {
	return "render"
}