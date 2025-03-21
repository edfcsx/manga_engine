package manga

import (
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/veandco/go-sdl2/sdl"
)

type manga struct {
	window         *sdl.Window
	renderer       *sdl.Renderer
	running        bool
	deltaTime      float64
	ticksLastFrame uint64
	scene          mangaI.Scene
}

var Engine = &manga{
	window:         nil,
	renderer:       nil,
	running:        false,
	deltaTime:      0.0,
	ticksLastFrame: 0,
	scene:          nil,
}

func (m *manga) Initialize(window mangaI.Window, scene mangaI.Scene) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	defer m.Stop()

	m.scene = scene
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

	m.window = sdlWindow
	m.renderer = renderer
	m.running = true
	Engine.scene.Initialize()

	for m.running {
		// run global scripts
		// process events
		m.update()
		// systems update
		m.render()
	}

}

func (m *manga) Stop() {
	Engine.running = false
}

func (m *manga) GetDeltaTime() float64 {
	return m.deltaTime
}

func (m *manga) GetTicksLastFrame() uint64 {
	return m.ticksLastFrame
}

func (m *manga) update() {

}

func (m *manga) render() {

}
