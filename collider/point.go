package collider

import (
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Point struct {
	x int32
	y int32
}

func MakePointShape() *Point {
	return &Point{
		x: 0,
		y: 0,
	}
}

func (p *Point) GetType() int32 {
	return mangaI.PointType
}

func (p *Point) X() int32 {
	return p.x
}

func (p *Point) Y() int32 {
	return p.y
}

func (p *Point) MoveTo(x, y int32) {
	p.x = x
	p.y = y
}

func (p *Point) Render(t mangaI.TransformComponent, r *sdl.Renderer) {
	// TODO: add errors in log
	err := r.SetDrawColor(255, 0, 0, 255)
	if err != nil {
		return
	}

	err = r.DrawPoint(p.x, p.y)
	if err != nil {
		return
	}
}

func (p *Point) Distance(x int32, y int32) float64 {
	deltaX := math.Abs(float64(x - p.x))
	deltaY := math.Abs(float64(y - p.y))

	return math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))
}

func (p *Point) CollidesWith(shape mangaI.ColliderShape) bool {
	switch shape.GetType() {
	case mangaI.PointType:
		return p.collidesPoint(shape)
	case mangaI.CircleType:
		return p.collidesCircle(shape)
	case mangaI.RectangleType:
		return p.collidesRectangle(shape)
	case mangaI.LineType:
		return p.collidesLine(shape)
	default:
		return false
	}
}

func (p *Point) collidesPoint(b mangaI.ColliderShape) bool {
	return p.X() == b.X() && p.Y() == b.X()
}

func (p *Point) collidesCircle(b interface{}) bool {
	circle, ok := b.(mangaI.CircleShape)

	if !ok {
		//TODO: remove panic and log error
		panic(ok)
	}

	return p.Distance(circle.X(), circle.Y()) <= circle.GetRadius()
}

func (p *Point) collidesRectangle(b interface{}) bool {
	rect, ok := b.(mangaI.RectangleShape)

	if !ok {
		panic(ok)
	}

	if p.X() >= rect.Left() && p.X() <= rect.Right() {
		if p.Y() >= rect.Top() && p.Y() <= rect.Bottom() {
			return true
		}
	}

	return false
}

func (p *Point) collidesLine(b interface{}) bool {
	l, ok := b.(mangaI.LineShape)

	if !ok {
		panic(ok)
	}

	// Coordenadas do ponto
	px, py := float64(p.X()), float64(p.Y())

	// Coordenadas dos pontos da linha
	ax, ay := float64(l.AX()), float64(l.AY())
	bx, by := float64(l.BX()), float64(l.BY())

	// Calcular a distância perpendicular do ponto à linha
	numerator := math.Abs((by-ay)*px - (bx-ax)*py + bx*ay - by*ax)
	denominator := math.Sqrt((by-ay)*(by-ay) + (bx-ax)*(bx-ax))
	distance := numerator / denominator

	// Verificar se o ponto está dentro dos limites da linha
	if distance == 0 {
		if (px >= math.Min(ax, bx) && px <= math.Max(ax, bx)) && (py >= math.Min(ay, by) && py <= math.Max(ay, by)) {
			return true
		}
	}

	return false
}
