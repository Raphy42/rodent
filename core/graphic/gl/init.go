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
	gl.Enable(gl.CULL_FACE)
	// gl.CullFace(gl.FRONT)
	// gl.DepthFunc(gl.MO)
	return nil
}
