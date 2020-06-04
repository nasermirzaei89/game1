package main

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

func main() {
	log.Println("Initializing")

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalln(errors.Wrap(err, "error on init sdl"))
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Game 1", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1024, 768, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer func() { _ = window.Destroy() }()

L1:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				log.Println("Quit Event")
				break L1
			}
		}
	}

	log.Println("Finalizing")
}
