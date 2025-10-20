package systems

import (
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

type AI struct {
	ecs.BaseSystem
}

func (a AI) GetName() string {
	return "System::AI"
}

func (a AI) Update(e *ecs.Entity) {
	component := e.GetComponent("ai")
	ai, ok := component.(*components.AI)
	if !ok || !ai.Enabled {
		return
	}

	component = e.GetComponent("position")
	position, ok := component.(*components.Position)
	if !ok {
		return
	}

	component = e.GetComponent("velocity")
	velocity, ok := component.(*components.Velocity)
	if !ok {
		return
	}

	// Calculate how far the enemy has moved from its starting position
	distanceFromStart := position.X - ai.StartX

	// Check if enemy has reached the patrol distance limit
	if distanceFromStart >= ai.PatrolDistance {
		ai.Direction = -1
	} else if distanceFromStart <= -ai.PatrolDistance {
		ai.Direction = 1
	}

	// Apply velocity based on AI direction and speed
	velocity.VX = ai.Direction * ai.Speed
}
