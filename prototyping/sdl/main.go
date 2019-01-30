package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"time"
)

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

const (
	HEIGHT           = 480
	WIDTH            = 640
	FRAMEBUFFER_SIZE = WIDTH * HEIGHT
)

var framebuffer = make([]Pixel, FRAMEBUFFER_SIZE)
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

	for i := range framebuffer{
		framebuffer[i] = Pixel{255, 255, 255, 255}
		draw()
	}

	sdl.Delay(20000)

	return 0
}

func draw() {
	t1 := time.Now()
	renderer.Clear()
	for i, p := range framebuffer {
		renderer.SetDrawColor(p.R, p.G, p.B, p.A)
		x := i % WIDTH
		y := i / WIDTH
		renderer.DrawPoint(int32(x), int32(y))
	}
	renderer.Present()
	t2 := time.Now().Sub(t1)
	fmt.Printf("frame rendered in %dms\n", t2.Nanoseconds() / 1000 / 1000)
}

func main() {
	os.Exit(run())
}
