package manga

import (
	"github.com/veandco/go-sdl2/sdl"
	mangaI "manga_engine/interface"
)

type manga struct {
	window         *sdl.Window
	renderer       *sdl.Renderer
	windowProps    mangaI.Window
	running        bool
	deltaTime      float64
	ticksLastFrame uint64
	scene          mangaI.Scene
	globalScript   map[string]mangaI.GlobalScript
}

var Engine = &manga{
	window:         nil,
	renderer:       nil,
	windowProps:    nil,
	running:        false,
	deltaTime:      0.0,
	ticksLastFrame: 0,
	scene:          nil,
	globalScript:   map[string]mangaI.GlobalScript{},
}

func (e *manga) Initialize(window mangaI.Window, scene mangaI.Scene) {
	Engine.windowProps = window
	Engine.scene = scene
	defer e.Stop()

	position := window.Position()
	size := window.Size()

	sdlWindow, err := sdl.CreateWindow(
		window.Title(),
		position.X,
		position.Y,
		size.X,
		size.Y,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(sdlWindow, -1, 0)

	if err != nil {
		panic(err)
	}

	Engine.window = sdlWindow
	Engine.renderer = renderer
	Engine.running = true
	Engine.scene.Initialize()

	for Engine.running {
		// run global scripts
		// process events
		e.update()
		// systems update
		e.render()
	}
}

func (e *manga) Stop() {
	Engine.running = false
}

func (e *manga) AddGlobalScript(name string, script mangaI.GlobalScript) {
	Engine.globalScript[name] = script
}

func (e *manga) GetDeltaTime() float64 {
	return e.deltaTime
}

func (e *manga) GetTicksLastFrame() uint64 {
	return e.ticksLastFrame
}

func (e *manga) update() {

}

func (e *manga) render() {

}
