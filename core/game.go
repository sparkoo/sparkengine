package core

import (
	"github.com/sparkoo/sparkengine/scene"
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

func (g *game) Start(conf *Conf, s *scene.Scene) {
	renderer := &sdlRenderer{}
	renderer.init(conf)
	defer renderer.destroy()

	go frameRenderer(g, renderer, conf)
	go gameTicker(g, conf, s)

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

func gameTicker(g *game, conf *Conf, s *scene.Scene) {
	timePerTick := time.Second / time.Duration(conf.tps)
	log.Println("timePerTick: ", timePerTick)
	ticker := time.NewTicker(timePerTick) // this ticker never stops
	for range ticker.C {
		if g.running {
			s.Tick()
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
