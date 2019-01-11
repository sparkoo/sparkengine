package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	fmt.Println("Hello from webassembly!")

	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "myCanvas")
	fpsField := doc.Call("getElementById", "fps")

	ctx := canvas.Call("getContext", "2d")

	ctx.Set("fillStyle", "green")

	done := make(chan struct{}, 0)

	blocksize := 10
	speed := 1

	x, y := 0, 0
	vx, vy := speed, speed
	width, height := canvas.Get("width").Int(), canvas.Get("height").Int()


	start := time.Now()
	var renderFrame js.Callback
	renderFrame = js.NewCallback(func(args []js.Value) {
		t := time.Now()
		diff := t.Sub(start).Nanoseconds() / 1000
		fps := (float64(1) / float64(diff)) * 1000 * 1000
		fpsField.Set("innerHTML", fmt.Sprintf("%.2f", fps))
		start = t
		if x >= width - blocksize {
			vx = -speed
		}
		if x <= 0 {
			vx = speed
		}
		if y >= height - blocksize {
			vy = -speed
		}
		if y <= 0 {
			vy = speed
		}

		x += vx
		y += vy
		ctx.Call("clearRect", 0, 0, width, height)
		ctx.Call("fillRect", x, y, blocksize, blocksize)
		js.Global().Call("requestAnimationFrame", renderFrame)
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
