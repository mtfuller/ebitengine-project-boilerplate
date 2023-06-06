package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	stateManager	*StateManager
}

func (g *Game) SetStateManager(s *StateManager) {
	g.stateManager = s
}

func (g *Game) Update() error {
	return g.stateManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stateManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}