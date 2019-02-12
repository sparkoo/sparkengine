package button

import (
	"github.com/sparkoo/sparkengine/scene"
	"github.com/sparkoo/sparkengine/scene/olib/shape"
)

type Button struct {
	*shape.Rect

	action func()
}

func NewButton(x float64, y float64, xsize int, ysize int, action func()) *Button {
	return &Button{
		Rect:   shape.NewRect(x, y, xsize, ysize, scene.RandomColor(255)),
		action: action}
}

func (b *Button) GetPixels() []scene.Pixel {
	return b.Rect.GetPixels()
}

func (b *Button) Submit() {
	b.action()
}
