package main

import (
	"github.com/edfcsx/manga_engine/game"
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
	win.SetSize(800, 600)

	manga.Engine.Initialize(win, game.HomeScene, manga.FPS_UNLIMITED)
}

// TODO: é importante testar o destroy das entidades
