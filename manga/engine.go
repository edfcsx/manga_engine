package manga

import (
	"errors"
	"github.com/edfcsx/manga_engine/collider"
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/edfcsx/manga_engine/keyboard"
	"github.com/edfcsx/manga_engine/texture"
	"github.com/edfcsx/manga_engine/vector"
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

/* scripts globais executados fora das entidades junto ao game loop */
var globalsScript []func()

/*
Se faz necessário criar a janela e o renderer na inicialização do pacote
caso ele seja criado depois e a primeira cena do jogo carregue alguma
textura, isso ira falhar, pois o ponteiro de render ainda não foi definido no
pacote texture.
*/
func init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	sdlWindow, err := sdl.CreateWindow(
		"",
		0,
		0,
		0,
		0,
		sdl.WINDOW_HIDDEN,
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
	texture.RendererPtr = renderer
}

func (m *manga) Initialize(window mangaI.Window, scene mangaI.Scene, fpsTarget uint) {
	defer m.Stop()

	m.frameTarget = fpsTarget
	m.frameTargetTime = 1000.0 / float64(fpsTarget)
	m.scene = scene
	position := window.Position()
	size := window.Size()

	m.window.SetSize(size.X, size.Y)

	if !window.IsResizable() {
		m.window.SetMinimumSize(size.X, size.Y)
		m.window.SetMaximumSize(size.X, size.Y)
	}

	window.SetGameWindow(m.window)
	m.running = true

	m.window.Show()
	m.window.SetPosition(position.X, position.Y)

	m.scene.Initialize()

	for m.running {
		for _, script := range globalsScript {
			script()
		}

		m.processEvents()
		m.update()
		m.render()
		collider.ResolveCollisions()
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
	m.scene.Update(m.deltaTime)
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

func (m *manga) GetRenderer() *sdl.Renderer {
	return m.renderer
}

func (m *manga) GetWindow() *sdl.Window {
	return m.window
}

func (m *manga) Draw(textureID string, src vector.Vec4[int32], dest vector.Vec4[int32], angle float64) error {
	srcRect := &sdl.Rect{X: src.X, Y: src.Y, W: src.W, H: src.H}
	destRect := &sdl.Rect{X: dest.X, Y: dest.Y, W: dest.W, H: dest.H}
	t := texture.GetTexture(textureID)

	if t == nil {
		return errors.New("engine: Texture not exists")
	}

	err := m.renderer.CopyEx(t.GetSource(), srcRect, destRect, angle, nil, sdl.FLIP_NONE)
	return err
}

func (m *manga) DrawTexture(texture mangaI.Texture, src vector.Vec4[int32], dest vector.Vec4[int32], angle float64, flip mangaI.FlipType) error {
	srcRect := &sdl.Rect{X: src.X, Y: src.Y, W: src.W, H: src.H}
	destRect := &sdl.Rect{X: dest.X, Y: dest.Y, W: dest.W, H: dest.H}

	err := m.renderer.CopyEx(texture.GetSource(), srcRect, destRect, angle, nil, sdl.RendererFlip(flip))
	return err
}

func (m *manga) AddGlobalScript(script func()) {
	globalsScript = append(globalsScript, script)
}

// check manga implements correct interface
var _ mangaI.Engine = (*manga)(nil)
