package line

import "github.com/sparkoo/sparkengine/scene"

type Line struct {
	*scene.Base
	p1x float64
	p1y float64
	p2x float64
	p2y float64

	color scene.Color

	pixels []scene.Pixel
}

func (l *Line) GetPixels() []scene.Pixel {
	return l.pixels
}

func NewLine(x1 float64, y1 float64, x2 float64, y2 float64, color scene.Color) *Line {
	x, y, xsize, ysize := coords(x1, y1, x2, y2)

	return &Line{
		Base:   scene.NewBase(x, y, xsize, ysize),
		pixels: initPixels(x1, y1, x2, y2, color)}
}

func coords(x1 float64, y1 float64, x2 float64, y2 float64) (x float64, y float64, xsize int, ysize int) {
	// determine most left point and X size
	if x1 < x2 {
		x = x1
		xsize = int(x2 - x1)
	} else {
		x = x2
		xsize = int(x1 - x2)
	}

	// determine most top point and y size
	if y1 < y2 {
		y = y1
		ysize = int(y2 - y1)
	} else {
		y = y2
		ysize = int(y1 - y2)
	}
	return
}

// TODO: this drawing works just for horizontal-ish lines. More vertical the line is, more broken it is.
//  implement this https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
func initPixels(x1 float64, y1 float64, x2 float64, y2 float64, color scene.Color) []scene.Pixel {
	pixels := make([]scene.Pixel, int(x2-x1))
	ci := 0

	dx := x2 - x1
	dy := y2 - y1

	for x := x1; x < x2; x++ {
		y := y1 + dy*(x-x1)/dx

		pixels[ci] = scene.NewPixel(int(x), int(y), color)
		ci++
	}

	return pixels
}
