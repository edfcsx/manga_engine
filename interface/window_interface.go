package mangaI

import "manga_engine/vector"

type Window interface {
	Title() string
	Position() vector.Vec2[int32]
	Size() vector.Vec2[int32]
	SetTitle(title string)
	SetPosition(x, y int32)
	SetSize(x, y int32)
}
