package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
)

const (
	WIDTH            = 320
	HEIGHT           = 240
	FRAMEBUFFER_SIZE = WIDTH * HEIGHT * 4
)

var framebuffer = make([]byte, FRAMEBUFFER_SIZE)
var renderer *sdl.Renderer

func run() int {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		panic(err)
	}
	defer renderer.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	defer surface.Free()

	fmt.Printf("%v\n", surface)

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, WIDTH, HEIGHT)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	for i := 0; i < FRAMEBUFFER_SIZE; i += 4 {
		framebuffer[i] = 255
		framebuffer[i + 1] = 255
		framebuffer[i + 2] = 255
		framebuffer[i + 3] = 255
		draw(texture)
	}

	sdl.Delay(2000)

	return 0
}

func draw(texture *sdl.Texture) {
	texture.Update(nil, framebuffer, WIDTH*4)
	renderer.Clear()
	renderer.Copy(texture, nil, nil)
	renderer.Present()
}

func main() {
	runtime.LockOSThread()
	os.Exit(run())
}
