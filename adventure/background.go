package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	backgroundSize = 64
)

type background struct {
	tex *sdl.Texture
}

func newBackground(renderer *sdl.Renderer) (bg background) {
	bg.tex = textureFromBMP(renderer, "Sprites/Background/Purple.bmp")
	return bg
}

func (bg *background) draw(renderer *sdl.Renderer) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			x := j * backgroundSize
			y := i * backgroundSize

			renderer.CopyEx(bg.tex,
				&sdl.Rect{X: 0, Y: 0, W: backgroundSize, H: backgroundSize},
				&sdl.Rect{X: int32(x), Y: int32(y), W: backgroundSize, H: backgroundSize},
				0,
				&sdl.Point{X: backgroundSize / 2, Y: backgroundSize / 2},
				sdl.FLIP_NONE)
		}
	}
}
