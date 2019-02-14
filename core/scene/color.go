package scene

import "math/rand"

type Color [4]byte

const (
	COLOR_R = 0
	COLOR_G = 1
	COLOR_B = 2
	COLOR_A = 3
)

func RandomColor(alpha byte) Color {
	return Color{
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		alpha}
}
