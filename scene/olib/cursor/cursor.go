package cursor

import (
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

var (
	color        = scene.Color{100, 100, 100, 255}
	cursorPixels []scene.Pixel
)

const (
	size = 10
)

func init() {
	cursorPixels = make([]scene.Pixel, (size*3)-2)
	pi := 0
	cursorPixels[pi] = scene.NewPixel(0, 0, color)
	pi++
	for i := 1; i < size; i++ {
		cursorPixels[pi] = scene.NewPixel(i, i, color)
		pi++
		cursorPixels[pi] = scene.NewPixel(0, i, color)
		pi++
		cursorPixels[pi] = scene.NewPixel(i, 0, color)
		pi++
	}
}

type Cursor struct {
	*scene.Base
}

func NewCursor(x float64, y float64) *Cursor {
	return &Cursor{scene.NewBase(x, y, 5, 5)}
}

func (c *Cursor) GetPixels() []scene.Pixel {
	return cursorPixels
}

func (c *Cursor) Listener(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.MouseMotionEvent:
		c.MoveTo(float64(e.X), float64(e.Y))
	case *sdl.MouseButtonEvent:
		log.Println("button pressed at[", e.X, e.Y, "]")
	}
}
