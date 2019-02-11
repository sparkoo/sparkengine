package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
	canvas2 "github.com/sparkoo/sparkengine/scene/olib/canvas"
	cursor2 "github.com/sparkoo/sparkengine/scene/olib/cursor"
	"github.com/sparkoo/sparkengine/scene/olib/frame"
	"github.com/sparkoo/sparkengine/scene/olib/line"
	"github.com/sparkoo/sparkengine/tools/pixedit/conf"
)

var (
	g *core.Game
	c *core.Conf
	s *scene.Scene
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

	canvasSize := conf.SHEIGHT - 10

	canvasFrame := frame.NewFrame(5, 5, canvasSize, canvasSize, scene.Color{100, 100, 100, 255})
	canvas := canvas2.NewCanvas(10, 10, canvasSize-10, canvasSize-10, scene.Color{200, 200, 200, 255})
	canvasControlLine := line.NewLine(float64(canvasSize + 10), 0, float64(canvasSize + 10), conf.SHEIGHT, scene.Color{127, 127, 127, 255})

	cursor := cursor2.NewCursor(0, 0)
	s.AddEventListener(cursor.Listener)

	s.AddObjects(canvasFrame, canvasControlLine, canvas, cursor)

	return s
}
