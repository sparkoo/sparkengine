package core

type conf struct {
	tps          int
	fps          int
	screenWidth  int32
	screenHeight int32
}

func NewConf(tps int, fps int, screenWidth int32, screenHeight int32) *conf {
	return &conf{tps: tps, fps: fps, screenWidth: screenWidth, screenHeight: screenHeight}
}