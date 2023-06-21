package ecs

//"fmt"

type Entity struct {
	components map[string]interface{}
}

func NewEntity() Entity {
	return Entity{components: make(map[string]interface{})}
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
