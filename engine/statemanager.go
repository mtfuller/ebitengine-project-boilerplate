package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yourname/yourgame/engine/ecs"
)

type StateManager struct {
	boards    		map[*ecs.Board]struct{}
	currentBoard	*ecs.Board
}

func (s *StateManager) AddBoard(b *ecs.Board) {
	if s.boards == nil {
		s.boards = make(map[*ecs.Board]struct{})
	}

	s.boards[b] = struct{}{}

	if s.currentBoard == nil {
		s.currentBoard = b
	}
}

func (s *StateManager) Update() error {
	s.currentBoard.Update()

	return nil
}

func (s *StateManager) Draw(screen *ebiten.Image) {
	s.currentBoard.Draw(screen)
}