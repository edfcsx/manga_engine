package game

import (
	"github.com/edfcsx/manga_engine/scene"
)

type Game struct {
	scene.Scene
}

var HomeScene = &Game{}

func init() {
	HomeScene.InitializeHandler = initialize
	HomeScene.UpdateHandler = update
	HomeScene.RenderHandler = render

	HomeScene.AddEntity(MakePlayer())
}

func initialize() {
}

func update(deltaTime float64) {

}

func render() {

}
