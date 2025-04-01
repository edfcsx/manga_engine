package collider

import mangaI "github.com/edfcsx/manga_engine/interfaces"

var showCollisionBox = true

var static = make(map[string]*Collider)
var moving = make(map[string]*Collider)

type Collider struct {
	label       string
	shape       mangaI.ColliderShape
	onCollision func(label string)
}

func SetShowCollisionBoxs(status bool) {
	showCollisionBox = status
}

func ShowCollisionBoxs() bool {
	return showCollisionBox
}

func ResolveCollisions() {
	if len(moving) <= 1 && len(static) == 0 {
		return
	}

	for idx1, moving1 := range moving {
		for idx2, moving2 := range moving {
			if idx1 != idx2 && moving1.shape.CollidesWith(moving2.shape) {
				if moving1.onCollision != nil {
					moving1.onCollision(moving2.label)
				}

				if moving2.onCollision != nil {
					moving2.onCollision(moving1.label)
				}
			}
		}
	}

	for _, movingObject := range moving {
		for _, staticObject := range static {
			if movingObject.shape.CollidesWith(staticObject.shape) {
				if movingObject.onCollision != nil {
					movingObject.onCollision(staticObject.label)
				}
			}
		}
	}
}

func AddCollision(id string, label string, shape mangaI.ColliderShape, onCollision func(label string), t mangaI.ColliderType) {
	collision := &Collider{
		label:       label,
		shape:       shape,
		onCollision: onCollision,
	}

	if t == mangaI.ColliderStatic {
		static[id] = collision
	} else {
		moving[id] = collision
	}
}
