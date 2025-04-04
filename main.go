package main

import (
	"github.com/edfcsx/manga_engine/game"
	"github.com/edfcsx/manga_engine/keyboard"
	"github.com/edfcsx/manga_engine/manga"
	"github.com/edfcsx/manga_engine/scene"
	"github.com/edfcsx/manga_engine/window"
)

type Home struct {
	scene.Scene
}

var win = window.Make()

func main() {
	win.SetTitle("Manga Engine - V.0.0.0")
	win.SetPosition(window.PosCentered, window.PosCentered)
	win.SetSize(1920, 1080)
	win.SetIsResizable(false)

	manga.Engine.AddGlobalScript(func() {
		if keyboard.IsKeyPressed(keyboard.GetKeyCode("escape")) {
			manga.Engine.Stop()
		}
	})

	manga.Engine.Initialize(win, game.HomeScene, 60)
}

// TODO: é importante testar o destroy das entidades
