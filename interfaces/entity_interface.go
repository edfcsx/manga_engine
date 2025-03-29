package mangaI

import "github.com/edfcsx/manga_engine/vector"

const (
	TransformComponentID = "TRANSFORM"
	SpriteComponentID    = "SPRITE"
	ScriptComponentID
)

type Component interface {
	GetType() string
	Update(deltaTime float64)
	Render()
}

type Entity interface {
	AddComponent(componentType string, c Component)
	GetComponent(componentType string) Component
	GetLabel() string
	Initialize()
	Update(deltaTime float64)
	Render()
	SetSelf(self interface{})
	IsActive() bool
	SetIsActive(status bool)
	Destroy()
	SetDestroy(destroy func())
}

type TransformComponent interface {
	Component
	Position(x, y int32)
	Velocity(x, y int32)
	Size(x, y int32)
	Scale(int32)
	GetPosition() vector.Vec2[int32]
	GetVelocity() vector.Vec2[int32]
	GetSize() vector.Vec2[int32]
	GetScale() int32
}
