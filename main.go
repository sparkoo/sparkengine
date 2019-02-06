package main

import (
	"fmt"
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
)

const SWIDTH = 320
const SHEIGHT = 240
const FPS = 30

func main() {
	game := core.NewGame(core.NewConf(FPS, FPS * 2, SWIDTH, SHEIGHT))

	b := NewBall(0, 0, 3.3, 4.5)

	s := scene.NewScene(func() {
		b.Move(1, SWIDTH, SHEIGHT)
	})
	s.AddObject(b)
	s.AddEventListener(sdl.MOUSEMOTION, func(event sdl.Event) {
		fmt.Println(event)
	})

	game.Start(s)
}
