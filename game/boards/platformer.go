package boards

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/yourname/yourgame/framework"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/entities"
	"github.com/yourname/yourgame/game/systems"
)

func NewPlatformer() ecs.Board {
	renderSystem := systems.NewRender()
	movementSystem := &systems.Movement{}
	inputSystem := &systems.Input{}
	audio.NewContext(44100)

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

		ctx := audio.CurrentContext()

		f, err := os.Open("./assets/audio/coin.ogg")
		if err != nil {
			fmt.Println("ERR:", err)
			panic("Can't open file")
		}

		s, err := vorbis.Decode(ctx, f)

		m, err := audio.NewPlayer(ctx, s)
		if err != nil {
			fmt.Println("ERR:", err)
			panic("Can't open file")
		}

		m.Play()

		coin := b.GetEntity(second)

		b.RemoveEntity(coin)
	})

	movementSystem.WhenEntityTouchesAnother("CHAR", "SPIKE", func(first uint64, second uint64) {
		fmt.Println("GAME OVER!")

		ctx := audio.CurrentContext()

		f, err := os.Open("./assets/audio/hit.ogg")
		if err != nil {
			fmt.Println("ERR:", err)
			panic("Can't open file")
		}

		s, err := vorbis.Decode(ctx, f)

		m, err := audio.NewPlayer(ctx, s)
		if err != nil {
			fmt.Println("ERR:", err)
			panic("Can't open file")
		}

		m.Play()

		player := b.GetEntity(first)

		b.RemoveEntity(player)
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
		case "SPIKE":
			e := entities.NewSpike(boardEntity)
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
