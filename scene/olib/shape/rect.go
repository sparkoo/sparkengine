package shape

import "github.com/sparkoo/sparkengine/scene"

type Rect struct {
	*scene.Base
	color  scene.Color
	pixels []scene.Pixel
}

func (f *Rect) GetPixels() []scene.Pixel {
	return f.pixels
}

func NewRect(xpos float64, ypos float64, xsize int, ysize int, color scene.Color) *Rect {
	return &Rect{
		Base: scene.NewBase(xpos, ypos, xsize, ysize), color: color,
		pixels: initRectPixels(xsize, ysize, color)}
}

func initRectPixels(xsize int, ysize int, color scene.Color) []scene.Pixel {
	// top and bottom lines -> xsize * 2
	// left and right lines -> ysize * 2
	// 4 corners are in both top/bottom and left/right lines -> -4
	pixels := make([]scene.Pixel, (xsize*2)+((ysize*2)-4))
	pi := 0

	// top and bottom lines
	for xi := 0; xi < xsize; xi++ {
		pixels[pi] = scene.NewPixel(xi, 0, color)
		pi++

		pixels[pi] = scene.NewPixel(xi, ysize-1, color)
		pi++
	}

	// left and right lines
	for yi := 1; yi < ysize-1; yi++ {
		pixels[pi] = scene.NewPixel(0, yi, color)
		pi++

		pixels[pi] = scene.NewPixel(xsize-1, yi, color)
		pi++
	}

	return pixels
}
