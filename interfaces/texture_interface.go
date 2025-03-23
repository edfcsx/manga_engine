package mangaI

import "github.com/veandco/go-sdl2/sdl"

type Texture interface {
	GetSource() *sdl.Texture
}
