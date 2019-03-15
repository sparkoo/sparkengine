package event

const (
	KEYDOWN       EventType = iota // key pressed
	KEYUP                          // key released
	TEXTEDITING                    // keyboard text editing (composition)
	TEXTINPUT                      // keyboard text input
	KEYMAPCHANGED                  // keymap changed due to a system event such as an input language or keyboard layout change (>= SDL 2.0.4)
)

type KeypressState uint

const (
	PRESSED KeypressState = iota
	RELEASED
)

type KeyboardEvent struct {
	CommonEvent
	eventType EventType
	state     KeypressState
	key       Key
}

func (e *KeyboardEvent) GetType() EventType {
	return e.eventType
}

func (e *KeyboardEvent) GetState() KeypressState {
	return e.state
}

func (e *KeyboardEvent) GetKey() Key {
	return e.key
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
