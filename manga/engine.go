package manga

import (
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/edfcsx/manga_engine/keyboard"
	"github.com/veandco/go-sdl2/sdl"
)

const FPS_UNLIMITED = 10001

type UpdateData struct {
	currentTicks uint64
	timeElapsed  uint64
	timeToWait   uint64
}

type manga struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	running         bool
	deltaTime       float64
	ticksLastFrame  uint64
	scene           mangaI.Scene
	frameTarget     uint
	frameTargetTime float64
	fps             *fpsCounter
}

var Engine = &manga{
	window:          nil,
	renderer:        nil,
	running:         false,
	deltaTime:       0.0,
	ticksLastFrame:  0,
	scene:           nil,
	frameTarget:     0,
	frameTargetTime: 0.0,
	fps:             makeFpsCounter(),
}

func (m *manga) Initialize(window mangaI.Window, scene mangaI.Scene, fpsTarget uint) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	defer m.Stop()

	m.frameTarget = fpsTarget
	m.frameTargetTime = 1000.0 / float64(fpsTarget)

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

	window.SetGameWindow(sdlWindow)

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
		m.processEvents()
		m.update()
		m.render()
	}
}

func (m *manga) Stop() {
	Engine.running = false

	if m.renderer != nil {
		err := m.renderer.Destroy()

		panic(err)
	}

	if m.window != nil {
		err := m.window.Destroy()

		if err != nil {
			panic(err)
		}
	}
}

func (m *manga) GetDeltaTime() float64 {
	return m.deltaTime
}

func (m *manga) GetTicksLastFrame() uint64 {
	return m.ticksLastFrame
}

// variables used in update - reduces memory allocation
var updt = &UpdateData{}

func (m *manga) update() {
	updt.currentTicks = sdl.GetTicks64()
	updt.timeElapsed = updt.currentTicks - m.ticksLastFrame

	if m.frameTarget != FPS_UNLIMITED {
		updt.timeToWait = uint64(m.frameTargetTime) - updt.timeElapsed

		if updt.timeToWait > 0 && updt.timeToWait <= 1000 {
			sdl.Delay(uint32(updt.timeToWait))
			updt.currentTicks = sdl.GetTicks64()
		}
	}

	m.deltaTime = float64(updt.timeElapsed) / 1000.0

	if m.deltaTime > 0.05 {
		m.deltaTime = 0.05
	}

	m.ticksLastFrame = updt.currentTicks
	m.scene.Update()
	m.fps.Update()
}

func (m *manga) render() {
	err := m.renderer.SetDrawColor(0, 0, 0, 255)

	if err != nil {
		panic(err)
	}

	err = m.renderer.Clear()
	if err != nil {
		panic(err)
	}

	m.scene.Render()
	m.renderer.Present()
}

func (m *manga) processEvents() {
	for {
		event := sdl.PollEvent()

		if event == nil {
			break
		}

		switch event.(type) {
		case *sdl.QuitEvent:
			m.Stop()
		case *sdl.KeyboardEvent:
			keyEvent := event.(*sdl.KeyboardEvent)
			keyCode := int(keyEvent.Keysym.Sym)

			if keyEvent.Type == sdl.KEYDOWN {
				keyboard.RegisterKeyPressed(keyCode)
			} else if keyEvent.Type == sdl.KEYUP {
				keyboard.RegisterKeyReleased(keyCode)
			}
		}
	}
}

func (m *manga) GetFPS() float64 {
	return m.fps.GetFPS()
}
