package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 600
)

func main() {
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
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(screenWidth), int32(screenHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()
	pixels := make([]byte, screenWidth*screenHeight*4)
	imgTex := loadPNG("Sprites/Menu/Buttons/Back.png")
	imgTex.draw(vector{0, 0}, pixels)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		tex.Update(nil, pixels, screenWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(16)
	}
}

type vector struct {
	x, y float32
}

type texture struct {
	pixels      []byte
	w, h, pitch int
}

func (tex *texture) draw(pos vector, pixels []byte) {
	for y := 0; y < tex.h; y++ {
		for x := 0; x < tex.w; x++ {
			screenX := x + int(pos.x)
			screenY := y + int(pos.y)
			if screenX >= 0 && screenX < screenWidth && screenY >= 0 && screenY < screenHeight {
				texIndex := y*tex.pitch + x*4
				scrIndex := screenY*screenWidth*4 + screenX*4

				pixels[scrIndex] = tex.pixels[texIndex]
				pixels[scrIndex+1] = tex.pixels[texIndex+1]
				pixels[scrIndex+2] = tex.pixels[texIndex+2]
				pixels[scrIndex+3] = tex.pixels[texIndex+3]
			}

		}
	}
}

func loadPNG(filneName string) *texture {
	file, err := os.Open(filneName)
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	imgPixels := make([]byte, w*h*4)
	index := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			imgPixels[index] = byte(r / 256)
			index++
			imgPixels[index] = byte(g / 256)
			index++
			imgPixels[index] = byte(b / 256)
			index++
			imgPixels[index] = byte(a / 256)
			index++
		}
	}
	return &texture{pixels: imgPixels, w: w, h: h, pitch: w * 4}
}
