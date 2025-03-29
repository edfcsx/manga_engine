package entity

import mangaI "github.com/edfcsx/manga_engine/interfaces"

type ScriptComponent struct {
	Entity  mangaI.Entity
	Handler func()
}

func MakeScriptComponent(entity mangaI.Entity, fn func()) *ScriptComponent {
	return &ScriptComponent{
		Entity:  entity,
		Handler: fn,
	}
}

func (s *ScriptComponent) GetType() string {
	return mangaI.ScriptComponentID
}

func (s *ScriptComponent) Update(deltaTime float64) {
	if s.Handler != nil {
		s.Handler()
	}
}

func (s *ScriptComponent) Render() {}
