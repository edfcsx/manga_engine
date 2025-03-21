package window

import (
	"github.com/edfcsx/manga_engine/vector"
	"github.com/veandco/go-sdl2/sdl"
)

const PosCentered = sdl.WINDOWPOS_CENTERED

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

func (w *Window) SetTitle(title string) {
	w.title = title
}

func (w *Window) Size() vector.Vec2[int32] {
	return w.size
}

func (w *Window) SetSize(width, height int32) {
	w.size.X = width
	w.size.Y = height
}

func (w *Window) Position() vector.Vec2[int32] {
	return w.pos
}

func (w *Window) SetPosition(x, y int32) {
	w.pos.X = x
	w.pos.Y = y
}
