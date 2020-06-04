package app

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Game interface {
	LoadTexture(filename string) (*sdl.Texture, error)
}

type game struct {
	window            *sdl.Window
	renderer          *sdl.Renderer
	scenes            []Scene
	currentSceneIndex int
	textures          []*sdl.Texture
	titleTexture      *sdl.Texture
	titleSrc          sdl.Rect
	titleDst          sdl.Rect
}

func NewGame(title string, screenWidth, screenHeight int32) (*game, error) {
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.Wrap(err, "error on sdl create window")
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, errors.Wrap(err, "error on sdl create renderer")
	}

	font1, err := ttf.OpenFont("resources/Arcade Classic.ttf", 16)
	if err != nil {
		return nil, errors.Wrap(err, "error on open font")
	}

	white := sdl.Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	surface, err := font1.RenderUTF8Solid(title, white)
	if err != nil {
		return nil, errors.Wrap(err, "error on font render utf-8 solid")
	}

	titleTexture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, errors.Wrap(err, "error on create texture from surface")
	}

	_, _, w, h, err := titleTexture.Query()
	if err != nil {
		return nil, errors.Wrap(err, "error on query title texture")
	}

	titleSrc := sdl.Rect{
		X: 0,
		Y: 0,
		W: w,
		H: h,
	}

	titleDst := sdl.Rect{
		X: 16,
		Y: 16,
		W: w,
		H: h,
	}

	game := game{
		window:            window,
		renderer:          renderer,
		scenes:            make([]Scene, 0),
		currentSceneIndex: -1,
		textures:          make([]*sdl.Texture, 0),
		titleTexture:      titleTexture,
		titleSrc:          titleSrc,
		titleDst:          titleDst,
	}

	return &game, nil
}

func (game *game) Close() error {
	var err error

	for i := range game.textures {
		err = game.textures[i].Destroy()
		if err != nil {
			return errors.Wrap(err, "error on destroy texture")
		}
	}

	err = game.renderer.Destroy()
	if err != nil {
		return errors.Wrap(err, "error on destroy renderer")
	}

	err = game.window.Destroy()
	if err != nil {
		return errors.Wrap(err, "error on destroy window")
	}

	return nil
}

func (game *game) Run() error {
	var (
		now       = sdl.GetPerformanceCounter()
		last      uint64
		deltaTime float32
	)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return nil
			}
		}

		last = now
		now = sdl.GetPerformanceCounter()
		deltaTime = float32((now - last) / sdl.GetPerformanceFrequency())

		err := game.update(deltaTime)
		if err != nil {
			return errors.Wrap(err, "error on update game")
		}

		err = game.render()
		if err != nil {
			return errors.Wrap(err, "error on render game")
		}
	}
}

func (game *game) update(deltaTime float32) error {
	currentScene := game.CurrentScene()

	if currentScene == nil {
		return nil
	}

	err := currentScene.Update(deltaTime)
	if err != nil {
		return errors.Wrap(err, "error on update current scene")
	}

	return nil
}

func (game *game) render() error {
	err := game.renderer.Clear()
	if err != nil {
		return errors.Wrap(err, "error on clear game renderer")
	}

	currentScene := game.CurrentScene()

	if currentScene == nil {
		return nil
	}

	err = currentScene.Render(game.renderer)
	if err != nil {
		return errors.Wrap(err, "error on render current scene")
	}

	err = game.renderer.Copy(game.titleTexture, &game.titleSrc, &game.titleDst)
	if err != nil {
		return errors.Wrap(err, "error on copy title texture to renderer")
	}

	game.renderer.Present()

	return nil
}

func (game *game) CurrentScene() Scene {
	if game.currentSceneIndex == -1 || len(game.scenes) == 0 {
		return nil
	}

	if game.currentSceneIndex >= len(game.scenes) {
		return nil
	}

	return game.scenes[game.currentSceneIndex]
}

func (game *game) AddScene(scene *scene) {
	game.scenes = append(game.scenes, scene)
	if game.currentSceneIndex == -1 {
		game.currentSceneIndex = 0
	}
}

func (game *game) LoadTexture(filename string) (*sdl.Texture, error) {
	texture, err := img.LoadTexture(game.renderer, filename)
	if err != nil {
		return nil, errors.Wrap(err, "error on load texture")
	}

	game.textures = append(game.textures, texture)

	return texture, nil
}
