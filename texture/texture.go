package texture

import (
	"errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var RendererPtr *sdl.Renderer = nil
var textures = map[string]*Texture{}

type Texture struct {
	Width  int32
	Height int32
	source *sdl.Texture
}

func (t *Texture) GetSource() *sdl.Texture {
	return t.source
}

func GetTexture(id string) *Texture {
	return textures[id]
}

func MakeTexture(id string, filePath string) error {
	surface, err := img.Load(filePath)

	if err != nil {
		return err
	}

	if RendererPtr == nil {
		panic(errors.New("texture: RendererPtr not initialized"))
	}

	texture, err := RendererPtr.CreateTextureFromSurface(surface)

	if err != nil {
		return err
	}

	textures[id] = &Texture{
		Width:  surface.W,
		Height: surface.H,
		source: texture,
	}

	surface.Free()
	return nil
}
