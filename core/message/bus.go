package message

import "go.uber.org/zap"

const (
	ChannelSize = 1
)

type Bus interface {
	Publish(Type) chan Message
	Subscribe(Type) chan Message
}

type EventBus struct {
	channels map[Type]chan Message
	clients  map[Type][]chan Message
}

func NewEventBus() *EventBus {
	return &EventBus{
		channels: make(map[Type]chan Message),
		clients:  make(map[Type][]chan Message),
	}
}

func (e *EventBus) bootstrap(xtype Type) {
	channel := make(chan Message)
	e.channels[xtype] = channel
	e.clients[xtype] = make([]chan Message, 0)

	zTypename := zap.String("type", xtype.String())

	log.Debug("new subscriber pool", zTypename)

	go func() {
		for {
			ev := <-e.channels[xtype]
			for _, client := range e.clients[xtype] {
				// log.Debug("event", zTypename, zap.Int("slot", slot))
				client <- ev
			}
		}
	}()
}

func (e *EventBus) Subscribe(xtype Type) chan Message {
	_, ok := e.channels[xtype]

	log.Info("new subscriber", zap.String("type", xtype.String()))

	if !ok {
		e.bootstrap(xtype)
	}
	client := make(chan Message, ChannelSize)
	e.clients[xtype] = append(e.clients[xtype], client)

	return client
}

func (e *EventBus) Publish(xtype Type) chan Message {
	_, ok := e.channels[xtype]

	log.Info("new publisher", zap.String("type", xtype.String()))

	if !ok {
		e.bootstrap(xtype)
	}
	return e.channels[xtype]
}
