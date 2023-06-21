package boards

import (
	"fmt"

	"github.com/yourname/yourgame/framework"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
	"github.com/yourname/yourgame/game/entities"
	"github.com/yourname/yourgame/game/systems"
)

func NewPlatformer() ecs.Board {
	b := ecs.Board{}

	renderSystem := systems.NewRender()
	movementSystem := systems.Movement{}
	inputSystem := systems.Input{}

	const tilemapFilepath = "assets/tilemaps/platformer.tmx"

	boardmap, err := framework.LoadBoardMapFromTilemap(tilemapFilepath)
	if err != nil {
		panic(err)
	}

	tiles := boardmap.GetAllTiles()

	for _, tile := range tiles {
		switch tile.Type {
		case "WALL":
			e := entities.NewWall(tile.X, tile.Y)

			e.SetComponent(&components.Size{
				OffsetX: 0,
				OffsetY: 0,
				W:       tile.W,
				H:       tile.H,
			})

			e.SetComponent(&components.Render{
				Spritesheet:  *tile.Spritesheet,
				EntityName:   "WALL",
				SpriteName:   tile.SpriteName,
				CurrentFrame: 0,
				FrameCount:   0,
				Z:            1,
			})

			movementSystem.AddEntity(&e)
			renderSystem.AddEntity(&e)
		default:
		}
	}

	objects := boardmap.GetAllObjects()

	for _, obj := range objects {
		switch obj.Type {
		case "CHAR1":
			e := entities.NewCharacter(10, 10)

			e.SetComponent(&components.Size{
				OffsetX: 16,
				OffsetY: 6,
				W:       6,
				H:       18,
			})

			e.SetComponent(&components.Render{
				Spritesheet:  *obj.Spritesheet,
				EntityName:   "CHAR1",
				SpriteName:   "0",
				CurrentFrame: 0,
				FrameCount:   0,
				Z:            2,
			})

			movementSystem.AddEntity(&e)
			inputSystem.AddEntity(&e)
			renderSystem.AddEntity(&e)
		default:
		}
	}

	fmt.Println(objects)

	b.SetRenderer(renderSystem)
	b.AddSystem(movementSystem)
	b.AddSystem(inputSystem)

	return b
}
