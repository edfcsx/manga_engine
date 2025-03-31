package game

import (
	"github.com/edfcsx/manga_engine/collider"
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

var controls = map[string][]int{
	"up":    []int{keyboard.GetKeyCode("w"), keyboard.GetKeyCode("up")},
	"down":  []int{keyboard.GetKeyCode("s"), keyboard.GetKeyCode("down")},
	"left":  []int{keyboard.GetKeyCode("a"), keyboard.GetKeyCode("left")},
	"right": []int{keyboard.GetKeyCode("d"), keyboard.GetKeyCode("right")},
}

var directionsKeys []int
var sprite mangaI.SpriteComponente

func (p *Player) Initialize() {
	p.Label = "player"
	p.SetIsActive(true)

	for _, keys := range controls {
		directionsKeys = append(directionsKeys, keys...)
	}

	transform := entity.MakeTransformComponent(p)
	transform.Size(192, 192)
	transform.Scale(4)
	transform.Velocity(500, 500)

	transform.Enable8DirectionsMove(controls["up"], controls["down"], controls["left"], controls["right"])

	p.AddComponent(mangaI.TransformComponentID, transform)

	//192 X 192

	err := texture.MakeTexture("goblin", "./assets/goblin.png")

	if err != nil {
		panic(err)
	}

	sprite = entity.MakeSpriteComponent(p, "goblin")
	sprite.AddAnimation("idle", 0, 7, 80, false, mangaI.FLIP_NONE)
	sprite.AddAnimation("right", 1, 6, 80, false, mangaI.FLIP_NONE)
	sprite.AddAnimation("left", 1, 6, 80, false, mangaI.FLIP_HORIZONTAL)

	sprite.PlayAnimation("idle")

	p.AddComponent(mangaI.SpriteComponentID, sprite)

	p.AddComponent(mangaI.ScriptComponentID, entity.MakeScriptComponent(p, func() {
		if !keyboard.IsAnyKeyPressed(directionsKeys) {
			sprite.PlayAnimation("idle")
		} else {
			if keyboard.IsAnyKeyPressed(controls["right"]) {
				sprite.PlayAnimation("right")
			} else if keyboard.IsAnyKeyPressed(controls["left"]) {
				sprite.PlayAnimation("left")
			}
		}
	}))

	shape := collider.MakeRectangleShape(70, 80)
	collision := entity.MakeColliderComponent(p, shape, nil)
	collision.SetMargins(55, 50)
	p.AddComponent(mangaI.ColliderComponentID, collision)
}
