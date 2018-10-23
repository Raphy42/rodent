package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	_ "runtime/pprof"
	"time"

	"github.com/raphy42/rodent/core/application"
	"github.com/raphy42/rodent/core/assets"
	"github.com/raphy42/rodent/core/event"
	"github.com/raphy42/rodent/core/graphic"
	"github.com/raphy42/rodent/core/system"
)

type Config struct {
	Application application.Config `json:"application"`
	Assets      assets.Config      `json:"assets"`
}

type Core struct {
	systems system.Systems
	config  Config
}

func New(config Config) *Core {
	core := new(Core)
	core.config = config
	core.systems = append(
		core.systems,
		application.Instance,
		event.Instance,
		graphic.Instance,
		assets.Instance)

	return core
}

func NewFrom(filename string) *Core {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	err = json.Unmarshal(buffer, &config)

	return New(config)
}

func (c *Core) Start() error {
	conf := make(map[string]interface{})
	conf["application"] = c.config.Application
	conf["assets"] = c.config.Assets

	return c.systems.StartAll(conf)
}

func (c *Core) Profile(addr string) {
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (c *Core) Wait() chan struct{} {
	shutdown := make(chan struct{})
	ticker := c.systems.Ticker()

	ShouldShutdown := event.Instance.Register("core.shouldShutdown", func(state *event.State) {
		if state.Application.Shutdown {
			shutdown <- struct{}{}
		}
	})

	go func() {
		for {
			event.Instance.CallHash(ShouldShutdown)
			ticker(time.Now())
		}
	}()
	return shutdown
}
