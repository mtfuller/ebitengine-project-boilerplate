package systems

import (
	"strconv"

	"github.com/solarlune/resolv"
	"github.com/yourname/yourgame/framework/ecs"
	"github.com/yourname/yourgame/game/components"
)

type Movement struct {
	ecs.BaseSystem
	collisionEvents  map[string]map[string]func(uint64, uint64)
	collisionSpace   *resolv.Space
	collisionObjects map[*ecs.Entity]*resolv.Object
}

func (m Movement) GetName() string {
	return "System::Movement"
}

func (m *Movement) WhenEntityTouchesAnother(first string, second string, handler func(uint64, uint64)) {
	if m.collisionEvents == nil {
		m.collisionEvents = make(map[string]map[string]func(uint64, uint64))
	}

	_, ok := m.collisionEvents[first]
	if !ok {
		m.collisionEvents[first] = make(map[string]func(uint64, uint64))
	}

	m.collisionEvents[first][second] = handler
}

func (m *Movement) HandleEntityCreated(e *ecs.Entity) {

	if m.collisionSpace == nil {
		m.collisionSpace = resolv.NewSpace(640, 480, 2, 2)
	}

	if m.collisionObjects == nil {
		m.collisionObjects = make(map[*ecs.Entity]*resolv.Object)
	}

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

		objectType := "NONE"
		if collision.Solid {
			objectType = "SOLID"
		}

		obj := resolv.NewObject(float64(x), float64(y), float64(w), float64(h), e.GetName(), strconv.FormatUint(e.GetID(), 10), objectType)

		m.collisionObjects[e] = obj

		m.collisionSpace.Add(obj)
	}
}

func (m *Movement) HandleEntityDestoryed(e *ecs.Entity) {
	collisionObject, hasObject := m.collisionObjects[e]
	if hasObject {
		m.collisionSpace.Remove(collisionObject)
	}

	delete(m.collisionObjects, e)
}

func (m Movement) Update(e *ecs.Entity) {
	component := e.GetComponent("position")
	position, ok := component.(*components.Position)
	if !ok {
		return
	}

	component = e.GetComponent("velocity")
	velocity, ok := component.(*components.Velocity)
	if !ok {
		return
	}

	component = e.GetComponent("gravity")
	gravity, ok := component.(*components.Gravity)
	if ok && gravity.Enabled {
		velocity.VY += 0.1
	}

	// component = e.GetComponent("collision")
	// collisionComponent, ok := component.(*components.Collision)

	if collision := m.collisionObjects[e].Check(velocity.VX, 0); collision != nil {
		_, hasEvents := m.collisionEvents[e.GetName()]

		for _, object := range collision.Objects {
			objType := object.Tags()[2]

			if objType == "SOLID" {
				velocity.VX = collision.ContactWithObject(collision.Objects[0]).X()
			}

			if hasEvents {
				for _, object := range collision.Objects {
					entityName := object.Tags()[0]

					entityID, err := strconv.ParseUint(object.Tags()[1], 10, 64)
					if err != nil {
						panic("Unable to parse entity ID")
					}

					_, hasHandler := m.collisionEvents[e.GetName()][entityName]
					if hasHandler {
						m.collisionEvents[e.GetName()][entityName](e.GetID(), entityID)
					}
				}
			}
		}
	}

	if collision := m.collisionObjects[e].Check(0, velocity.VY); collision != nil {
		_, hasEvents := m.collisionEvents[e.GetName()]

		for _, object := range collision.Objects {
			objType := object.Tags()[2]

			if objType == "SOLID" {
				velocity.VY = collision.ContactWithObject(collision.Objects[0]).Y()
			}

			if hasEvents {
				for _, object := range collision.Objects {
					entityName := object.Tags()[0]

					entityID, err := strconv.ParseUint(object.Tags()[1], 10, 64)
					if err != nil {
						panic("Unable to parse entity ID")
					}

					_, hasHandler := m.collisionEvents[e.GetName()][entityName]
					if hasHandler {
						m.collisionEvents[e.GetName()][entityName](e.GetID(), entityID)
					}
				}
			}
		}
	}

	m.collisionObjects[e].X += velocity.VX
	m.collisionObjects[e].Y += velocity.VY
	m.collisionObjects[e].Update()

	position.X += velocity.VX
	position.Y += velocity.VY
}
