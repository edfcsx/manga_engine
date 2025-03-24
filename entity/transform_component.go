package entity

import (
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/edfcsx/manga_engine/vector"
)

type TransformComponent struct {
	Entity   mangaI.Entity
	position vector.Vec2[int32]
	velocity vector.Vec2[int32]
	size     vector.Vec2[int32]
	scale    int32
}

func MakeTransformComponent(entity mangaI.Entity) *TransformComponent {
	return &TransformComponent{
		Entity:   entity,
		position: vector.MakeVec2[int32](0, 0),
		velocity: vector.MakeVec2[int32](0, 0),
		size:     vector.MakeVec2[int32](0, 0),
		scale:    1,
	}
}

func (t *TransformComponent) GetType() int32 {
	return mangaI.TransformComponentID
}

func (t *TransformComponent) Update(deltaTime float64) {
	t.position.X += int32(float64(t.velocity.X) * deltaTime)
	t.position.Y += int32(float64(t.velocity.Y) * deltaTime)
}

func (t *TransformComponent) Render() {}

func (t *TransformComponent) Position(x, y int32) {
	t.position.X = x
	t.position.Y = y
}

func (t *TransformComponent) Velocity(x, y int32) {
	t.velocity.X = x
	t.velocity.Y = y
}

func (t *TransformComponent) Size(width, height int32) {
	t.size.X = width
	t.size.Y = height
}

func (t *TransformComponent) Scale(s int32) {
	t.scale = s
}

func (t *TransformComponent) GetPosition() vector.Vec2[int32] {
	return t.position
}

func (t *TransformComponent) GetVelocity() vector.Vec2[int32] {
	return t.velocity
}

func (t *TransformComponent) GetSize() vector.Vec2[int32] {
	return t.size
}

func (t *TransformComponent) GetScale() int32 {
	return t.scale
}
