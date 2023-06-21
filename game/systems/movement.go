package systems

import (
	//"fmt"
	"github.com/solarlune/resolv"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

type Movement struct {
	ecs.BaseSystem
	collisionSpace   *resolv.Space
	collisionObjects map[*ecs.Entity]*resolv.Object
}

func (m *Movement) AddEntity(e *ecs.Entity) {
	if m.Entities == nil {
		m.Entities = make(map[*ecs.Entity]struct{})
	}

	if m.collisionSpace == nil {
		m.collisionSpace = resolv.NewSpace(640, 480, 2, 2)
	}

	if m.collisionObjects == nil {
		m.collisionObjects = make(map[*ecs.Entity]*resolv.Object)
	}

	m.Entities[e] = struct{}{}

	component := e.GetComponent("position")
	position, _ := component.(*components.Position)

	component = e.GetComponent("size")
	size, _ := component.(*components.Size)

	component = e.GetComponent("collision")
	collision, ok := component.(*components.Collision)

	if ok && collision.Enabled {
		x := position.X + float64(size.OffsetX)
		y := position.Y + float64(size.OffsetY)
		w := size.W
		h := size.H

		obj := resolv.NewObject(float64(x), float64(y), float64(w), float64(h))

		m.collisionObjects[e] = obj

		m.collisionSpace.Add(obj)
	}
}

func (m Movement) Update() {
	for ptr, _ := range m.Entities {
		e := *ptr

		component := e.GetComponent("position")
		position, ok := component.(*components.Position)
		if !ok {
			continue
		}

		component = e.GetComponent("velocity")
		velocity, ok := component.(*components.Velocity)
		if !ok {
			continue
		}

		component = e.GetComponent("gravity")
		gravity, ok := component.(*components.Gravity)
		if ok && gravity.Enabled {
			velocity.VY += 0.1
		}

		if collision := m.collisionObjects[ptr].Check(velocity.VX, 0); collision != nil {
			velocity.VX = collision.ContactWithObject(collision.Objects[0]).X()
		}

		if collision := m.collisionObjects[ptr].Check(0, velocity.VY); collision != nil {
			velocity.VY = collision.ContactWithObject(collision.Objects[0]).Y()
		}

		m.collisionObjects[ptr].X += velocity.VX
		m.collisionObjects[ptr].Y += velocity.VY
		m.collisionObjects[ptr].Update()

		position.X += velocity.VX
		position.Y += velocity.VY
	}
}
