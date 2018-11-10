package message

type Bus interface {
	Channel(string) chan Message
}

type EventBus struct {
	channels map[string]chan Message
}

func NewEventBus() *EventBus {
	return &EventBus{
		channels: make(map[string]chan Message, 16),
	}
}

func (e *EventBus) Channel(typename string) chan Message {
	channel, ok := e.channels[typename]

	if !ok {
		channel = make(chan Message)
		e.channels[typename] = channel
	}

	return channel
}
