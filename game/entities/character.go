package entities

import (
	"github.com/yourname/yourgame/framework"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

func NewCharacter(boardEntity framework.BoardMapEntity) ecs.Entity {
	myEntity := ecs.NewEntity("CHAR")

	myEntity.SetComponent(&components.Input{
		Enabled: true,
	})

	myEntity.SetComponent(&components.Position{
		X: float64(boardEntity.X),
		Y: float64(10),
	})

	myEntity.SetComponent(&components.Size{
		OffsetX: 16,
		OffsetY: 6,
		W:       6,
		H:       18,
	})

	myEntity.SetComponent(&components.Gravity{
		Enabled: true,
	})

	myEntity.SetComponent(&components.Collision{
		Enabled: true,
		Solid:   true,
	})

	myEntity.SetComponent(&components.Velocity{
		VX: 0,
		VY: 0,
	})

	myEntity.SetComponent(&components.Render{
		Spritesheet:  *boardEntity.Spritesheet,
		EntityName:   "CHAR1",
		SpriteName:   "0",
		CurrentFrame: 0,
		FrameCount:   0,
		Z:            2,
	})

	return myEntity
}
