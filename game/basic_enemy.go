package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySize = 105
)

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (be basicEnemy) {
	be.tex = createTextureFromBMP(renderer, "sprites/basic_enemy.bmp")
	be.x = x
	be.y = y
	return be
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	x := be.x - basicEnemySize/2
	y := be.y - basicEnemySize/2
	renderer.CopyEx(be.tex,
		&sdl.Rect{X: 0, Y: 0, W: basicEnemySize, H: basicEnemySize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemySize, H: basicEnemySize},
		180.0,
		&sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE)
}
