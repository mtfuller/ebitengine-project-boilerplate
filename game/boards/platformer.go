package boards

import (
	"github.com/yourname/yourgame/framework"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/entities"
	"github.com/yourname/yourgame/game/systems"
)

func NewPlatformer() ecs.Board {
	b := ecs.Board{}

	renderSystem := systems.NewRender()
	movementSystem := &systems.Movement{}
	inputSystem := &systems.Input{}

	b.SetRenderer(renderSystem)
	b.AddSystem(movementSystem)
	b.AddSystem(inputSystem)

	const tilemapFilepath = "assets/tilemaps/platformer.tmx"

	boardmap, err := framework.LoadBoardMapFromTilemap(tilemapFilepath)
	if err != nil {
		panic(err)
	}

	boardEntities := boardmap.GetAllEntities()

	for _, boardEntity := range boardEntities {
		switch boardEntity.Type {
		case "WALL":
			e := entities.NewWall(boardEntity)
			movementSystem.AddEntity(&e)
			renderSystem.AddEntity(&e)
		case "CHAR1":
			e := entities.NewCharacter(boardEntity)
			movementSystem.AddEntity(&e)
			inputSystem.AddEntity(&e)
			renderSystem.AddEntity(&e)
		default:
		}
	}

	return b
}
