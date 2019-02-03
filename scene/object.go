package scene

type Object interface {
	GetXoffset() int
	GetYoffset() int
	GetXsize() int
	GetYsize() int
	GetPixels() []Pixel
}

type Pixel struct {
	X     int
	Y     int
	Color [4]byte
}

func NewPixel(x int, y int, color [4]byte) Pixel {
	return Pixel{X: x, Y: y, Color: color}
}
