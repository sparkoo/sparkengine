package core

import (
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

type game struct {
	running bool
	conf *Conf

	currentScene *scene.Scene
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
	go gameTicker(g, g.conf)

	g.run()
	for g.running {
		time.Sleep(1 * time.Second)
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

func gameTicker(g *game, conf *Conf) {
	timePerTick := time.Second / time.Duration(conf.tps)
	log.Println("timePerTick: ", timePerTick)
	ticker := time.NewTicker(timePerTick) // this ticker never stops
	for range ticker.C {
		if event := sdl.PollEvent(); event != nil {
			handleEvent(event)
		}

		if g.running {
			g.currentScene.Tick()
		}
	}
}

func frameRenderer(g *game, renderer renderer, conf *Conf) {
	timePerFrame := time.Second / time.Duration(conf.fps)
	log.Println("timePerFrame: ", timePerFrame)
	frameTicker := time.NewTicker(timePerFrame) // this ticker never stops
	for range frameTicker.C {
		if g.running {
			renderer.renderFrame(g.currentScene.GetObjects())
		}
	}
}

func handleEvent(event scene.Event) {

}