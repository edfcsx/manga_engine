package main

import (
	"github.com/edfcsx/manga_engine/manga"
	"github.com/edfcsx/manga_engine/window"
)

type Home struct{}

func (h *Home) Initialize() {}
func (h *Home) Update()     {}
func (h *Home) Render()     {}

func main() {
	win := window.Make()
	win.SetTitle("Manga Engine - V.0.0.0")
	win.SetPosition(window.PosCentered, window.PosCentered)
	win.SetSize(800, 600)

	manga.Engine.Initialize(win, &Home{})
}
