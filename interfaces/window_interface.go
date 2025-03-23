package mangaI

import (
	"github.com/edfcsx/manga_engine/vector"
	"github.com/veandco/go-sdl2/sdl"
)

type Window interface {
	Title() string
	SetTitle(string)
	Size() vector.Vec2[int32]
	SetSize(w, h int32)
	Position() vector.Vec2[int32]
	SetPosition(x, y int32)
	SetGameWindow(window *sdl.Window)
}
