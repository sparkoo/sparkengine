package scene

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	objects        []Object
	tickAction     func()
	eventListeners []func(sdl.Event)
}

func NewScene(tickAction func()) *Scene {
	return &Scene{
		objects:        make([]Object, 0),
		tickAction:     tickAction,
		eventListeners: make([]func(event sdl.Event), 0)}
}

func (*Scene) Start() {

}

func (*Scene) Stop() {

}

func (s *Scene) AddObject(o Object) {
	s.objects = append(s.objects, o)
}

func (s *Scene) AddObjects(o ...Object) {
	s.objects = append(s.objects, o...)
}

func (s *Scene) GetObjects() []Object {
	return s.objects
}

func (s *Scene) Tick() {
	s.tickAction()
}

func (s *Scene) AddEventListener(action func(sdl.Event)) {
	s.eventListeners = append(s.eventListeners, action)
}

func (s *Scene) HandleEvents(event sdl.Event) {
	for _, e := range s.eventListeners {
		e(event)
	}
}

func NoopTick() {

}
