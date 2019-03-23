package scene

import (
	"github.com/sparkoo/sparkengine/core/event"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Scene struct {
	objects        []Object
	tickAction     func(gameTickCounter int64, sceneTickCounter int64)
	eventListeners []func(event.Event)

	sceneTickCounter int64
}

func NewScene(tickAction func(gameTickCounter int64, sceneTickCounter int64)) *Scene {
	return &Scene{
		objects:          make([]Object, 0),
		tickAction:       tickAction,
		eventListeners:   make([]func(event.Event), 0),
		sceneTickCounter: 0}
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

func (s *Scene) Tick(gameTickCounter int64) {
	s.tickAction(gameTickCounter, s.sceneTickCounter)
	s.sceneTickCounter++
}

func (s *Scene) AddEventListener(action func(event.Event)) {
	s.eventListeners = append(s.eventListeners, action)
}

func (s *Scene) HandleEvents(event sdl.Event) {
	for _, e := range s.eventListeners {
		e(createEvent(event))
	}
}

func createEvent(e sdl.Event) event.Event {
	switch t := e.(type) {
	case *sdl.KeyboardEvent:
		return event.NewKeyboardEvent(t)
	default:
		log.Println("Unknown event", e)
		return nil
	}
}

func (s *Scene) ClearTickCounter() {
	s.sceneTickCounter = 0
}

func NoopTick(_ int64, _ int64) {
}
