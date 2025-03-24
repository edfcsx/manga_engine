package main

import (
	"github.com/edfcsx/manga_engine/manga"
	"github.com/edfcsx/manga_engine/scene"
	"github.com/edfcsx/manga_engine/texture"
	"github.com/edfcsx/manga_engine/vector"
	"github.com/edfcsx/manga_engine/window"
)

type Home struct {
	scene.Scene
}

var win = window.Make()

func main() {
	win.SetTitle("Manga Engine - V.0.0.0")
	win.SetPosition(window.PosCentered, window.PosCentered)
	win.SetSize(800, 600)

	home := &Home{}

	home.InitializeHandler = func() {
		err := texture.MakeTexture("start", "./assets/start.png")

		if err != nil {
			panic(err)
		}
	}

	home.UpdateHandler = func() {
		//win.SetTitle(fmt.Sprintf("Manga Engine - V.0.0.0   FPS: %f", manga.Engine.GetFPS()))
	}

	home.RenderHandler = func() {
		src := vector.Vec4[int32]{X: 0, Y: 0, W: 32, H: 32}
		dest := vector.Vec4[int32]{X: 32, Y: 32, W: 32 * 4, H: 32 * 4}

		err := manga.Engine.Draw("start", src, dest, 0.0)

		if err != nil {
			panic(err)
		}
	}

	manga.Engine.Initialize(win, &Home{}, manga.FPS_UNLIMITED)
}
