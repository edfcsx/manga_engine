package texture

import (
	"github.com/edfcsx/manga_engine/manga"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	Width  int32
	Height int32
	source *sdl.Texture
}

func Draw(texture *Texture, src *sdl.Rect, dst *sdl.Rect, flip sdl.RendererFlip) {
	err := manga.Engine.GetRenderer().CopyEx(texture, src, dst, 0, nil, flip)

	if err != nil {
		panic(err)
	}
}

func MakeTexture(filePath string) *Texture {
	surface, err := img.Load(filePath)

	if err != nil {
		panic(err)
	}

	texture, err := manga.Engine.GetRenderer().CreateTextureFromSurface(surface)

	if err != nil {
		panic(err)
	}

	text := &Texture{
		Width:  surface.W,
		Height: surface.H,
		source: texture,
	}

	surface.Free()
	return text
}
