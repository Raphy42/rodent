package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/pkg/errors"
)

func Init() error {
	if err := gl.Init(); err != nil {
		return errors.Wrap(err, "OpenGL initialisation")
	}
	gl.Enable(gl.DEPTH_TEST)
	// gl.DepthFunc(gl.MO)
	return nil
}