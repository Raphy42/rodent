package gl

import "github.com/go-gl/gl/v4.1-core/gl"

func Renderer() string {
	return gl.GoStr(gl.GetString(gl.RENDERER))
}

func Version() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}