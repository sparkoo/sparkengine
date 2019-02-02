package core

import (
	"log"
	"time"
)

type game struct {
	running bool
	conf *conf
}

func NewGame(conf *conf) *game {
	return &game{running: false, conf: conf}
}

func (g *game) Start(conf *conf, action func()) {
	renderer := &sdlRenderer{}
	renderer.init(conf)
	defer renderer.destroy()

	go frameRenderer(g, renderer, conf)
	go gameTicker(g, conf, action)

	time.Sleep(time.Second * 1)
	g.run()
	time.Sleep(time.Second * 1)
	g.pause()
	time.Sleep(time.Second * 1)
}

func (g *game) run() {
	g.running = true
}

func (g *game) pause() {
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
