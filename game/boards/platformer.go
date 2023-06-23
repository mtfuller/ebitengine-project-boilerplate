package boards

import (
	"github.com/yourname/yourgame/framework"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/entities"
	"github.com/yourname/yourgame/game/systems"
)

func NewPlatformer() ecs.Board {
	renderSystem := systems.NewRender()
	movementSystem := &systems.Movement{}
	inputSystem := &systems.Input{}

	b := ecs.Board{

		OnEntityCreated: func(e *ecs.Entity) {
			movementSystem.HandleEntityCreated(e)
			renderSystem.HandleEntityCreated(e)
		},

		OnEntityDestroyed: func(e *ecs.Entity) {

		},
	}

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
			b.AddEntityToSystems(&e, movementSystem.GetName(), renderSystem.GetName())
		case "CHAR1":
			e := entities.NewCharacter(boardEntity)
			b.AddEntityToSystems(&e, inputSystem.GetName(), movementSystem.GetName(),
				renderSystem.GetName())
		default:
		}
	}

	return b
}
