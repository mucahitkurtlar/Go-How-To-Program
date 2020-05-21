package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerJumpCooldown = time.Millisecond * 200
	playerSize         = 32
	playerSpeed        = 0.2
	playerJumpSpeed    = 0.1
	right              = sdl.FLIP_NONE
	left               = sdl.FLIP_HORIZONTAL
)

type player struct {
	renderer  *sdl.Renderer
	x, y      float64
	tex       *sdl.Texture
	frame     uint16
	direction sdl.RendererFlip
	lastJump  time.Time
	jump      uint16
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromBMP(renderer, "Sprites/MainCharacters/MaskDude/Run.bmp")
	p.renderer = renderer
	p.x = screenWidth / 2.0
	p.y = screenHeight - playerSize/2.0
	p.frame = 0
	p.direction = right
	p.jump = 0
	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0

	renderer.CopyEx(p.tex,
		&sdl.Rect{X: playerSize * (int32(p.frame) % 12), Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize},
		0,
		&sdl.Point{X: playerSize / 2, Y: playerSize / 2},
		p.direction)

}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if p.x-(playerSize/2.0) > 0 {
			p.x -= playerSpeed
			p.frame++
			p.direction = left
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if p.x+(playerSize/2.0) < screenWidth {
			p.x += playerSpeed
			p.frame++
			p.direction = right
		}
	} else {
		p.frame = 0
	}
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastJump) >= playerJumpCooldown {
			p.jump = 100
			p.lastJump = time.Now()
		}
	}
	if p.jump > 0 {
		p.y--
		p.jump--
	}
}
