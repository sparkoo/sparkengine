package shape

import "github.com/sparkoo/sparkengine/core/scene"

type Ball struct {
	*scene.Base
	color  scene.Color
	pixels []scene.Pixel
}

func (b *Ball) GetPixels() []scene.Pixel {
	return b.pixels
}

func NewBall(x float64, y float64, diameter int, color scene.Color) *Ball {
	return &Ball{
		Base:   scene.NewBase(x, y, diameter, diameter),
		color:  color,
		pixels: initBallPixels(diameter/2, color),
	}
}

func initBallPixels(radius int, color scene.Color) []scene.Pixel {
	pixels := make([]scene.Pixel, 0)

	x := radius - 1
	y := 0
	dx := 1
	dy := 1
	err := dx - (radius << 1)

	x0, y0 := radius, radius

	for x >= y {
		pixels = append(pixels, scene.NewPixel(x0+x, y0+y, color))
		pixels = append(pixels, scene.NewPixel(x0+y, y0+x, color))
		pixels = append(pixels, scene.NewPixel(x0-y, y0+x, color))
		pixels = append(pixels, scene.NewPixel(x0-x, y0+y, color))
		pixels = append(pixels, scene.NewPixel(x0-x, y0-y, color))
		pixels = append(pixels, scene.NewPixel(x0-y, y0-x, color))
		pixels = append(pixels, scene.NewPixel(x0+y, y0-x, color))
		pixels = append(pixels, scene.NewPixel(x0+x, y0-y, color))

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}

		if err > 0 {
			x--
			dx += 2
			err += dx - (radius << 1)
		}
	}

	return pixels
}
