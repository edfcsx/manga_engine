package mangaI

type Engine interface {
	Initialize(Window, Scene)
	Stop()
	AddGlobalScript(name string, script GlobalScript)
	GetDeltaTime() float64
	GetTicksLastFrame() uint64
}
