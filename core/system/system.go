package system

import (
	"math"
	"time"
)

const (
	LowestPriority = math.MinInt32
	HighestPriority = math.MaxInt32
)

type IModule interface {
	Priority() int
	Name() string
}

type ITickable interface {
	Ticker() func(time.Time)
}

type ISystem interface {
	IModule

	// PreInit should provide the context for the underlying system
	PreInit(interface{}) error

	// Init bootstraps the system
	// After this step the system must be stable and ready
	Init() error

	// PostInit is the end of the bootstrap loop
	// The system can now start jobs or communicate to other systems
	PostInit() error

	// Ticker returns a ticker
	Ticker() func(time.Time)
}

type Module struct {
	name     string
	priority int
}

func NewModule(name string, priority int) Module {
	return Module{
		name: name,
		priority: priority,
	}
}

func (m Module) Priority() int {
	return m.priority
}

func (m Module) Name() string {
	return m.name
}
