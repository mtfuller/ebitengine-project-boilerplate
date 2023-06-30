package ecs

var globalEntityIdCounter uint64 = 0

type Entity struct {
	id         uint64
	name       string
	components map[string]interface{}
}

func NewEntity(name string) Entity {
	e := Entity{
		id:         globalEntityIdCounter,
		name:       name,
		components: make(map[string]interface{}),
	}

	globalEntityIdCounter++

	return e
}

func (e Entity) GetID() uint64 {
	return e.id
}

func (e Entity) GetName() string {
	return e.name
}

func (e Entity) GetComponent(name string) interface{} {
	return e.components[name]
}

func (e Entity) HasAllComponents(names []string) bool {
	hasAllComponents := true

	for _, element := range names {
		_, isFound := e.components[element]

		if !isFound {
			hasAllComponents = false
			break
		}
	}

	return hasAllComponents
}

func (e *Entity) SetComponent(c Component) {
	e.components[c.GetName()] = c
}
