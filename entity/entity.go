package entity

import mangaI "github.com/edfcsx/manga_engine/interfaces"

type Entity struct {
	Label          string
	IsActive       bool
	components     map[string]mangaI.Component
	Self           interface{}
	destroyHandler []func()
}

func (e *Entity) AddComponent(componentType string, c mangaI.Component) {
	e.components[componentType] = c
}

func (e *Entity) GetComponent(componentType string) mangaI.Component {
	return e.components[componentType]
}

func (e *Entity) Update(deltaTime float64) {
	for _, c := range e.components {
		c.Update(deltaTime)
	}
}

func (e *Entity) Render() {
	for _, c := range e.components {
		c.Render()
	}
}

func (e *Entity) SetSelf(self interface{}) {
	e.Self = self
}

func (e *Entity) Destroy() {
	for _, fn := range e.destroyHandler {
		fn()
	}
}

func (e *Entity) SetDestroy(fn func()) {
	e.destroyHandler = append(e.destroyHandler, fn)
}
