package event

import "github.com/veandco/go-sdl2/sdl"

const (
	KEYDOWN       = sdl.KEYDOWN       // key pressed
	KEYUP         = sdl.KEYUP         // key released
	TEXTEDITING   = sdl.TEXTEDITING   // keyboard text editing (composition)
	TEXTINPUT     = sdl.TEXTINPUT     // keyboard text input
	KEYMAPCHANGED = sdl.KEYMAPCHANGED // keymap changed due to a system event such as an input language or keyboard layout change (>= SDL 2.0.4)
)

type KeypressState uint8

const (
	PRESSED  = sdl.PRESSED
	RELEASED = sdl.RELEASED
)

type KeyboardEvent struct {
	CommonEvent
	eventType EventType
	state     KeypressState
	key       Key
}

func NewKeyboardEvent(event sdl.Event) Event {
	if e, ok := event.(*sdl.KeyboardEvent); ok {
		return &KeyboardEvent{
			CommonEvent: CommonEvent{
				timestamp: e.GetTimestamp(),
			},
			eventType: EventType(e.Type),
			state:     KeypressState(e.State),
			key: Key{
				keycode: uint32(e.Keysym.Scancode),
				mod:     e.Keysym.Mod,
			},
		}
	} else {
		panic("invalid event type")
	}
}

func (e *KeyboardEvent) GetType() EventType {
	return e.eventType
}

func (e *KeyboardEvent) GetState() KeypressState {
	return e.state
}

func (e *KeyboardEvent) GetKey() *Key {
	return &e.key
}

type Key struct {
	keycode uint32
	mod     uint16
}

func (k *Key) GetKeycode() uint32 {
	return k.keycode
}

func (k *Key) GetMod() uint16 {
	return k.mod
}
