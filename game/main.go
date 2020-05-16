package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 600
)

func createTextureFromBMP(renderer *sdl.Renderer, fileName string) *sdl.Texture {
	img, err := sdl.LoadBMP(fileName)
	if err != nil {
		panic(fmt.Errorf("Loading %v: %v", fileName, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("Creating texture %v: %v", fileName, err))
	}
	return tex
}

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

	plr := newPlayer(renderer)
	/*
		enemy, err := newBasicEnemy(renderer, screenHeight/2, screenWidth/2)
		if err != nil {
			fmt.Println("Creating basic enemy: ", err)
			return
		}
	*/

	var enemies []basicEnemy

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0)

			enemy := newBasicEnemy(renderer, x, y)
			enemies = append(enemies, enemy)
		}
	}
	initBulletPool(renderer)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		//enemy.draw(renderer)
		for _, enemy := range enemies {
			enemy.draw(renderer)
		}
		plr.draw(renderer)
		plr.update()
		for _, bul := range bulletPool {
			bul.draw(renderer)
			bul.update()
		}
		renderer.Present()
	}
}
