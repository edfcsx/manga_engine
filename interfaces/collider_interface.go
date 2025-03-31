package mangaI

import "github.com/veandco/go-sdl2/sdl"

type ColliderShape interface {
	X() int32
	Y() int32
	MoveTo(x, y int32)
	GetType() int32
	CollidesWith(shape ColliderShape) bool
	Render(t TransformComponent, r *sdl.Renderer)
	Scale(x int32)
}

const (
	PointType = iota
	LineType
	CircleType
	RectangleType
)

type PointShape interface {
	ColliderShape
	Distance(x, y int32) float64
}

type CircleShape interface {
	ColliderShape
	GetRadius() float64
}

type LineShape interface {
	ColliderShape
	AX() int32
	AY() int32
	BX() int32
	BY() int32
}

type RectangleShape interface {
	ColliderShape
	Left() int32
	Right() int32
	Top() int32
	Bottom() int32
}

type ColliderType int32

const (
	ColliderStatic = iota
	ColliderMoving
)
