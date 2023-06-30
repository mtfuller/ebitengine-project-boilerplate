package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BoardEntity struct {
	ecsEntity         *Entity
	subscribedSystems []string
}

type Board struct {
	entities          map[uint64]BoardEntity
	systems           map[string]ISystem
	renderer          IRenderableSystem
	OnEntityCreated   func(*Entity)
	OnEntityDestroyed func(*Entity)
}

func (b *Board) AddEntity(e *Entity) {
	b.entities[e.id] = BoardEntity{
		ecsEntity:         e,
		subscribedSystems: []string{},
	}
}

func (b Board) GetEntity(id uint64) *Entity {
	return b.entities[id].ecsEntity
}

func (b *Board) RemoveEntity(e *Entity) {
	b.OnEntityDestroyed(e)

	delete(b.entities, e.id)
}

func (b *Board) AddEntityToSystems(e *Entity, systemNames ...string) {
	if b.entities == nil {
		b.entities = make(map[uint64]BoardEntity)
	}

	b.entities[e.id] = BoardEntity{
		ecsEntity:         e,
		subscribedSystems: systemNames,
	}

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
	for _, boardEntity := range b.entities {
		for _, systemName := range boardEntity.subscribedSystems {
			system, ok := b.systems[systemName]
			if ok {
				system.Update(boardEntity.ecsEntity)
			}
		}
	}

	return nil
}

func (b *Board) Draw(screen *ebiten.Image) {
	b.renderer.Render(screen)
}
