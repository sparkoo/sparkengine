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

type Cursor struct {
	xpos   float64
	ypos   float64
	xsize  int
	ysize  int
	pixels []Pixel
}

func NewCursor(x float64, y float64) *Cursor {
	xsize, ysize := 10, 10
	col := [4]byte{127, 127, 127, 255}
	p := make([]Pixel, 100)
	for i := 0; i < xsize; i++ {
		for j := 0; j < ysize; j++ {
			p[(i * 10) + j] = Pixel{i, j, col}
		}
	}
	return &Cursor{x, y, xsize, ysize, p}
}

func (b *Cursor) getXoffset() int {
	return int(b.xpos)
}

func (b *Cursor) getYoffset() int {
	return int(b.ypos)
}

func (b *Cursor) getXsize() int {
	return b.xsize
}

func (b *Cursor) getYsize() int {
	return b.ysize
}

func (b *Cursor) getPixels() []Pixel {
	return b.pixels
}

func (b *Cursor) String() string {
	return fmt.Sprintf("x: [%v], y: [%v] | xvel: [%v], yvel: [%v] | xsize: [%v], ysize: [%v]",
		b.xpos, b.ypos, b.xsize, b.ysize)
}
