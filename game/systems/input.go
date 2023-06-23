package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

type Input struct {
	ecs.BaseSystem
}

func (i Input) GetName() string {
	return "System::Input"
}

func (i Input) Update(e *ecs.Entity) {
	component := e.GetComponent("input")
	input, ok := component.(*components.Input)
	if !ok {
		return
	}

	component = e.GetComponent("velocity")
	velocity, ok := component.(*components.Velocity)
	if !ok {
		return
	}

	if input.Enabled {
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) && (-0.01 < velocity.VY && velocity.VY < 0.01) {
			velocity.VY -= 5
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			velocity.VX = -1
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			velocity.VX = 1
		} else {
			velocity.VX = 0
		}
	}
}
