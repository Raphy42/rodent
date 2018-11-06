package renderer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func Background(rgba mgl32.Vec4) {
	gl.ClearColor(rgba.X(), rgba.Y(), rgba.Z(), rgba.W())
}

func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
}