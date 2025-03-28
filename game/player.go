package game

import (
	"github.com/edfcsx/manga_engine/entity"
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/edfcsx/manga_engine/keyboard"
	"github.com/edfcsx/manga_engine/texture"
)

type Player struct {
	entity.Entity
}

func MakePlayer() *Player {
	player := &Player{}
	return player
}

func (p *Player) Initialize() {
	p.Label = "player"
	p.SetIsActive(true)

	transform := entity.MakeTransformComponent(p)
	transform.Size(32, 32)
	transform.Scale(4)
	transform.Velocity(200, 200)

	upkeys := []int{keyboard.GetKeyCode("w"), keyboard.GetKeyCode("up")}
	downkeys := []int{keyboard.GetKeyCode("s"), keyboard.GetKeyCode("down")}
	leftkeys := []int{keyboard.GetKeyCode("a"), keyboard.GetKeyCode("left")}
	rightkeys := []int{keyboard.GetKeyCode("d"), keyboard.GetKeyCode("right")}

	transform.Enable8DirectionsMove(upkeys, downkeys, leftkeys, rightkeys)

	p.AddComponent(mangaI.TransformComponentID, transform)

	err := texture.MakeTexture("start", "./assets/start.png")

	if err != nil {
		panic(err)
	}

	p.AddComponent(mangaI.SpriteComponentID, entity.MakeSpriteComponent(p, "start"))

	p.AddComponent(mangaI.ScriptComponentID, entity.MakeScriptComponent(p, func() {
	}))
}
