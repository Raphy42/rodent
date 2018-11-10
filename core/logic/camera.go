package logic

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/raphy42/rodent/core/application"
	"github.com/raphy42/rodent/core/camera"
	"github.com/raphy42/rodent/core/message"
)

type Camera struct {
	*camera.Camera

	keys   chan message.Message
	cursor chan message.Message
}

func NewCamera() *Camera {
	return &Camera{Camera: camera.NewPerspective()}
}

func (c *Camera) Register(bus *message.EventBus) {
	c.keys = bus.Channel(message.Keyboard.String())
	c.cursor = bus.Channel(message.Cursor.String())

	updateMouse := false

	go func() {
		for {
			ev := <-c.keys
			key := ev.(*application.KeyboardEvent)
			delta := float32(glfw.GetTime())
			switch key.Action {
			case glfw.Press, glfw.Repeat:
				switch key.Key {
				case glfw.KeyUp, glfw.KeyW:
					c.Move(camera.Forward, delta)
				case glfw.KeyDown, glfw.KeyS:
					c.Move(camera.Backward, delta)
				case glfw.KeyLeft, glfw.KeyA:
					c.Move(camera.Left, delta)
				case glfw.KeyRight, glfw.KeyD:
					c.Move(camera.Right, delta)
				case glfw.KeySpace:
					updateMouse = !updateMouse
				}
			}

		}
	}()

	go func() {
		lastX := float32(0)
		lastY := float32(0)
		for {
			ev := <-c.cursor
			cursor := ev.(*application.CursorEvent)
			if updateMouse {
				c.Center(cursor.X-lastX, cursor.Y-lastY)
			}
			lastX = cursor.X
			lastY = cursor.Y
		}
	}()
}
