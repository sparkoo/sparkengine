package core

type Conf struct {
	tps          int
	fps          int
	screenWidth  int32
	screenHeight int32
}

func NewConf(tps int, fps int, screenWidth int32, screenHeight int32) *Conf {
	return &Conf{tps: tps, fps: fps, screenWidth: screenWidth, screenHeight: screenHeight}
}

func NewConf30Ticks(width int32, height int32) *Conf {
	return NewConf(30, 60, width, height)
}

func NewConf30T320W() *Conf {
	return NewConf(30, 60, 320, 240)
}

func NewConf30T640W() *Conf {
	return NewConf(30, 60, 640, 480)
}
