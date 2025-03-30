package mangaI

import (
	"github.com/edfcsx/manga_engine/vector"
	"github.com/veandco/go-sdl2/sdl"
)

type Engine interface {
	Initialize(w Window, s Scene, fpsTarget uint)
	Stop()
	GetDeltaTime() float64
	GetTicksLastFrame() uint64
	GetFPS() float64
	GetRenderer() *sdl.Renderer
	GetWindow() *sdl.Window
	Draw(textureID string, src vector.Vec4[int32], dest vector.Vec4[int32], angle float64) error
	DrawTexture(textureID Texture, src vector.Vec4[int32], dest vector.Vec4[int32], angle float64, flip FlipType) error
}
