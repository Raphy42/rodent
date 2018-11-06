package compute

import (
	"log"
	"time"

	"github.com/raphy42/rodent/tokeep/cl"
	"github.com/raphy42/rodent/tokeep/clcl"
	"github.com/raphy42/rodent/core/system"
)

var (
	Instance = NewCompute()
)

type Compute struct {
	system.Module
	platform *cl.Platform
}

func (c *Compute) PreInit(interface{}) error {
	platforms, err := cl.ListPlatforms()
	if err != nil {
		return err
	}
	log.Printf("%d platforms detected\n", len(platforms))
	c.platform = platforms[0]
	return nil
}

func (c *Compute) Init() error {
	devices, err := c.platform.GetDevices(cl.DeviceTypeAll)
	if err != nil {
		return err
	}
	for _, device := range devices {
		name, err := device.Name()
		if err != nil {
			return err
		}
		version, err := device.Version()
		if err != nil {
			return err
		}
		log.Printf("Device: name=%s version=%s\n", name, version)
	}
	return nil
}

func (c Compute) PostInit() error {
	return nil
}

func (c Compute) Ticker() func(time.Time) {
	return func(delta time.Time) {

	}
}

func (c Compute) Shutdown() error {
	return nil
}

func NewCompute() *Compute {
	return &Compute{
		Module: system.NewModule("compute", system.LowestPriority),
	}
}
