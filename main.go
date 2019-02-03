package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
)

const SWIDTH = 320
const SHEIGHT = 240
const FPS = 30

func main() {
	conf := core.NewConf(FPS, FPS * 2, SWIDTH, SHEIGHT)
	g := core.NewGame(conf)
	b := scene.NewBall(0, 0, 1.2, .8)
	s := scene.NewScene(func() {
		b.Move(1, SWIDTH, SHEIGHT)
	})
	s.AddObject(b)
	g.Start(conf, s)
}
