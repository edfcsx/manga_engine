package scene

import (
	mangaI "github.com/edfcsx/manga_engine/interfaces"
)

type Scene struct {
	entities          []mangaI.Entity
	garbage           map[int]interface{}
	InitializeHandler func()
	UpdateHandler     func(deltaTime float64)
	RenderHandler     func()
}

func (s *Scene) Initialize() {
	//s.entities = make([]mangaI.Entity, 0, 10)
	s.garbage = make(map[int]interface{})

	if s.InitializeHandler != nil {
		s.InitializeHandler()
	}
}

func (s *Scene) Update(deltaTime float64) {
	for _, entity := range s.entities {
		if entity.IsActive() {
			entity.Update(deltaTime)
		}
	}

	if s.UpdateHandler != nil {
		s.UpdateHandler(deltaTime)
	}

	// remove garbage entities
	s.removeGarbageEntities()
}

func (s *Scene) Render() {
	for _, entity := range s.entities {
		if entity.IsActive() {
			entity.Render()
		}
	}

	if s.RenderHandler != nil {
		s.RenderHandler()
	}
}

func (s *Scene) AddEntity(entity mangaI.Entity) {
	index := len(s.entities)
	s.entities = append(s.entities, entity)
	entity.Initialize()

	entity.SetDestroy(func() {
		s.garbage[index] = nil
	})
}

func (s *Scene) removeGarbageEntities() {
	buffer := make([]mangaI.Entity, 0, len(s.entities))

	for i, entity := range s.entities {
		if _, found := s.garbage[i]; !found {
			buffer = append(buffer, entity)
		}
	}

	s.entities = buffer
	s.garbage = make(map[int]interface{})
}
