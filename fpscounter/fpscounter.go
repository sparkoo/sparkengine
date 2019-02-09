package fpscounter

import (
	"log"
	"time"
)

type FpsCounter struct {
	name       string
	ticks      int
	ticksLimit int
	startTime  time.Time
}

func NewFpsCounter(name string, eachNFrames int) *FpsCounter {
	return &FpsCounter{name: name, ticks: 0, ticksLimit: eachNFrames, startTime: time.Now()}
}

func (c *FpsCounter) Tick() {
	c.ticks++
	if c.ticks == c.ticksLimit {
		tDiff := float64(time.Now().Sub(c.startTime).Nanoseconds()) / 1000 / 1000
		fps := float64(c.ticks) / (tDiff / 1000)
		log.Printf("[%s] %.2f tps", c.name, fps)
		c.startTime = time.Now()
		c.ticks = 0
	}
}
