package shape

import (
	"github.com/sparkoo/sparkengine/core/scene"
	"math"
)

// [x1;y1] determines position of the line
// pixels can go to negative positions, thus lines are not positioned by top-left corner, but by p1
type Line struct {
	*scene.Base
	x1 float64
	y1 float64
	x2 float64
	y2 float64

	color scene.Color

	pixels []scene.Pixel
}

func (l *Line) GetPixels() []scene.Pixel {
	return l.pixels
}

func NewLine(x1 float64, y1 float64, x2 float64, y2 float64, color scene.Color) *Line {
	xsize := int(math.Abs(x1 - x2))
	ysize := int(math.Abs(y1 - y2))

	return &Line{
		Base:   scene.NewBase(x1, y1, xsize, ysize),
		x1:     x1,
		y1:     y1,
		x2:     x2,
		y2:     y2,
		color:  color,
		pixels: initLinePixels(0, 0, x2-x1, y2-y1, color)}
}

func initLinePixels(x0 float64, y0 float64, x1 float64, y1 float64, color scene.Color) []scene.Pixel {
	pixels := make([]scene.Pixel, 0)

	if math.Abs(y1-y0) < math.Abs(x1-x0) {
		if x0 > x1 {
			pixels = append(pixels, plotLineLow(x1, y1, x0, y0, color)...)
		} else {
			pixels = append(pixels, plotLineLow(x0, y0, x1, y1, color)...)
		}
	} else {
		if y0 > y1 {
			pixels = append(pixels, plotLineHigh(x1, y1, x0, y0, color)...)
		} else {
			pixels = append(pixels, plotLineHigh(x0, y0, x1, y1, color)...)
		}
	}

	return pixels
}

func plotLineLow(x1 float64, y1 float64, x2 float64, y2 float64, color scene.Color) []scene.Pixel {
	var pixels []scene.Pixel

	dx := x2 - x1
	dy := y2 - y1
	yi := 1.0
	if dy < 0 {
		yi = -1
		dy = -dy
	}
	D := 2*dy - dx
	y := y1

	for x := x1; x < x2; x++ {
		pixels = append(pixels, scene.NewPixel(int(x), int(y), color))
		if D > 0 {
			y = y + yi
			D = D - 2*dx
		}
		D = D + 2*dy
	}

	return pixels
}

func plotLineHigh(x1 float64, y1 float64, x2 float64, y2 float64, color scene.Color) []scene.Pixel {
	var pixels []scene.Pixel

	dx := x2 - x1
	dy := y2 - y1
	xi := 1.0
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	D := 2*dx - dy
	x := x1

	for y := y1; y < y2; y++ {
		pixels = append(pixels, scene.NewPixel(int(x), int(y), color))
		if D > 0 {
			x = x + xi
			D = D - 2*dy
		}
		D = D + 2*dx
	}

	return pixels
}
