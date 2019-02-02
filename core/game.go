package core

import (
	"github.com/sparkoo/sparkengine/scene"
	"log"
	"time"
)

type game struct {
	running bool
	conf *conf

	objects []scene.Object
}

func NewGame(conf *conf) *game {
	return &game{running: false, conf: conf, objects: make([]scene.Object, 0)}
}

func (g *game) AddObject(o scene.Object) {
	g.objects = append(g.objects, o)
}

func (g *game) Start(conf *conf, action func()) {
	renderer := &sdlRenderer{}
	renderer.init(conf)
	defer renderer.destroy()

	go frameRenderer(g, renderer, conf)
	go gameTicker(g, conf, action)

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

func gameTicker(g *game, conf *conf, action func()) {
	timePerTick := time.Second / time.Duration(conf.tps)
	log.Println("timePerTick: ", timePerTick)
	ticker := time.NewTicker(timePerTick) // this ticker never stops
	for range ticker.C {
		if g.running {
			action()
		}
	}
}

func frameRenderer(g *game, renderer renderer, conf *conf) {
	timePerFrame := time.Second / time.Duration(conf.fps)
	log.Println("timePerFrame: ", timePerFrame)
	frameTicker := time.NewTicker(timePerFrame) // this ticker never stops
	for range frameTicker.C {
		if g.running {
			renderer.renderFrame()
		}
	}
}
