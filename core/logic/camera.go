package logic

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/raphy42/rodent/core/application"
	"github.com/raphy42/rodent/core/application/input"
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
	c.keys = bus.Subscribe(message.Keyboard.String())
	c.cursor = bus.Subscribe(message.Cursor.String())

	updateMouse := false

	go func() {
		for {
			ev := <-c.keys
			key := ev.(*input.KeyboardAction)
			delta := float32(glfw.GetTime())
			if key.IsReleased() {
				continue
			}
			switch key.Key {
			case input.KeyUp, input.KeyW:
				c.Move(camera.Forward, delta)
			case input.KeyDown, input.KeyS:
				c.Move(camera.Backward, delta)
			case input.KeyLeft, input.KeyA:
				c.Move(camera.Left, delta)
			case input.KeyRight, input.KeyD:
				c.Move(camera.Right, delta)
			case input.KeySpace:
				updateMouse = !updateMouse
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
