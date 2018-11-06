package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/raphy42/rodent/core/logger"
)

var log = logger.New()

func Viewport(width, height int32) {
	gl.Viewport(0, 0, width, height)
}
