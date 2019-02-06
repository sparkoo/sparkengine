package scene

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	objects        []Object
	tickAction     func()
	eventListeners map[uint32][]func(sdl.Event)
}

func NewScene(tickAction func()) *Scene {
	return &Scene{
		objects:        make([]Object, 0),
		tickAction:     tickAction,
		eventListeners: make(map[uint32][]func(event sdl.Event))}
}

func (*Scene) Start() {

}

func (*Scene) Stop() {

}

func (s *Scene) AddObject(o Object) {
	s.objects = append(s.objects, o)
}

func (s *Scene) GetObjects() []Object {
	return s.objects
}

func (s *Scene) Tick() {
	s.tickAction()
}

func (s *Scene) AddEventListener(event uint32, action func(sdl.Event)) {
	s.eventListeners[event] = append(s.eventListeners[event], action)
}

func (s *Scene) HandleEvents(event sdl.Event) {
	for _, e := range s.eventListeners[event.GetType()] {
		e(event)
	}
}
