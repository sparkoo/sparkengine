package core

import (
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

type game struct {
	running bool
	conf    *Conf

	currentScene *scene.Scene

	gameTicker  *time.Ticker
	frameTicker *time.Ticker
}

func NewGame(conf *Conf) *game {
	return &game{running: false, conf: conf}
}

func (g *game) Start(s *scene.Scene) {
	renderer := &sdlRenderer{}
	renderer.init(g.conf)
	defer renderer.destroy()

	g.currentScene = s

	go frameRenderer(g, renderer, g.conf)
	go gameTick(g, g.conf)

	g.run()

	for g.running {
		handleEvents(g)
	}

	time.Sleep(1 * time.Second)
}

func (g *game) run() {
	log.Println("run the game!")
	g.running = true
}

func (g *game) stop() {
	log.Println("stop the game!")
	g.running = false
}

func gameTick(g *game, conf *Conf) {
	timePerTick := time.Second / time.Duration(conf.tps)
	log.Println("timePerTick: ", timePerTick)
	g.gameTicker = time.NewTicker(timePerTick) // this ticker never stops
	for range g.gameTicker.C {

		if g.running {
			g.currentScene.Tick()
		}
	}
}

func frameRenderer(g *game, renderer renderer, conf *Conf) {
	timePerFrame := time.Second / time.Duration(conf.fps)
	log.Println("timePerFrame: ", timePerFrame)
	g.frameTicker = time.NewTicker(timePerFrame) // this ticker never stops
	for range g.frameTicker.C {
		if g.running {
			renderer.renderFrame(g.currentScene.GetObjects())
		}
	}
}

func handleEvents(g *game) {
	if event := sdl.PollEvent(); event != nil {
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			if t.Keysym.Scancode == 41 {
				log.Println("pressed esc. quitting game...")
				g.running = false
				g.gameTicker.Stop()
				g.frameTicker.Stop()
			}
		}

		handleEvent(g, event)
	}
}

func handleEvent(g *game, event sdl.Event) {
	g.currentScene.HandleEvents(event)
}
