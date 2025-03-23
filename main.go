package main

import (
	"github.com/edfcsx/manga_engine/manga"
	"github.com/edfcsx/manga_engine/window"
)

type Home struct{}

var win = window.Make()

func (h *Home) Initialize() {}
func (h *Home) Update() {
	//win.SetTitle(fmt.Sprintf("Manga Engine - V.0.0.0   FPS: %f", manga.Engine.GetFPS()))
}
func (h *Home) Render() {}

func main() {
	win.SetTitle("Manga Engine - V.0.0.0")
	win.SetPosition(window.PosCentered, window.PosCentered)
	win.SetSize(800, 600)

	manga.Engine.Initialize(win, &Home{}, manga.FPS_UNLIMITED)
}
