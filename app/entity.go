package app

import "github.com/veandco/go-sdl2/sdl"

type Entity interface {
	Update(deltaTime float32) error
	Render(renderer *sdl.Renderer) error
}
