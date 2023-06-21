package entities

import (
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

func NewWall(x, y int) ecs.Entity {
	myEntity := ecs.NewEntity()

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

	myEntity.SetComponent(&components.Collision{
		Enabled: true,
	})

	// myEntity.SetComponent(&components.Render{
	// 	Spritesheet:   loaders.Wall,
	// 	CurrentSprite: "test",
	// 	CurrentFrame:  0,
	// 	MaxCountdown:  20,
	// 	Countdown:     20,
	// 	Z:             0,
	// })

	return myEntity
}
