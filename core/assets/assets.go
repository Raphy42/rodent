package assets

import (
	"github.com/raphy42/rodent/core/system"
	"log"
	"time"
)

var (
	Instance = NewAssets()
)

type Builder func(name string, filenames []string, metadata map[string]string) error

type Config struct {
	Root      string `json:"root"`
	Resources []struct {
		Name     string            `json:"name"`
		Type     string            `json:"type"`
		Files    []string          `json:"files"`
		Metadata map[string]string `json:"metadata"`
	} `json:"resources"`
}

type Assets struct {
	system.Module
	config   Config
	builders map[string]Builder
}

func NewAssets() *Assets {
	return &Assets{
		// setting the system priority allows all the other systems to register their builder during `system.Init()`
		Module:   system.NewModule("assets", system.LowestPriority),
		builders: make(map[string]Builder),
	}
}

func (a *Assets) PreInit(v interface{}) error {
	a.config = v.(Config)
	return nil
}

func (a Assets) Init() error {
	return nil
}

func (a Assets) PostInit() error {
	log.Printf("found %d resources", len(a.config.Resources))

	for _, resource := range a.config.Resources {
		builder, ok := a.builders[resource.Type]
		if !ok {
			log.Printf("unknown builder type: '%s'\n", resource.Type)
		}
		if resource.Metadata == nil {
			resource.Metadata = make(map[string]string)
		}
		resource.Metadata["root"] = a.config.Root
		err := builder(resource.Name, resource.Files, resource.Metadata)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Assets) Ticker() func(time.Time) {
	return func(delta time.Time) {

	}
}

func (a *Assets) RegisterBuilder(xtype string, builder Builder) {
	log.Printf("new builder registered: '%s'", xtype)
	a.builders[xtype] = builder
}
