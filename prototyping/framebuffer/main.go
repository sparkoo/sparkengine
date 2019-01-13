package main

import (
	"fmt"
	"log"
	"math/rand"
	"syscall/js"
	"time"
)

var canvas js.Value
var doc js.Value
var ctx js.Value
var blocksize = 10

const WIDTH = 320
const HEIGHT = 200

const BUFFERSIZE = WIDTH * HEIGHT * 4

var framebuffer = [BUFFERSIZE]int8{}
var imageData js.Value

func main() {
	fmt.Println("Hello from webassembly!")

	doc = js.Global().Get("document")
	canvas = doc.Call("getElementById", "myCanvas")
	ctx = canvas.Call("getContext", "2d")
	imageData = ctx.Call("createImageData", WIDTH, HEIGHT)

	canvas.Set("width", WIDTH)
	canvas.Set("height", HEIGHT)

	fmt.Println("before")
	go randNoise()
	fmt.Println("after")

	loop()
	fmt.Println("bye")
}

func randNoise() {
	fmt.Println("generating random noise")
	time.Sleep(250 * time.Millisecond)
	for i := 0; i < BUFFERSIZE; i += 4 {
		imageData.Get("data").SetIndex(i, rand.Int()%255)
		imageData.Get("data").SetIndex(i+1, rand.Int()%255)
		imageData.Get("data").SetIndex(i+2, rand.Int()%255)
		imageData.Get("data").SetIndex(i+3, rand.Int()%255)
	}
	go randNoise()
}

func loop() {
	done := make(chan struct{}, 0)

	fmt.Println("hsi")

	fpsField := doc.Call("getElementById", "fps")

	start := time.Now()
	var renderFrame js.Callback
	renderFrame = js.NewCallback(func(args []js.Value) {
		t := time.Now()
		diff := t.Sub(start).Nanoseconds() / 1000
		fps := (float64(1) / float64(diff)) * 1000 * 1000
		fpsField.Set("innerHTML", fmt.Sprintf("%.2f", fps))
		start = t

		log.Println("start")
		ctx.Call("putImageData", imageData, 0, 0)
		log.Println("end")
		js.Global().Call("requestAnimationFrame", renderFrame)
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
