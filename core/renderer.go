package core

import (
	"log"
)

type renderer interface {
	init(conf *Conf)
	renderFrame()
}

type sdlRenderer struct {
}

func (sdlRenderer) renderFrame() {
	log.Println("frame rendered")
}

func (sdlRenderer) init(conf *Conf) {
	log.Println("initializing SDL renderer ...")
	log.Println("done")
}
