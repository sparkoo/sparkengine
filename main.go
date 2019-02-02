package main

import (
	"github.com/sparkoo/sparkengine/core"
	"log"
)

func main() {
	conf := core.NewConf(30, 30, 320, 240)
	g := core.NewGame(conf)
	g.Start(conf, func() {
		log.Println("Hello")
	})
}
