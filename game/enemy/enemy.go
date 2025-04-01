package enemy

import (
	"github.com/edfcsx/manga_engine/collider"
	"github.com/edfcsx/manga_engine/entity"
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/edfcsx/manga_engine/texture"
)

type Enemy struct {
	entity.Entity
}

func MakeEnemy() *Enemy {
	return &Enemy{}
}

func (e *Enemy) Initialize() {
	e.Label = "Enemy"
	e.SetIsActive(true)

	transform := entity.MakeTransformComponent(e)
	transform.Position(600, 200)
	transform.Size(192, 192)
	transform.Scale(3)

	e.AddComponent(mangaI.TransformComponentID, transform)

	err := texture.MakeTexture("warrior", "./assets/warrior.png")

	if err != nil {
		panic(err)
	}

	sprite := entity.MakeSpriteComponent(e, "warrior")
	sprite.AddAnimation("idle", 0, 6, 80, false, mangaI.FLIP_NONE)
	sprite.PlayAnimation("idle")
	e.AddComponent(mangaI.SpriteComponentID, sprite)

	shape := collider.MakeRectangleShape(70, 80)
	collision := entity.MakeColliderComponent(e, shape, mangaI.ColliderMoving, nil)
	collision.SetMargins(55, 50)
	e.AddComponent(mangaI.ColliderComponentID, collision)
}
