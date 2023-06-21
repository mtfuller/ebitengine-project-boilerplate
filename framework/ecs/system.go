package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BaseSystem struct {
	Entities map[*Entity]struct{}
}

func (b *BaseSystem) AddEntity(e *Entity)  {
	if b.Entities == nil {
		b.Entities = make(map[*Entity]struct{})
	}

	b.Entities[e] = struct{}{}
}

type ISystem interface {
	Update()
}

type IRenderableSystem interface {
	Render(screen *ebiten.Image)
}