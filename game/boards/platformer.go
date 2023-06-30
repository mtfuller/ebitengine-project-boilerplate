package boards

import (
	"fmt"

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
			movementSystem.HandleEntityDestoryed(e)
			renderSystem.HandleEntityDestoryed(e)
		},
	}

	movementSystem.WhenEntityTouchesAnother("CHAR", "COIN", func(first uint64, second uint64) {
		fmt.Println("SCORE!")

		coin := b.GetEntity(second)

		b.RemoveEntity(coin)
	})

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
		case "COIN":
			e := entities.NewCoin(boardEntity)
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
