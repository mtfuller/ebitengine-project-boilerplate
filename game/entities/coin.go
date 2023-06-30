package entities

import (
	"github.com/yourname/yourgame/framework"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

func NewCoin(boardEntity framework.BoardMapEntity) ecs.Entity {
	myEntity := ecs.NewEntity("COIN")

	myEntity.SetComponent(&components.Position{
		X: float64(boardEntity.X),
		Y: float64(boardEntity.Y),
	})

	myEntity.SetComponent(&components.Velocity{
		VX: 0,
		VY: 0,
	})

	myEntity.SetComponent(&components.Size{
		OffsetX: 0,
		OffsetY: 0,
		W:       boardEntity.W,
		H:       boardEntity.H,
	})

	myEntity.SetComponent(&components.Collision{
		Enabled: true,
		Solid:   false,
	})

	myEntity.SetComponent(&components.Gravity{
		Enabled: false,
	})

	myEntity.SetComponent(&components.Render{
		Spritesheet:  *boardEntity.Spritesheet,
		EntityName:   "COIN",
		SpriteName:   boardEntity.SpriteName,
		CurrentFrame: 0,
		FrameCount:   0,
		Z:            1,
	})

	return myEntity
}
