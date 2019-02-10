package line

import (
	"github.com/sparkoo/sparkengine/scene"
	"log"
	"math"
)

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
		pixels: initPixels(0, 0, x2-x1, y2-y1, color)}
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
func initPixels(x0 float64, y0 float64, x1 float64, y1 float64, color scene.Color) []scene.Pixel {
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

	log.Println(pixels)

	return pixels
}

func plotLineLow(x0 float64, y0 float64, x1 float64, y1 float64, color scene.Color) []scene.Pixel {
	var pixels []scene.Pixel

	dx := x1 - x0
	dy := y1 - y0
	yi := 1.0
	if dy < 0 {
		yi = -1
		dy = -dy
	}
	D := 2*dy - dx
	y := y0

	for x := x0; x < x1; x++ {
		pixels = append(pixels, scene.NewPixel(int(x), int(y), color))
		log.Println(x, y)
		if D > 0 {
			y = y + yi
			D = D - 2*dx
		}
		D = D + 2*dy
	}

	return pixels
}

func plotLineHigh(x0 float64, y0 float64, x1 float64, y1 float64, color scene.Color) []scene.Pixel {
	var pixels []scene.Pixel

	dx := x1 - x0
	dy := y1 - y0
	xi := 1.0
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	D := 2*dx - dy
	x := x0

	for y := y0; y < y1; y++ {
		pixels = append(pixels, scene.NewPixel(int(x), int(y), color))
		log.Println(x, y)
		if D > 0 {
			x = x + xi
			D = D - 2*dy
		}
		D = D + 2*dx
	}

	return pixels
}
