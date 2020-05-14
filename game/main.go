package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 600
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(
		"GoGame",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	plr, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("Creating player: ", err)
		return
	}

	enemy, err := newBasicEnemy(renderer, screenHeight / 2, screenWidth / 2)
	if err != nil {
		fmt.Println("Creating basic enemy: ", err)
		return
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		enemy.draw(renderer)
		plr.draw(renderer)
		plr.update()
		renderer.Present()
	}
}
