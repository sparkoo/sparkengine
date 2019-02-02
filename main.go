package main

import (
	"fmt"
	"github.com/sparkoo/sparkengine/core"
)

func main() {
	conf := &core.Conf{TicksPS: 30, FPS: 60}
	core.Start(conf, func() {
		fmt.Println("Hello")
	})
}
