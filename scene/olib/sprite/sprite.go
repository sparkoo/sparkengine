package sprite

import "github.com/sparkoo/sparkengine/scene"

type Sprite struct {
	*scene.Base

	pixels []scene.Pixel
}

func (s *Sprite) GetPixels() []scene.Pixel {
	return s.pixels
}

func NewSprite(x float64, y float64, width int, height int, getPixels func() []scene.Pixel) *Sprite {
	return &Sprite{Base: scene.NewBase(x, y, width, height), pixels: getPixels()}
}
