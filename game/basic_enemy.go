package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 105
)

type basicEnemy struct {
	tex *sdl.Texture
	x, y float64 
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (be basicEnemy, err error) {
	img, err := sdl.LoadBMP("sprites/basic_enemy.bmp")
	if err != nil {
		return basicEnemy{}, fmt.Errorf("Loading basic enemy image: %v", err)
	}
	defer img.Free()

	be.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return basicEnemy{}, fmt.Errorf("Creating basic enemy texture: %v", err)
	}
	be.x = x
	be.y = y
	return be, nil
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	x := be.x - basicEnemySize / 2
	y := be.y - basicEnemySize / 2
	renderer.Copy(be.tex,
		&sdl.Rect{X: 0, Y:0, W: basicEnemySize, H: basicEnemySize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemySize, H: basicEnemySize},
	)
}