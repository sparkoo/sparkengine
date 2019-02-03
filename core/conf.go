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