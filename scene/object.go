package scene

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

func NewPixel(x int, y int, color [4]byte) *Pixel {
	return &Pixel{x: x, y: y, color: color}
}
