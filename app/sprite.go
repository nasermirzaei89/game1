package app

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite interface {
	GetSrc() sdl.Rect
	GetDst() sdl.Rect
	GetTexture() *sdl.Texture
}

type sprite struct {
	texture *sdl.Texture
	src     sdl.Rect
	dst     sdl.Rect
}

func NewSprite(texture *sdl.Texture) (*sprite, error) {
	res := sprite{
		texture: texture,
		src:     sdl.Rect{},
		dst:     sdl.Rect{},
	}

	_, _, w, h, err := texture.Query()
	if err != nil {
		return nil, errors.Wrap(err, "error on query texture")
	}

	res.src.W = w
	res.src.H = h

	res.dst.W = w
	res.dst.H = h

	return &res, nil
}

func (s sprite) GetTexture() *sdl.Texture {
	return s.texture
}

func (s sprite) GetSrc() sdl.Rect {
	return s.src
}

func (s sprite) GetDst() sdl.Rect {
	return s.dst
}

func (s *sprite) SetFrame(frame sdl.Rect) {
	s.src = frame
	s.dst = frame
	s.dst.X = 0
	s.dst.Y = 0
}
