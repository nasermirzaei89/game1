package app

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene interface {
	Render(renderer *sdl.Renderer) error
	Update(deltaTime float32) error
}

type scene struct {
	entities []Entity
}

func NewScene() *scene {
	res := scene{
		entities: make([]Entity, 0),
	}

	return &res
}

func (s *scene) AddEntity(entity Entity) {
	s.entities = append(s.entities, entity)
}

func (s *scene) Update(deltaTime float32) error {
	for i := range s.entities {
		err := s.entities[i].Update(deltaTime)
		if err != nil {
			return errors.Wrap(err, "error on update scene entity")
		}
	}

	return nil
}

func (s *scene) Render(renderer *sdl.Renderer) error {
	for i := range s.entities {
		err := s.entities[i].Render(renderer)
		if err != nil {
			return errors.Wrap(err, "error on render scene entity")
		}
	}

	return nil
}
