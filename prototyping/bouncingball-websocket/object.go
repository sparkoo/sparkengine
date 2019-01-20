package main

import "fmt"

type Object interface {
	getXoffset() int
	getYoffset() int
	getXsize() int
	getYsize() int
	getPixels() []Pixel
}

type Pixel struct {
	x     int
	y     int
	color [4]byte
}

type Ball struct {
	xpos   float64
	ypos   float64
	xsize  int
	ysize  int
	xvel   float64
	yvel   float64
	pixels []Pixel
}

func NewBall(x float64, y float64, xsize int, ysize int, xvel float64, yvel float64) *Ball {
	col := [4]byte{127, 127, 127, 255}
	p := make([]Pixel, 100)
	for i := 0; i < xsize; i++ {
		for j := 0; j < ysize; j++ {
			p = append(p, Pixel{i, j, col})
		}
	}
	return &Ball{x, y, xsize, ysize, xvel, yvel, p}
}

func (b *Ball) getXoffset() int {
	return int(b.xpos)
}

func (b *Ball) getYoffset() int {
	return int(b.ypos)
}

func (b *Ball) getXsize() int {
	return b.xsize
}

func (b *Ball) getYsize() int {
	return b.ysize
}

func (b *Ball) getPixels() []Pixel {
	return b.pixels
}

func (b *Ball) String() string {
	return fmt.Sprintf("x: [%v], y: [%v] | xvel: [%v], yvel: [%v] | xsize: [%v], ysize: [%v]",
		b.xpos, b.ypos, b.xvel, b.yvel, b.xsize, b.ysize)
}
