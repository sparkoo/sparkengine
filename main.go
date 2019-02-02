package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
)

func main() {
	conf := core.NewConf(30, 30, 320, 240)
	g := core.NewGame(conf)
	b := scene.NewBall()
	g.AddObject(b)
	g.Start(conf, func() {
	})
}
