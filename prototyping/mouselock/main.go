package main

import (
	"fmt"
	"syscall/js"
	"time"
)

var canvas js.Value
var doc js.Value
var updateMouseCallback js.Callback
var ctx js.Value
var width, height int
var blocksize = 10

func main() {

	fmt.Println("Hello from webassembly!")

	doc = js.Global().Get("document")
	canvas = doc.Call("getElementById", "myCanvas")

	ctx = canvas.Call("getContext", "2d")

	fmt.Println("hello www")

	canvas.Call("addEventListener", "mousedown", js.NewCallback(lockmouse))
	doc.Call("addEventListener", "pointerlockchange", js.NewCallback(lockChange))
	updateMouseCallback = js.NewCallback(updateCanvasPosition)
	width, height = canvas.Get("width").Int(), canvas.Get("height").Int()

	loop()
	fmt.Println("bye")
}

func loop() {
	done := make(chan struct{}, 0)

	fmt.Println("hsi")

	fpsField := doc.Call("getElementById", "fps")

	width, height := canvas.Get("width").Int(), canvas.Get("height").Int()

	start := time.Now()
	var renderFrame js.Callback
	renderFrame = js.NewCallback(func(args []js.Value) {
		t := time.Now()
		diff := t.Sub(start).Nanoseconds() / 1000
		fps := (float64(1) / float64(diff)) * 1000 * 1000
		fpsField.Set("innerHTML", fmt.Sprintf("%.2f", fps))
		start = t

		ctx.Call("clearRect", 0, 0, width, height)
		ctx.Call("fillRect", x, y, blocksize, blocksize)
		js.Global().Call("requestAnimationFrame", renderFrame)
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}

func lockmouse(args []js.Value) {
	fmt.Println("trying lock mouse ...")
	canvas.Call("requestPointerLock")
}

func lockChange(args []js.Value) {
	if doc.Get("pointerLockElement") == canvas {
		fmt.Println("locked. listening for mousemove")
		canvas.Call("addEventListener", "mousemove", updateMouseCallback)
	} else {
		fmt.Println("unlocked. stop listening")
		canvas.Call("removeEventListener", "mousemove", updateMouseCallback)
	}
}

var x = 0
var y = 0

func updateCanvasPosition(args []js.Value) {
	mx := args[0].Get("movementX").Int()
	if x+mx >= 0 {
		x += mx
	} else {
		x = 0
	}
	if x+mx < width-blocksize {
		x += mx
	} else {
		x = width - blocksize
	}

	my := args[0].Get("movementY").Int()
	if y+my >= 0 {
		y += my
	} else {
		y = 0
	}
	if y+my < height-blocksize {
		y += my
	} else {
		y = height - blocksize
	}
}
