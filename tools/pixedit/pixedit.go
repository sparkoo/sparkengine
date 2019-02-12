package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/sparkoo/sparkengine/scene/olib/button"
	canvas2 "github.com/sparkoo/sparkengine/scene/olib/canvas"
	cursor2 "github.com/sparkoo/sparkengine/scene/olib/cursor"
	"github.com/sparkoo/sparkengine/scene/olib/shape"
	"github.com/sparkoo/sparkengine/tools/pixedit/conf"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

var (
	g      *core.Game
	c      *core.Conf
	s      *scene.Scene
	canvas *canvas2.Canvas
)

func init() {
	c = core.NewConf(conf.FPS, conf.FPS*2, conf.SWIDTH, conf.SHEIGHT)
	g = core.NewGame(c)
	s = initScene()
}

func main() {
	g.Start(s)
}

func initScene() *scene.Scene {
	s := scene.NewScene(scene.NoopTick)

	canvasFrame := shape.NewRect(5, 5, conf.CANVAS_SIZE, conf.CANVAS_SIZE, scene.Color{100, 100, 100, 255})

	canvasControlLine := shape.NewLine(float64(conf.CANVAS_SIZE+10), 0, float64(conf.CANVAS_SIZE+10), conf.SHEIGHT, scene.Color{127, 127, 127, 255})

	canvas = canvas2.NewCanvas(10, 10, conf.CANVAS_SIZE-10, conf.CANVAS_SIZE-10, scene.Color{200, 200, 200, 255})

	cursor := cursor2.NewCursor(0, 0)
	s.AddEventListener(cursor.Listener)
	s.AddEventListener(drawListener)

	testButton := button.NewButton(conf.SWIDTH * .8, conf.SHEIGHT * .9, 50, 20, func() {
		log.Println("test button pressed")
	})

	s.AddObjects(canvasFrame, canvasControlLine, canvas, testButton, cursor)

	return s
}

var drawing = false

func drawListener(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.MouseMotionEvent:
		if drawing {
			canvas.Draw(int(e.X - 10), int(e.Y - 10), scene.Color{0, 0, 0, 255})
		}
	case *sdl.MouseButtonEvent:
		drawing = e.State == 1
	}
}
