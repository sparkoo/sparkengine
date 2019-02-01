package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"reflect"
	"runtime"
	"time"
)

const (
	WIDTH            = 640
	HEIGHT           = 480
	FRAMEBUFFER_SIZE = WIDTH * HEIGHT * 4
)

var framebuffer = make([]byte, FRAMEBUFFER_SIZE)
var renderer *sdl.Renderer

var running = true
var ball *Ball

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

	ball = NewBall(0, 0, 10, 10, 1.2, 1.5)

	err = sdl.WarpMouseGlobal(WIDTH / 2, HEIGHT / 2)
	if err != nil {
		panic(err)
	}

	sdl.SetRelativeMouseMode(true)

	for running {
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

		if event := sdl.PollEvent(); event != nil {
			handleEvent(event)
		}

		draw(texture)
	}

	sdl.SetRelativeMouseMode(false)

	sdl.Delay(200)

	return 0
}

func handleEvent(event sdl.Event) {
	switch t := event.(type) {
	case *sdl.MouseMotionEvent:
		fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
		ball.xpos += float64(t.XRel)
		if ball.xpos <= 0 {
			ball.xpos = 0
		}
		if ball.xpos + float64(ball.xsize) >= WIDTH {
			ball.xpos = float64(WIDTH - ball.xsize)
		}
		ball.ypos += float64(t.YRel)
		if ball.ypos <= 0 {
			ball.ypos = 0
		}
		if ball.ypos + float64(ball.ysize) >= HEIGHT {
			ball.ypos = float64(HEIGHT - ball.ysize)
		}
		fmt.Println("[x:", ball.xpos, "; y:", ball.ypos, "]")
	case *sdl.MouseButtonEvent:
		fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
	case *sdl.MouseWheelEvent:
		fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
			t.Timestamp, t.Type, t.Which, t.X, t.Y)
	case *sdl.KeyboardEvent:
		fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%dscancode:%d\n",
			t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat, t.Keysym.Scancode)
		if t.Keysym.Scancode == 41 {
			running = false
		}
	case *sdl.UserEvent:
		fmt.Printf("[%d ms] UserEvent\tcode:%d\n", t.Timestamp, t.Code)
	default:
		fmt.Printf("Unknown event: [%v] - [%v]\n", event, reflect.TypeOf(event).String())
	}
}

func draw(texture *sdl.Texture) {
	texture.Update(nil, framebuffer, WIDTH*4)
	renderer.Clear()
	renderer.Copy(texture, nil, nil)
	renderer.Present()
	handlefps()

	//time.Sleep(15 * time.Millisecond)
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

	fmt.Println("width:", WIDTH, "; height:", HEIGHT, "; pixels:", WIDTH*HEIGHT)

	os.Exit(run())
}
