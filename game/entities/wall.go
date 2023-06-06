package entities

import (
	"github.com/yourname/yourgame/engine/ecs"
	"github.com/yourname/yourgame/game/components"
	"github.com/yourname/yourgame/game/loaders"
)

func NewWall(x, y int) ecs.Entity {
	myEntity := ecs.NewEntity()

	myEntity.AddComponent(&components.Position{
		X: float64(x),
		Y: float64(y),
	})

	myEntity.AddComponent(&components.Size{
		W: 32,
		H: 32,
	})

	myEntity.AddComponent(&components.Collision{
		Enabled: true,
	})
	
	myEntity.AddComponent(&components.Render{
		Spritesheet: loaders.Wall,
		CurrentSprite: "test",
		CurrentFrame: 0,
		MaxCountdown: 20,
		Countdown: 20,
		Z: 0,
	})

	return myEntity
}