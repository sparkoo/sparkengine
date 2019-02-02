package scene

type Ball struct {
	xpos float64
	ypos float64
	xvel float64
	yvel float64
}

func NewBall() *Ball {
	return &Ball{0, 0, 1, 2}
}

const size = 10

var ballPixels []Pixel

func init() {
	ballPixels := make([]*Pixel, size*size)
	c := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			ballPixels[c] = NewPixel(i, j, [4]byte{127, 127, 127, 127})
		}
	}
}

func (b *Ball) getXoffset() int {
	return int(b.xpos)
}

func (b *Ball) getYoffset() int {
	return int(b.ypos)
}

func (b *Ball) getXsize() int {
	return size
}

func (b *Ball) getYsize() int {
	return size
}

func (b *Ball) getPixels() []Pixel {
	return ballPixels
}
