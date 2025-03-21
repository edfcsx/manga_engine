package mangaI

type Engine interface {
	Initialize(Window, Scene)
	Stop()
	GetDeltaTime() float64
	GetTicksLastFrame() uint64
}
