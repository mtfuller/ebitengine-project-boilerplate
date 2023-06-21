package components

import (
	"github.com/yourname/yourgame/framework"
)

type Render struct {
	Spritesheet  framework.Spritesheet
	EntityName   string
	SpriteName   string
	CurrentFrame int
	FrameCount   int
	Z            int
}

func (r Render) GetName() string {
	return "render"
}
