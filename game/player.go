package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	speed               = 0.5
	playerSize          = 105
	playerShootCoolDown = time.Millisecond * 250
)

type player struct {
	tex       *sdl.Texture
	x, y      float64
	lastShoot time.Time
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = createTextureFromBMP(renderer, "sprites/player.bmp")
	p.x = screenWidth / 2.0
	p.y = screenHeight - playerSize/2.0
	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= speed
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += speed
	}
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShoot) >= playerShootCoolDown {
			p.shoot(p.x+25, p.y-20)
			p.shoot(p.x-25, p.y-20)
			p.lastShoot = time.Now()
		}
	}
}

func (p *player) shoot(x, y float64) {
	if b, ok := bulletFromPool(); ok {
		b.active = true
		b.x = x
		b.y = y
		b.angle = 270 * (math.Pi / 180)
	}
}
