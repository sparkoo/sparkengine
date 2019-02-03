package main

import "github.com/sparkoo/sparkengine/scene"

type Ball struct {
	xpos float64
	ypos float64
	xvel float64
	yvel float64
}

func NewBall(xpos float64, ypos float64, xvel float64, yvel float64) *Ball {
	return &Ball{xpos, ypos, xvel, yvel}
}

const size = 10

var ballPixels []scene.Pixel

func init() {
	ballPixels = make([]scene.Pixel, size*size)
	c := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			ballPixels[c] = scene.NewPixel(i, j, [4]byte{127, 127, 127, 127})
			c++
		}
	}
}

func (b *Ball) GetXoffset() int {
	return int(b.xpos)
}

func (b *Ball) GetYoffset() int {
	return int(b.ypos)
}

func (b *Ball) GetXsize() int {
	return size
}

func (b *Ball) GetYsize() int {
	return size
}

func (b *Ball) GetPixels() []scene.Pixel {
	return ballPixels
}

func (b *Ball) Move(ticks int, sWidth int, sHeight int) {
	for ti := 0; ti < ticks; ti++ {
		xPot := int(b.xpos + b.xvel)
		if xPot < 0 || xPot+b.GetXsize() >= sWidth {
			b.xvel *= -1
		}

		yPot := int(b.ypos + b.yvel)
		if yPot < 0 || yPot+b.GetYsize() >= sHeight {
			b.yvel *= -1
		}

		b.xpos += b.xvel
		b.ypos += b.yvel
	}
}