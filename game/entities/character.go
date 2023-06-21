package entities

import (
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

func NewCharacter(x int, y int) ecs.Entity {
	myEntity := ecs.NewEntity()

	myEntity.SetComponent(&components.Input{
		Enabled: true,
	})

	myEntity.SetComponent(&components.Position{
		X: float64(x),
		Y: float64(y),
	})

	myEntity.SetComponent(&components.Size{
		OffsetX: 0,
		OffsetY: 0,
		W:       32,
		H:       32,
	})

	myEntity.SetComponent(&components.Gravity{
		Enabled: true,
	})

	myEntity.SetComponent(&components.Collision{
		Enabled: true,
	})

	myEntity.SetComponent(&components.Velocity{
		VX: 0,
		VY: 0,
	})

	// myEntity.SetComponent(&components.Render{
	// 	Spritesheet: ,
	// 	CurrentSprite: "test",
	// 	CurrentFrame:  1,
	// 	MaxCountdown:  20,
	// 	Countdown:     20,
	// 	Z:             1,
	// })

	return myEntity
}
