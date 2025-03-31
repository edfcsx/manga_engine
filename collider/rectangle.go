package collider

import (
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/veandco/go-sdl2/sdl"
)

type Rectangle struct {
	x      int32
	y      int32
	width  int32
	height int32
}

func MakeRectangleShape(width int32, height int32) *Rectangle {
	return &Rectangle{
		x:      0,
		y:      0,
		width:  width,
		height: height,
	}
}

func (r *Rectangle) GetType() int32 {
	return mangaI.RectangleType
}

func (r *Rectangle) X() int32 {
	return r.x
}

func (r *Rectangle) Y() int32 {
	return r.y
}

func (r *Rectangle) MoveTo(x, y int32) {
	r.x = x
	r.y = y
}

func (r *Rectangle) Render(t mangaI.TransformComponent, renderer *sdl.Renderer) {
	err := renderer.SetDrawColor(255, 0, 0, 255)

	if err != nil {
		// TODO: add errors in log
		return
	}

	err = renderer.DrawRect(&sdl.Rect{
		X: r.x,
		Y: r.y,
		W: r.width,
		H: r.height,
	})

	if err != nil {
		// TODO: add error in log
		return
	}
}

func (r *Rectangle) Left() int32 {
	return r.x
}

func (r *Rectangle) Right() int32 {
	return r.x + r.width
}

func (r *Rectangle) Top() int32 {
	return r.y
}

func (r *Rectangle) Bottom() int32 {
	return r.y + r.height
}

func (r *Rectangle) CollidesWith(shape mangaI.ColliderShape) bool {
	switch shape.GetType() {
	case mangaI.PointType:
		return r.collidesPoint(shape)
	case mangaI.CircleType:
		return r.collidesCircle(shape)
	case mangaI.RectangleType:
		return r.collidesRect(shape)
	case mangaI.LineType:
		return r.collidesLine(shape)
	default:
		// TODO: add in log unknown shape
		return false
	}
}

func (r *Rectangle) collidesPoint(b interface{}) bool {
	point, ok := b.(mangaI.PointShape)

	if !ok {
		// Todo: remove panic and log error
		panic(ok)
	}

	return point.CollidesWith(r)
}

func (r *Rectangle) collidesCircle(b interface{}) bool {
	circle, ok := b.(mangaI.CircleShape)

	if !ok {
		// Todo: remove panic and log error
		panic(ok)
	}

	return circle.CollidesWith(r)
}

func (r *Rectangle) collidesRect(c interface{}) bool {
	b, ok := c.(mangaI.RectangleShape)

	if !ok {
		// Todo: remove panic and log error
		panic(ok)
	}

	// verificando se existe sobreposição no eixo X
	if b.Left() <= r.Right() && r.Left() <= b.Right() {
		// verificando se existe sobreposição no eixo y
		if b.Top() <= r.Bottom() && r.Top() <= b.Bottom() {
			return true
		}
	}

	return false
}

func (r *Rectangle) collidesLine(b interface{}) bool {
	line, ok := b.(mangaI.LineShape)

	if !ok {
		// Todo: remove panic and log error
		panic(ok)
	}

	// Coordenadas dos quatro lados do retângulo
	left := r.Left()
	right := r.Right()
	top := r.Top()
	bottom := r.Bottom()

	// Coordenadas da linha
	ax, ay := line.AX(), line.AY()
	bx, by := line.BX(), line.BY()

	// Verificar colisão com cada lado do retângulo
	if linesIntersect(ax, ay, bx, by, left, top, left, bottom) ||
		linesIntersect(ax, ay, bx, by, left, top, right, top) ||
		linesIntersect(ax, ay, bx, by, right, top, right, bottom) ||
		linesIntersect(ax, ay, bx, by, left, bottom, right, bottom) {
		return true
	}

	return false
}

func linesIntersect(ax, ay, bx, by, cx, cy, dx, dy int32) bool {
	denominator := (bx-ax)*(dy-cy) - (by-ay)*(dx-cx)
	if denominator == 0 {
		return false // Linhas são paralelas
	}

	numerator1 := (ay-cy)*(dx-cx) - (ax-cx)*(dy-cy)
	numerator2 := (ay-cy)*(bx-ax) - (ax-cx)*(by-ay)

	r := float64(numerator1) / float64(denominator)
	s := float64(numerator2) / float64(denominator)

	return (r >= 0 && r <= 1) && (s >= 0 && s <= 1)
}
