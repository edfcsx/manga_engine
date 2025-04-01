package game

import (
	"github.com/edfcsx/manga_engine/game/enemy"
	"github.com/edfcsx/manga_engine/game/player"
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

	HomeScene.AddEntity(player.MakePlayer())
	HomeScene.AddEntity(enemy.MakeEnemy())
}

func initialize() {
}

func update(deltaTime float64) {

}

func render() {

}
