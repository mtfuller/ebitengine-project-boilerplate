package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yourname/yourgame/framework"
	"github.com/yourname/yourgame/game/boards"
)

var game framework.Game

func init() {

	board := boards.NewPlatformer()

	gsm := framework.StateManager{}

	gsm.AddBoard(&board)

	game.SetStateManager(&gsm)

}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Testing")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
