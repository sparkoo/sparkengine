package core

import (
	"fmt"
	"time"
)

var running = true

func Start(conf *Conf, action func()) {
	renderer := sdlRenderer{}
	renderer.init(conf)

	go frameRenderer(renderer, conf)
	go gameTicker(conf, action)
	fmt.Println("doing something else")
	time.Sleep(time.Millisecond * 500)
}

func gameTicker(conf *Conf, action func()) {
	timePerTick := time.Microsecond / time.Duration(conf.TicksPS)
	ticker := time.NewTicker(timePerTick * time.Microsecond) // this ticker never stops
	for range ticker.C {
		action()
	}
}

func frameRenderer(renderer renderer, conf *Conf) {
	timePerFrame := time.Microsecond / time.Duration(conf.FPS)
	frameTicker := time.NewTicker(timePerFrame * time.Microsecond)	// this ticker never stops
	for range frameTicker.C {
		renderer.renderFrame()
	}
}
