package main

import (
	"manga_engine/manga"
	mangaWindow "manga_engine/manga/window"
)

type Home struct {
}

func (h *Home) Initialize() {

}

func main() {
	window := mangaWindow.Make()
	window.SetTitle("Hello world!")
	window.SetSize(800, 600)
	window.SetPosition(mangaWindow.WindowCentered, mangaWindow.WindowCentered)

	manga.Engine.Initialize(window, &Home{})
}
