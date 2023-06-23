package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BoardEntity struct {
	subscribedSystems []string
}

type Board struct {
	entities          map[*Entity]BoardEntity
	systems           map[string]ISystem
	renderer          IRenderableSystem
	OnEntityCreated   func(*Entity)
	OnEntityDestroyed func(*Entity)
}

func (b *Board) AddEntity(e *Entity) {
	b.entities[e] = BoardEntity{subscribedSystems: []string{}}
}

func (b *Board) AddEntityToSystems(e *Entity, systemNames ...string) {
	if b.entities == nil {
		b.entities = make(map[*Entity]BoardEntity)
	}
	b.entities[e] = BoardEntity{subscribedSystems: systemNames}

	b.OnEntityCreated(e)
}

func (b *Board) AddSystem(s ISystem) {
	if b.systems == nil {
		b.systems = make(map[string]ISystem)
	}

	b.systems[s.GetName()] = s
}

func (b *Board) SetRenderer(r IRenderableSystem) {
	b.renderer = r
}

func (b *Board) Update() error {
	for e, boardEntity := range b.entities {
		for _, systemName := range boardEntity.subscribedSystems {
			system, ok := b.systems[systemName]
			if ok {
				system.Update(e)
			}
		}
	}

	return nil
}

func (b *Board) Draw(screen *ebiten.Image) {
	b.renderer.Render(screen)
}
