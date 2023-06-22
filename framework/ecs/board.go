package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	entities map[*Entity]struct{}
	systems  []ISystem
	renderer IRenderableSystem
}

func (b *Board) AddEntity(e *Entity) {
	b.entities[e] = struct{}{}
}

func (b *Board) AddSystem(s ISystem) {
	b.systems = append(b.systems, s)
}

func (b *Board) SetRenderer(r IRenderableSystem) {
	b.renderer = r
}

func (b *Board) Update() error {
	for _, s := range b.systems {
		s.Update()
	}

	return nil
}

func (b *Board) Draw(screen *ebiten.Image) {
	b.renderer.Render(screen)
}
