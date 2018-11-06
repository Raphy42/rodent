package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/pkg/errors"
)

type State struct {
}

func NewState() *State {
	if err := gl.Init(); err != nil {
		panic(errors.Wrap(err, "unable to init opengl"))
	}
	return &State{}
}

