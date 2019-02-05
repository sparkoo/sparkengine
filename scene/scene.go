package scene

type Scene struct {
	objects []Object
	tickAction func()
}

func NewScene(tickAction func()) *Scene {
	return &Scene{objects: make([]Object, 0), tickAction: tickAction}
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

func (s *Scene) AddEventListener(event, action func()) {

}
