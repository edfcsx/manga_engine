package collider

import mangaI "github.com/edfcsx/manga_engine/interfaces"

var showCollisionBox = true

type Collider struct {
	shape       mangaI.ColliderShape
	onCollision func()
}

func SetShowCollisionBoxs(status bool) {
	showCollisionBox = status
}

func ShowCollisionBoxs() bool {
	return showCollisionBox
}
