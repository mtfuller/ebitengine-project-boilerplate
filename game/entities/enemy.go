package entities

import (
	"github.com/yourname/yourgame/framework"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

func NewEnemy(boardEntity framework.BoardMapEntity) ecs.Entity {
	myEntity := ecs.NewEntity("ENEMY")

	myEntity.SetComponent(&components.Position{
		X: float64(boardEntity.X),
		Y: float64(boardEntity.Y),
	})

	myEntity.SetComponent(&components.Size{
		OffsetX: 4,
		OffsetY: 6,
		W:       16,
		H:       18,
	})

	myEntity.SetComponent(&components.Velocity{
		VX: 0,
		VY: 0,
	})

	myEntity.SetComponent(&components.Gravity{
		Enabled: true,
	})

	myEntity.SetComponent(&components.Collision{
		Enabled: true,
		Solid:   true,
	})

	myEntity.SetComponent(&components.AI{
		Enabled:        true,
		Direction:      1,
		Speed:          0.5,
		PatrolDistance: 50,
		StartX:         float64(boardEntity.X),
	})

	myEntity.SetComponent(&components.Render{
		Spritesheet:  *boardEntity.Spritesheet,
		EntityName:   boardEntity.Type,
		SpriteName:   boardEntity.SpriteName,
		CurrentFrame: 0,
		FrameCount:   0,
		Z:            2,
	})

	return myEntity
}
