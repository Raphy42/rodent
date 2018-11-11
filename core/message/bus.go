package message

import "go.uber.org/zap"

const (
	ChannelSize = 1
)

type Bus interface {
	Publish(string) chan Message
	Subscribe(string) chan Message
}

type EventBus struct {
	channels map[string]chan Message
	clients  map[string][]chan Message
}

func NewEventBus() *EventBus {
	return &EventBus{
		channels: make(map[string]chan Message),
		clients:  make(map[string][]chan Message),
	}
}

func (e *EventBus) bootstrap(typename string) {
	channel := make(chan Message)
	e.channels[typename] = channel
	e.clients[typename] = make([]chan Message, 0)

	zTypename := zap.String("type", typename)

	log.Debug("new subscriber pool", zTypename)

	go func() {
		for {
			ev := <-e.channels[typename]
			for _, client := range e.clients[typename] {
				// go func() {
				// log.Debug("event", zTypename, zap.Int("slot", slot))
				client <- ev
				// }()
			}
		}
	}()
}

func (e *EventBus) Subscribe(typename string) chan Message {
	_, ok := e.channels[typename]

	log.Info("new subscriber", zap.String("type", typename))

	if !ok {
		e.bootstrap(typename)
	}
	client := make(chan Message, ChannelSize)
	e.clients[typename] = append(e.clients[typename], client)

	return client
}

func (e *EventBus) Publish(typename string) chan Message {
	_, ok := e.channels[typename]

	log.Info("new publisher", zap.String("type", typename))

	if !ok {
		e.bootstrap(typename)
	}
	return e.channels[typename]
}
