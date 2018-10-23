package event

import (
	"github.com/raphy42/rodent/core/system"
	"time"
)

var (
	Instance = NewEvent()
)

type Event struct {
	system.Module
	state       State
	actions     []Action
	hashName    map[string]int
	reverseHash []string
	actionCache []Action
}

func NewEvent() *Event {
	event := Event{
		Module:      system.NewModule("event-bus", system.HighestPriority),
		reverseHash: make([]string, 0),
		hashName: make(map[string]int),
	}
	return &event
}

func (e *Event) PreInit(interface{}) error {
	return nil
}

func (e *Event) Init() error {

	return nil
}

func (e *Event) PostInit() error {
	return nil
}

func (e *Event) Ticker() func(time.Time) {
	return func(delta time.Time) {
		for _, action := range e.actions {
			action(&e.state)
		}
		e.actions = e.actions[:0]
	}
}

func (e *Event) Call(actions ...Action) {
	e.actions = append(e.actions, actions...)
}

func (e *Event) CallHash(hashes ...int) {
	for _, hash := range hashes {
		e.actions = append(e.actions, e.actionCache[hash])
	}
}

func (e *Event) Register(name string, action Action) int {
	if hash, ok := e.hashName[name]; ok {
		return hash
	}

	hash := len(e.reverseHash)
	e.actionCache = append(e.actionCache, action)
	e.reverseHash = append(e.reverseHash, name)
	e.hashName[name] = hash
	return hash
}

func (e *Event) RegisterMultiple(actionMap map[string]Action) {
	for name, action := range actionMap {
		e.Register(name, action)
	}
}

func (e *Event) Hash(name string) (int, bool) {
	v, ok := e.hashName[name]
	return v, ok
}
