package canvas

import "github.com/sparkoo/sparkengine/scene"

type Canvas struct {
	*scene.Base

	backgroundColor scene.Color
	canvasPixels    []scene.Pixel
}

func NewCanvas(x float64, y float64, xsize int, ysize int, backgroundColor scene.Color) *Canvas {
	return &Canvas{
		Base:            scene.NewBase(x, y, xsize, ysize),
		backgroundColor: backgroundColor,
		canvasPixels:    initPixels(backgroundColor, xsize, ysize),
	}
}

func initPixels(color scene.Color, width int, height int) []scene.Pixel {
	pixels := make([]scene.Pixel, width*height)
	pi := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pixels[pi] = scene.NewPixel(x, y, color)
			pi++
		}
	}
	return pixels
}

func (c *Canvas) GetPixels() []scene.Pixel {
	return c.canvasPixels
}

func (c *Canvas) Draw(x int, y int, color scene.Color) {
	c.canvasPixels = append(c.canvasPixels, scene.NewPixel(x, y, color))
}
