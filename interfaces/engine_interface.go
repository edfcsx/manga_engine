package mangaI

import "github.com/veandco/go-sdl2/sdl"

type Engine interface {
	Initialize(w Window, s Scene, fpsTarget uint)
	Stop()
	GetDeltaTime() float64
	GetTicksLastFrame() uint64
	GetFps() float64
	GetRenderer() *sdl.Renderer
	GetWindow() *sdl.Window
}
