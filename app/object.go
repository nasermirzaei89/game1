package app

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Object interface {
}

type object struct {
	sprite Sprite
	x      float32
	y      float32
	hSpeed float32
	vSpeed float32
}

func NewObject(sprite Sprite) *object {
	res := object{
		sprite: sprite,
		x:      0,
		y:      0,
		hSpeed: 0,
		vSpeed: 0,
	}

	return &res
}

func (obj object) GetSprite() Sprite {
	return obj.sprite
}

func (obj object) GetSrc() sdl.Rect {
	return obj.sprite.GetSrc()
}

func (obj object) GetDst() sdl.Rect {
	dst := obj.sprite.GetDst()
	dst.X += int32(obj.x)
	dst.Y += int32(obj.y)

	return dst
}

func (obj *object) SetPosition(x, y float32) {
	obj.x = x
	obj.y = y
}

func (obj *object) SetVelocity(hSpeed, vSpeed float32) {
	obj.hSpeed = hSpeed
	obj.vSpeed = vSpeed
}

func (obj *object) Update(deltaTime float32) error {
	obj.x += obj.hSpeed * deltaTime
	obj.y += obj.vSpeed * deltaTime

	return nil
}

func (obj *object) Render(renderer *sdl.Renderer) error {
	src := obj.GetSrc()
	dst := obj.GetDst()
	texture := obj.GetSprite().GetTexture()
	err := renderer.Copy(texture, &src, &dst)
	if err != nil {
		return errors.Wrap(err, "error on copy texture to renderer")
	}

	return nil
}
