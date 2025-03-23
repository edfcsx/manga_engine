package mangaI

type Engine interface {
	Initialize(w Window, s Scene, fpsTarget uint)
	Stop()
	GetDeltaTime() float64
	GetTicksLastFrame() uint64
	GetFps() float64
}
