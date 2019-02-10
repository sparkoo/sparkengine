package scene

type Object interface {
	GetXoffset() int
	GetYoffset() int
	GetXsize() int
	GetYsize() int
	GetPixels() []Pixel
	GetPos() (float64, float64)

	MoveBy(x float64, y float64)
	MoveTo(x float64, y float64)
}

type Base struct {
	xpos  float64
	ypos  float64
	xsize int
	ysize int
}

func NewBase(xpos float64, ypos float64, xsize int, ysize int) *Base {
	return &Base{xpos: xpos, ypos: ypos, xsize: xsize, ysize: ysize}
}

func (o *Base) GetXoffset() int {
	return int(o.xpos)
}

func (o *Base) GetYoffset() int {
	return int(o.ypos)
}

func (o *Base) GetPos() (float64, float64) {
	return o.xpos, o.ypos
}

func (o *Base) GetXsize() int {
	return o.xsize
}

func (o *Base) GetYsize() int {
	return o.ysize
}

func (o *Base) MoveBy(x float64, y float64) {
	o.xpos += x
	o.ypos += y
}

func (o *Base) MoveTo(x float64, y float64) {
	o.xpos = x
	o.ypos = y
}

type Pixel struct {
	X     int
	Y     int
	Color Color
}

type Color [4]byte

func NewPixel(x int, y int, color Color) Pixel {
	return Pixel{X: x, Y: y, Color: color}
}

func Collides(o1 Object, o2 Object) bool {
	xColision := o1.GetXoffset()+o1.GetXsize() > o2.GetXoffset() &&
		o1.GetXoffset() < o2.GetXoffset()+o2.GetXsize()

	yColision := o1.GetYoffset()+o1.GetYsize() > o2.GetYoffset() &&
		o1.GetYoffset() < o2.GetYoffset()+o2.GetYsize()

	return xColision && yColision
}
