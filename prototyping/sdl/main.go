package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
	"time"
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

	ball := NewBall(0, 0, 10, 10, 1.2, 1.5)

	for {
		framebuffer = make([]byte, FRAMEBUFFER_SIZE)
		for _, p := range ball.getPixels() {
			x := ball.getXoffset() + p.x
			y := ball.getYoffset() + p.y
			i := (x + (WIDTH * y)) * 4
			framebuffer[i] = p.color[0]
			framebuffer[i+1] = p.color[1]
			framebuffer[i+2] = p.color[2]
			framebuffer[i+3] = p.color[3]
		}

		xPot := int(ball.xpos + ball.xvel)
		if xPot < 0 || xPot+ball.xsize >= WIDTH {
			ball.xvel *= -1
		}

		yPot := int(ball.ypos + ball.yvel)
		if yPot < 0 || yPot+ball.ysize >= HEIGHT {
			ball.yvel *= -1
		}

		ball.xpos += ball.xvel
		ball.ypos += ball.yvel
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
	handlefps()

	time.Sleep(15 * time.Millisecond)
}

var frames = 0
var t1 = time.Now()

func handlefps() {
	frames++
	if frames >= 1000 {
		duration := float64(time.Now().Sub(t1).Nanoseconds()) / 1000 / 1000
		fps := int(float64(frames) / (duration / 1000))
		fmt.Println(fps, "fps")
		t1 = time.Now()
		frames = 0
	}
}

func main() {
	runtime.LockOSThread()
	os.Exit(run())
}
