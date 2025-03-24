package entity

import mangaI "github.com/edfcsx/manga_engine/interfaces"

type Entity struct {
	Label      string
	IsActive   bool
	components map[string]mangaI.EntityComponent
	Self       interface{}
}

func (e *Entity) AddComponent(componentType string, c mangaI.EntityComponent) {
	e.components[componentType] = c
}

func (e *Entity) GetComponent(componentType string) mangaI.EntityComponent {
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
