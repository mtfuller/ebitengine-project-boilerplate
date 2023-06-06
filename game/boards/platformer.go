package boards

import (
	"fmt"
	"github.com/yourname/yourgame/engine/ecs"
	"github.com/yourname/yourgame/game/entities"
	"github.com/yourname/yourgame/game/systems"
)

func NewPlatformer() ecs.Board {
	a := [][]uint8{
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 0, 1},
		{1, 2, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 1, 1, 1, 1, 1},
	}

	b := ecs.Board{}

	renderSystem := systems.NewRender()
	movementSystem := systems.Movement{}
	inputSystem := systems.Input{}

	for row, h := range a {
		for col, cell := range h {
			x := col * 32
			y := row * 32
			fmt.Println("(", x, ",", y, ") =", cell)

			switch cell {
			case 1:
				e := entities.NewWall(x, y)
				movementSystem.AddEntity(&e)
				renderSystem.AddEntity(&e)
			case 2:
				e := entities.NewCharacter()
				inputSystem.AddEntity(&e)
				movementSystem.AddEntity(&e)
				renderSystem.AddEntity(&e)
			}
		}
	}

	b.SetRenderer(renderSystem)
	b.AddSystem(movementSystem)
	b.AddSystem(inputSystem)

	return b
}
