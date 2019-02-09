package core

import (
	"github.com/sparkoo/sparkengine/fpscounter"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

type Game struct {
	running bool
	conf    *Conf

	currentScene *scene.Scene

	gameTicker  *time.Ticker
	frameTicker *time.Ticker
}

func NewGame(conf *Conf) *Game {
	return &Game{running: false, conf: conf}
}

func (g *Game) Start(s *scene.Scene) {
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

func (g *Game) Quit() {
	g.running = false
	g.gameTicker.Stop()
	g.frameTicker.Stop()
}

func (g *Game) SwitchScene(s *scene.Scene) {
	g.stop()
	g.currentScene = s
	g.run()
}

func (g *Game) run() {
	log.Println("run the Game!")
	g.running = true
}

func (g *Game) stop() {
	log.Println("stop the Game!")
	g.running = false
}

func gameTick(g *Game, conf *Conf) {
	timePerTick := time.Second / time.Duration(conf.tps)
	log.Println("timePerTick: ", timePerTick)
	g.gameTicker = time.NewTicker(timePerTick) // this ticker never stops

	fps := fpscounter.NewFpsCounter("ticks", 1000)
	for range g.gameTicker.C {
		if g.running {
			g.currentScene.Tick()
			fps.Tick()
		}
	}
}

func frameRenderer(g *Game, renderer renderer, conf *Conf) {
	timePerFrame := time.Second / time.Duration(conf.fps)
	log.Println("timePerFrame: ", timePerFrame)
	g.frameTicker = time.NewTicker(timePerFrame) // this ticker never stops

	fps := fpscounter.NewFpsCounter("frames", 2000)
	for range g.frameTicker.C {
		if g.running {
			renderer.renderFrame(g.currentScene.GetObjects())
			fps.Tick()
		}
	}
}

func handleEvents(g *Game) {
	if event := sdl.PollEvent(); event != nil {
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			if t.Keysym.Scancode == 41 {
				log.Println("pressed esc. quitting Game...")
				g.Quit()
			}
		}

		handleEvent(g, event)
	}
}

func handleEvent(g *Game, event sdl.Event) {
	g.currentScene.HandleEvents(event)
}
