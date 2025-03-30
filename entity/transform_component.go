package entity

import (
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/edfcsx/manga_engine/keyboard"
	"github.com/edfcsx/manga_engine/vector"
)

type moveKeys struct {
	activate   bool
	directions vector.Vec2[int32]
	up         []int
	down       []int
	left       []int
	right      []int
}

type TransformComponent struct {
	Entity   mangaI.Entity
	position vector.Vec2[float64]
	velocity vector.Vec2[int32]
	size     vector.Vec2[int32]
	scale    int32
	move     moveKeys
}

func MakeTransformComponent(entity mangaI.Entity) *TransformComponent {
	return &TransformComponent{
		Entity:   entity,
		position: vector.MakeVec2[float64](0.0, 0.0),
		velocity: vector.MakeVec2[int32](0, 0),
		size:     vector.MakeVec2[int32](0, 0),
		scale:    1,
		move: moveKeys{
			activate:   false,
			directions: vector.MakeVec2[int32](0, 0),
			up:         nil,
			down:       nil,
			left:       nil,
			right:      nil,
		},
	}
}

func (t *TransformComponent) GetType() string {
	return mangaI.TransformComponentID
}

func (t *TransformComponent) Update(deltaTime float64) {
	if !t.move.activate {
		t.position.X += float64(t.velocity.X) * deltaTime
		t.position.Y += float64(t.velocity.Y) * deltaTime

		return
	}

	t.move.directions.X = 0
	t.move.directions.Y = 0

	if keyboard.IsAnyKeyPressed(t.move.up) && keyboard.IsAnyKeyPressed(t.move.down) {
		t.move.directions.Y = 0
	} else {
		if keyboard.IsAnyKeyPressed(t.move.up) {
			t.move.directions.Y = -1
		} else if keyboard.IsAnyKeyPressed(t.move.down) {
			t.move.directions.Y = 1
		}
	}

	if keyboard.IsAnyKeyPressed(t.move.left) && keyboard.IsAnyKeyPressed(t.move.right) {
		t.move.directions.X = 0
	} else {
		if keyboard.IsAnyKeyPressed(t.move.left) {
			t.move.directions.X = -1
		} else if keyboard.IsAnyKeyPressed(t.move.right) {
			t.move.directions.X = 1
		}
	}

	t.position.X += float64(t.velocity.X*t.move.directions.X) * deltaTime
	t.position.Y += float64(t.velocity.Y*t.move.directions.Y) * deltaTime
}

func (t *TransformComponent) Render() {}

func (t *TransformComponent) Position(x, y float64) {
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

func (t *TransformComponent) GetPosition() vector.Vec2[float64] {
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

func (t *TransformComponent) Enable8DirectionsMove(upKeys []int, downKeys []int, leftKeys []int, rightKeys []int) {
	t.move.activate = true
	t.move.up = upKeys
	t.move.down = downKeys
	t.move.right = rightKeys
	t.move.left = leftKeys
}
