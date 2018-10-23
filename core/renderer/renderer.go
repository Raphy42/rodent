package renderer

import (
	"github.com/raphy42/rodent/core/system"
)

var (
	Instance = NewRenderer()
)

type Renderer struct {
	system.Module
}

func NewRenderer() *Renderer {
	renderer := Renderer{
		Module: system.NewModule("renderer", 0),
	}
	return &renderer
}

func (r Renderer) PreInit(interface{}) error {
	return nil
}

func (r Renderer) Init() error {
	return nil
}

func (r Renderer) PostInit() error {
	panic("implement me")
}

