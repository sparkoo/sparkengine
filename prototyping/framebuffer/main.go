package main

import (
	"fmt"
	"log"
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

var framebuffer []int8
var imageData js.Value

func main() {
	fmt.Println("Hello from webassembly!")

	doc = js.Global().Get("document")
	canvas = doc.Call("getElementById", "myCanvas")
	ctx = canvas.Call("getContext", "2d")
	imageData = ctx.Call("createImageData", WIDTH, HEIGHT)

	canvas.Set("width", WIDTH)
	canvas.Set("height", HEIGHT)

	framebuffer = make([]int8, BUFFERSIZE)

	loop(randNoise)
	fmt.Println("bye")
}

//// 0.63fps
//func randNoise() {
//	log.Println("generating random noise")
//	for i := 0; i < BUFFERSIZE; i += 4 {
//		c := i % 255
//		imageData.Get("data").SetIndex(i, c)
//		imageData.Get("data").SetIndex(i+1, c)
//		imageData.Get("data").SetIndex(i+2, c)
//		imageData.Get("data").SetIndex(i+3, 255)
//	}
//	log.Println("done")
//}

//// 2.3fps
//func randNoise() {
//	log.Println("generating random noise")
//	data := imageData.Get("data")
//	for i := 0; i < BUFFERSIZE; i += 4 {
//		c := i % 255
//		data.SetIndex(i, c)
//		data.SetIndex(i+1, c)
//		data.SetIndex(i+2, c)
//		data.SetIndex(i+3, 255)
//	}
//	log.Println("done")
//}

type imgdata struct {
	uint8
}

func randNoise() {
	log.Println("generating random noise")
	data := imageData.Get("data")
	//f := make([]interface{}, BUFFERSIZE)
	for i := 0; i < BUFFERSIZE; i += 4 {
		c := uint8(i % 255)
		data.SetIndex(i, c)
		data.SetIndex(i+1, c)
		data.SetIndex(i+2, c)
		data.SetIndex(i+3, 255)
	}
	// fmt.Println(arr)

	//imageData.Set("data", f)
	//imageData.Get("data").SetIndex(25, 255)
	//imageData.Get("data").SetIndex(26, 255)
	//imageData.Get("data").SetIndex(27, 255)
	//imageData.Get("data").SetIndex(28, 255)
	fmt.Println(imageData.Get("data").Type())
	log.Println("done")
}

func loop(foo func()) {
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
		foo()
		ctx.Call("putImageData", imageData, 0, 0)
		log.Println("end")
		js.Global().Call("requestAnimationFrame", renderFrame)
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
