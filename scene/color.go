package scene

import "math/rand"

type Color [4]byte

func RandomColor(alpha byte) Color {
	return Color{
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		alpha}
}
