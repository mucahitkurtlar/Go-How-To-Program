package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	speed = 0.1
	playerSize = 105
)

type player struct {
	tex *sdl.Texture
	x, y float64
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		return player{}, fmt.Errorf("Loading player image: %v", err)
	}
	defer img.Free()

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("Creating player texture: %v", err)
	}

	p.x = screenWidth / 2 - playerSize / 2
	p.y = screenHeight - playerSize
	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - playerSize / 2
	y := p.y - playerSize / 2
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize},
	)
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()
	
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= speed
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += speed
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		p.y -= speed
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		p.y += speed
	}
}