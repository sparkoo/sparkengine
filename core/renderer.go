package core

import (
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"runtime"
)

type renderer interface {
	init(conf *Conf)
	renderFrame(objects []scene.Object)
	destroy()
}

type sdlRenderer struct {
	framebuffer     []byte
	framebufferSize int
	conf *Conf

	texture *sdl.Texture
	window *sdl.Window
	renderer *sdl.Renderer
}

func (r *sdlRenderer) init(conf *Conf) {
	log.Println("initializing SDL renderer ...")

	r.conf = conf

	r.framebufferSize = int(conf.screenWidth * conf.screenHeight * 4)
	r.framebuffer = make([]byte, r.framebufferSize)

	initSDL()
	r.window = initWindow(conf)
	r.renderer = initRenderer(r.window)
	r.texture = initTexture(r.renderer, conf)
	r.texture.Update(nil, r.framebuffer, int(conf.screenWidth * 4))

	sdl.SetRelativeMouseMode(true)
	log.Println("done")
}

func (r *sdlRenderer) destroy() {
	log.Println("cleaning up ...")
	sdl.SetRelativeMouseMode(false)
	r.window.Destroy()
	r.renderer.Destroy()
	r.texture.Destroy()
	sdl.Quit()
	log.Println("done")
}

func (r *sdlRenderer) renderFrame(objects []scene.Object) {
	framebuffer := make([]byte, r.framebufferSize)
	for _, o := range objects {
		for _, p := range o.GetPixels() {
			x := o.GetXoffset() + p.X
			y := o.GetYoffset() + p.Y
			i := (x + (int(r.conf.screenWidth) * y)) * 4

			// fit pixel into the framebuffer ?
			if i + 3 <= r.framebufferSize {
				framebuffer[i] = p.Color[0]
				framebuffer[i+1] = p.Color[1]
				framebuffer[i+2] = p.Color[2]
				framebuffer[i+3] = p.Color[3]
			}
		}
	}

	r.texture.Update(nil, framebuffer, int(r.conf.screenWidth * 4))
	r.renderer.Clear()
	r.renderer.Copy(r.texture, nil, nil)
	r.renderer.Present()
}

func initSDL() {
	runtime.LockOSThread()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func initWindow(conf *Conf) *sdl.Window {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		conf.screenWidth, conf.screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	return window
}

func initRenderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	return renderer
}

func initTexture(renderer *sdl.Renderer, conf *Conf) *sdl.Texture {
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING,
		conf.screenWidth, conf.screenHeight)
	if err != nil {
		panic(err)
	}
	return texture
}

