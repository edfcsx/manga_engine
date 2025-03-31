package collider

import (
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/edfcsx/manga_engine/vector"
	"github.com/veandco/go-sdl2/sdl"
)

type Line struct {
	position vector.Vec2[int32]
	a        vector.Vec2[int32]
	b        vector.Vec2[int32]
}

func MakeLineShape(x1 int32, y1 int32, x2 int32, y2 int32) *Line {
	return &Line{
		position: vector.MakeVec2[int32](0, 0),
		a:        vector.MakeVec2[int32](x1, y1),
		b:        vector.MakeVec2[int32](x2, y2),
	}
}

func (l *Line) X() int32 {
	return l.position.X
}

func (l *Line) Y() int32 {
	return l.position.Y
}

func (l *Line) GetType() int32 {
	return mangaI.LineType
}

func (l *Line) AX() int32 {
	return l.position.X + l.a.X
}

func (l *Line) AY() int32 {
	return l.position.Y + l.a.Y
}

func (l *Line) BX() int32 {
	return l.position.X + l.b.X
}

func (l *Line) BY() int32 {
	return l.position.Y + l.b.Y
}

func (l *Line) MoveTo(x, y int32) {
	l.position.X = x
	l.position.Y = y
}

func (l *Line) Render(t mangaI.TransformComponent, renderer *sdl.Renderer) {
	err := renderer.SetDrawColor(255, 0, 0, 255)

	if err != nil {
		// TODO: add errors in log
		return
	}

	err = renderer.DrawLine(l.AX(), l.AY(), l.BX(), l.BY())

	if err != nil {
		// TODO: add errors in log
		return
	}
}

func (l *Line) CollidesWith(shape mangaI.Shape) bool {
	switch shape.GetType() {
	case mangaI.PointType:
		return l.collidesPoint(shape)
	case mangaI.CircleType:
		return l.collidesCircle(shape)
	case mangaI.RectangleType:
		return l.collidesRectangle(shape)
	case mangaI.LineType:
		return l.collidesLine(shape)
	}

	return false
}

func (l *Line) collidesPoint(b interface{}) bool {
	point, ok := b.(mangaI.PointShape)

	if !ok {
		// TODO: add in log unknown shape
		panic(ok)
	}

	return point.CollidesWith(l)
}

func (l *Line) collidesCircle(b interface{}) bool {
	circle, ok := b.(mangaI.CircleShape)

	if !ok {
		// Todo: add error in log
		panic(ok)
	}

	return circle.CollidesWith(l)
}

func (l *Line) collidesRectangle(b interface{}) bool {
	rectangle, ok := b.(mangaI.RectangleShape)

	if !ok {
		// Todo: add error in log
		panic(ok)
	}

	return rectangle.CollidesWith(l)
}

func (l *Line) collidesLine(c interface{}) bool {
	b, ok := c.(mangaI.LineShape)

	if !ok {
		//Todo: add error in log
		panic(ok)
	}

	ax1, ay1 := l.AX(), l.AY()
	ax2, ay2 := l.BX(), l.BY()
	bx1, by1 := b.AX(), b.AY()
	bx2, by2 := b.BX(), b.BY()

	// Helper function to calculate the orientation of the triplet (p, q, r)
	// 0 -> p, q and r are collinear
	// 1 -> Clockwise
	// 2 -> Counterclockwise
	orientation := func(px, py, qx, qy, rx, ry int32) int {
		val := (qy-py)*(rx-qx) - (qx-px)*(ry-qy)
		if val == 0 {
			return 0
		}
		if val > 0 {
			return 1
		}
		return 2
	}

	// Check if point q lies on segment pr
	onSegment := func(px, py, qx, qy, rx, ry int32) bool {
		if qx <= max(px, rx) && qx >= min(px, rx) && qy <= max(py, ry) && qy >= min(py, ry) {
			return true
		}
		return false
	}

	// Find the four orientations needed for the general and special cases
	o1 := orientation(ax1, ay1, ax2, ay2, bx1, by1)
	o2 := orientation(ax1, ay1, ax2, ay2, bx2, by2)
	o3 := orientation(bx1, by1, bx2, by2, ax1, ay1)
	o4 := orientation(bx1, by1, bx2, by2, ax2, ay2)

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Special cases
	// a1, a2 and b1 are collinear and b1 lies on segment a1a2
	if o1 == 0 && onSegment(ax1, ay1, bx1, by1, ax2, ay2) {
		return true
	}

	// a1, a2 and b2 are collinear and b2 lies on segment a1a2
	if o2 == 0 && onSegment(ax1, ay1, bx2, by2, ax2, ay2) {
		return true
	}

	// b1, b2 and a1 are collinear and a1 lies on segment b1b2
	if o3 == 0 && onSegment(bx1, by1, ax1, ay1, bx2, by2) {
		return true
	}

	// b1, b2 and a2 are collinear and a2 lies on segment b1b2
	if o4 == 0 && onSegment(bx1, by1, ax2, ay2, bx2, by2) {
		return true
	}

	// Doesn't fall in any of the above cases
	return false
}
