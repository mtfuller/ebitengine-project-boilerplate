package entities

import (
	"github.com/yourname/yourgame/engine/ecs"
	"github.com/yourname/yourgame/game/components"
	"github.com/yourname/yourgame/game/loaders"
)

func NewCharacter() ecs.Entity {
	myEntity := ecs.NewEntity()

	myEntity.AddComponent(&components.Input{
		Enabled: true,
	})

	myEntity.AddComponent(&components.Position{
		X: 100,
		Y: 100,
	})

	myEntity.AddComponent(&components.Size{
		W: 32,
		H: 32,
	})

	myEntity.AddComponent(&components.Gravity{
		Enabled: true,
	})

	myEntity.AddComponent(&components.Collision{
		Enabled: true,
	})

	myEntity.AddComponent(&components.Velocity{
		VX: 0,
		VY: 0,
	})
	
	myEntity.AddComponent(&components.Render{
		Spritesheet: loaders.Player,
		CurrentSprite: "test",
		CurrentFrame: 1,
		MaxCountdown: 20,
		Countdown: 20,
		Z: 1,
	})

	return myEntity
}