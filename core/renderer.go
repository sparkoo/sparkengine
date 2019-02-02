package core

import (
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type renderer interface {
	init(conf *conf)
	renderFrame(objects []scene.Object)
	destroy()
}

type sdlRenderer struct {
	texture *sdl.Texture
	window *sdl.Window
	renderer *sdl.Renderer
}

func (r sdlRenderer) destroy() {
	log.Println("cleaning up ...")
	r.window.Destroy()
	r.renderer.Destroy()
	r.texture.Destroy()
	sdl.Quit()
	log.Println("done")
}

func (sdlRenderer) renderFrame(objects []scene.Object) {
	//log.Println("frame rendered")
}

func (r sdlRenderer) init(conf *conf) {
	log.Println("initializing SDL renderer ...")
	initSDL()
	window := initWindow(conf)
	renderer := initRenderer(window)
	r.texture = initTexture(renderer, conf)
	log.Println("done")
}

func initSDL() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func initWindow(conf *conf) *sdl.Window {
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

func initTexture(renderer *sdl.Renderer, conf *conf) *sdl.Texture {
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING,
		conf.screenWidth, conf.screenHeight)
	if err != nil {
		panic(err)
	}
	return texture
}

