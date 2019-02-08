package scene

type Object interface {
	GetXoffset() int
	GetYoffset() int
	GetXsize() int
	GetYsize() int
	GetPixels() []Pixel

	MoveBy(x float64, y float64)
	MoveTo(x float64, y float64)
}

type Pixel struct {
	X     int
	Y     int
	Color [4]byte
}

func NewPixel(x int, y int, color [4]byte) Pixel {
	return Pixel{X: x, Y: y, Color: color}
}

func Collides(o1 Object, o2 Object) bool {
	xColision := o1.GetXoffset() + o1.GetXsize() > o2.GetXoffset() &&
		o1.GetXoffset() < o2.GetXoffset() + o2.GetXsize()

	yColision := o1.GetYoffset() + o1.GetYsize() > o2.GetYoffset() &&
		o1.GetYoffset() < o2.GetYoffset() + o2.GetYsize()

	return xColision && yColision
}
