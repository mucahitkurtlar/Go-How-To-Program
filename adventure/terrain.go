package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	terrainSize = 48
)

type terrain struct {
	tex  *sdl.Texture
	rect rectangle
}

func newTerrain(renderer *sdl.Renderer) (t terrain) {
	t.tex = textureFromBMP(renderer, "Sprites/Terrain/Terrain.bmp")
	t.rect.width = terrainSize * 13
	t.rect.height = terrainSize
	t.rect.x = screenWidth / 2
	t.rect.y = screenHeight - terrainSize/2
	return t
}

func (t *terrain) draw(renderer *sdl.Renderer) {
	for i := 0; i < 14; i++ {
		x := i * terrainSize
		y := screenHeight - terrainSize
		renderer.CopyEx(t.tex,
			&sdl.Rect{X: 96, Y: 0, W: terrainSize, H: terrainSize},
			&sdl.Rect{X: int32(x), Y: int32(y), W: terrainSize, H: terrainSize},
			0,
			&sdl.Point{X: terrainSize / 2, Y: terrainSize / 2},
			sdl.FLIP_NONE)

	}
}
