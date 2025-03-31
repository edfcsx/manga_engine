package entity

import (
	"github.com/edfcsx/manga_engine/collider"
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/edfcsx/manga_engine/manga"
	"github.com/edfcsx/manga_engine/vector"
)

type ColliderComponent struct {
	Entity      mangaI.Entity
	Shape       mangaI.ColliderShape
	OnCollision func()
	transform   mangaI.TransformComponent
	position    vector.Vec2[float64]
	margins     vector.Vec2[int32]
}

func MakeColliderComponent(entity mangaI.Entity, shape mangaI.ColliderShape, onCollision func()) *ColliderComponent {
	transform := entity.GetComponent(mangaI.TransformComponentID)

	if transform == nil {
		panic("collider component: transform component is required in entity")
	}

	return &ColliderComponent{
		Entity:      entity,
		Shape:       shape,
		OnCollision: onCollision,
		transform:   transform.(mangaI.TransformComponent),
		margins:     vector.MakeVec2[int32](0, 0),
	}
}

func (c *ColliderComponent) GetType() string {
	return mangaI.ColliderComponentID
}

func (c *ColliderComponent) Update(deltaTime float64) {
	c.position = c.transform.GetPosition()
	c.Shape.MoveTo(int32(c.position.X)+c.margins.X, int32(c.position.Y)+c.margins.Y)
}

func (c *ColliderComponent) Render() {
	if collider.ShowCollisionBoxs() {
		c.Shape.Render(c.transform, manga.Engine.GetRenderer())
	}
}

func (c *ColliderComponent) SetMargins(x, y int32) {
	c.margins.X = x
	c.margins.Y = y
}
