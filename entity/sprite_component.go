package entity

import (
	"errors"
	mangaI "github.com/edfcsx/manga_engine/interfaces"
	"github.com/edfcsx/manga_engine/manga"
	"github.com/edfcsx/manga_engine/texture"
	"github.com/edfcsx/manga_engine/vector"
	"github.com/veandco/go-sdl2/sdl"
)

type Animation struct {
	index     int32
	numFrames int32
	speed     int32
	isFixed   bool
}

type SpriteComponent struct {
	Entity        mangaI.Entity
	transform     mangaI.TransformComponent
	texture       mangaI.Texture
	src           vector.Vec4[int32]
	dst           vector.Vec4[int32]
	animations    map[string]*Animation
	currAnimation *Animation
}

func MakeSpriteComponent(e mangaI.Entity, textureID string) *SpriteComponent {
	t := texture.GetTexture(textureID)

	if t == nil {
		panic("sprite component: texture is nil")
	}

	transform := e.GetComponent(mangaI.TransformComponentID)

	if transform == nil {
		panic("sprite component: sprite component requires a transform component in entity")
	}

	return &SpriteComponent{
		Entity:     e,
		transform:  transform.(mangaI.TransformComponent),
		texture:    t,
		src:        vector.MakeVec4[int32](0, 0, 0, 0),
		dst:        vector.MakeVec4[int32](0, 0, 0, 0),
		animations: make(map[string]*Animation),
	}
}

func (s *SpriteComponent) GetType() string {
	return mangaI.SpriteComponentID
}

func (s *SpriteComponent) SetTexture(textureID string) error {
	text := texture.GetTexture(textureID)

	if text == nil {
		return errors.New("sprite component: texture is nil")
	}

	s.texture = text
	return nil
}

func (s *SpriteComponent) Update(deltaTime float64) {
	size := s.transform.GetSize()
	position := s.transform.GetPosition()
	scale := s.transform.GetScale()

	s.src.W = size.X
	s.src.H = size.Y

	if s.currAnimation != nil {
		s.src.X = s.src.W * ((int32(sdl.GetTicks64()) / s.currAnimation.speed) % s.currAnimation.numFrames)
		s.src.Y = s.currAnimation.index * s.src.H
	}

	s.dst.X = int32(position.X)
	s.dst.Y = int32(position.Y)
	s.dst.W = size.X * scale
	s.dst.H = size.Y * scale
}

func (s *SpriteComponent) Render() {
	err := manga.Engine.DrawTexture(s.texture, s.src, s.dst, 0.0)

	if err != nil {
		panic(err)
	}
}
