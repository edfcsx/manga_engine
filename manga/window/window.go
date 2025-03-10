package mangaWindow

import (
	"github.com/veandco/go-sdl2/sdl"
	"manga_engine/vector"
)

const WindowCentered = sdl.WINDOWPOS_CENTERED

type Window struct {
	title string
	pos   vector.Vec2[int32]
	size  vector.Vec2[int32]
}

func Make() *Window {
	return &Window{
		title: "",
		pos:   vector.MakeVec2[int32](0, 0),
		size:  vector.MakeVec2[int32](0, 0),
	}
}

func (w *Window) Title() string {
	return w.title
}

func (w *Window) Position() vector.Vec2[int32] {
	return w.pos
}

func (w *Window) Size() vector.Vec2[int32] {
	return w.size
}

func (w *Window) SetTitle(title string) {
	w.title = title
}

func (w *Window) SetPosition(x, y int32) {
	w.pos.X = x
	w.pos.Y = y
}

func (w *Window) SetSize(weight, height int32) {
	w.size.X = weight
	w.size.Y = height
}
