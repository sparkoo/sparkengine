package event

type EventType uint

type Event interface {
	GetTimestamp() uint32
}

type CommonEvent struct {
	timestamp uint32
}

func (e *CommonEvent) GetTimestamp() uint32 {
	return e.timestamp
}
