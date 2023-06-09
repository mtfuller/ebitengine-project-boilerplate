package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BaseSystem struct {
}

func (b BaseSystem) GetName() string {
	return "System::BaseSystem"
}

// func (b *BaseSystem) AddEntity(e *Entity) {
// 	if b.Entities == nil {
// 		b.Entities = make(map[*Entity]struct{})
// 	}

// 	b.Entities[e] = struct{}{}
// }

type ISystem interface {
	GetName() string
	Update(entity *Entity)
}

type IRenderableSystem interface {
	Render(screen *ebiten.Image)
}
