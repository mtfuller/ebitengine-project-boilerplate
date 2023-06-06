package main

import (
	_ "image/png"
	"fmt"
	"log"
	"github.com/yourname/yourgame/engine"
	"github.com/yourname/yourgame/game/boards"
	//"github.com/yourname/yourgame/game/entities"
	//"github.com/yourname/yourgame/game/systems"
	"github.com/yourname/yourgame/game/loaders"
	"github.com/hajimehoshi/ebiten/v2"
)

var game engine.Game;

func init() {
	loaders.Load()

	board := boards.NewPlatformer()

	gsm := engine.StateManager{}

	gsm.AddBoard(&board)

	game.SetStateManager(&gsm)

	fmt.Println(board)

	// renderSystem := systems.NewRender()
	// movementSystem := systems.Movement{}
	// inputSystem := systems.Input{}

	// player := entities.NewCharacter()
	// wall1 := entities.NewWall(100, 150)
	// wall2 := entities.NewWall(132, 150)
	// wall3 := entities.NewWall(164, 150)
	// wall4 := entities.NewWall(196, 150)

	// renderSystem.AddEntity(&player)
	// renderSystem.AddEntity(&wall1)
	// renderSystem.AddEntity(&wall2)
	// renderSystem.AddEntity(&wall3)
	// renderSystem.AddEntity(&wall4)

	// movementSystem.AddEntity(&player)
	// movementSystem.AddEntity(&wall1)
	// movementSystem.AddEntity(&wall2)
	// movementSystem.AddEntity(&wall3)
	// movementSystem.AddEntity(&wall4)

	// inputSystem.AddEntity(&player)
	
	// game.SetRenderer(renderSystem)
	// game.AddSystem(movementSystem)
	// game.AddSystem(inputSystem)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Testing")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}