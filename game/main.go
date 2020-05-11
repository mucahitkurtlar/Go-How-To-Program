package main

import (
  "fmt"
  "github.com/veandco/go-sdl2/sdl"
)

const (
  screenWidth = 600
  screenHeight = 800
)

func main () {
  if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
    panic(err)
  }

  window, err := sdl.CreateWindow(
		"Gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
  defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()
  for {
    renderer.SetDrawColor(255, 0, 0, 255)
    renderer.Clear()
    renderer.Present()
  }
}
