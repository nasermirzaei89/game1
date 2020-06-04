package main

import (
	"github.com/nasermirzaei89/game1/app"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
)

func main() {
	log.Println("Initializing")

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error on init sdl"))
	}
	defer sdl.Quit()

	err = img.Init(img.INIT_PNG)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error on init img"))
	}
	defer img.Quit()

	err = ttf.Init()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error on init ttf"))
	}
	defer ttf.Quit()

	game1, err := app.NewGame("Game 1", 1024, 768)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error on new game"))
	}
	defer func() { _ = game1.Close() }()

	scene1 := app.NewScene()
	game1.AddScene(scene1)

	tileMap1, err := app.NewTileMap("resources/map1.tmx", game1)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error on load texture"))
	}
	scene1.AddEntity(tileMap1)

	texture1, err := game1.LoadTexture("resources/Human Buildings Summer.png")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error on load texture"))
	}

	sprite1, err := app.NewSprite(texture1)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error on new sprite"))
	}
	sprite1.SetFrame(sdl.Rect{
		X: 0,
		Y: 0,
		W: 125,
		H: 125,
	})

	object1 := app.NewObject(sprite1)
	object1.SetPosition(480, 240)
	scene1.AddEntity(object1)

	sprite2, err := app.NewSprite(texture1)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error on new sprite"))
	}
	sprite2.SetFrame(sdl.Rect{
		X: 390,
		Y: 0,
		W: 75,
		H: 75,
	})

	object2 := app.NewObject(sprite2)
	object2.SetPosition(640, 320)
	scene1.AddEntity(object2)

	object3 := app.NewObject(sprite2)
	object3.SetPosition(640, 200)
	scene1.AddEntity(object3)

	err = game1.Run()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error on run game"))
	}

	log.Println("Finalizing")
}
