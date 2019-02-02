package core

import (
	"fmt"
	"time"
)

var running = true

func Start(conf *Conf, action func()) {
	go gameTicker(conf, action)
	go frameRenderer(conf)
	fmt.Println("doing something else")
	time.Sleep(time.Millisecond * 1500)
}

func gameTicker(conf *Conf, action func()) {
	timePerTick := time.Microsecond / time.Duration(conf.TicksPS)
	ticker := time.NewTicker(timePerTick * time.Microsecond) // this ticker never stops
	for range ticker.C {
		action()
	}
}

func frameRenderer(conf *Conf) {
	timePerFrame := time.Microsecond / time.Duration(conf.FPS)
	frameTicker := time.NewTicker(timePerFrame * time.Microsecond)	// this ticker never stops
	for range frameTicker.C {
		renderFrame()
	}
}
