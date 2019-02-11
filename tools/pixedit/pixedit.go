package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/sparkoo/sparkengine/scene/olib/cursor"
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
	canvasControlLine := line.NewLine(float64(canvasSize + 10), 0, float64(canvasSize + 10), conf.SHEIGHT, scene.Color{127, 127, 127, 255})

	c := cursor.NewCursor(0, 0)
	s.AddEventListener(c.Listener)

	s.AddObjects(canvasFrame, canvasControlLine, c)

	return s
}
