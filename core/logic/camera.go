package logic

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/raphy42/rodent/core/application"
	"github.com/raphy42/rodent/core/application/input"
	"github.com/raphy42/rodent/core/camera"
	"github.com/raphy42/rodent/core/message"
)

type Camera struct {
	*camera.Camera

	keys, cursor, scroll chan message.Message
}

func NewCamera() *Camera {
	return &Camera{Camera: camera.NewPerspective()}
}

func (c *Camera) Register(bus *message.EventBus) {
	c.keys = bus.Subscribe(message.Keyboard.String())
	c.cursor = bus.Subscribe(message.Cursor.String())
	c.scroll = bus.Subscribe(message.Scroll.String())

	updateMouse := false

	// keyboard events
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
			if key.Mods == input.Shift {
				c.Speed = mgl32.DegToRad(float32(int(c.Speed) ^ 3))
			}
		}
	}()

	// cursor position
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

	// mouse wheel
	go func() {
		for {
			ev := <-c.scroll
			scroll := ev.(*application.ScrollEvent)
			c.Zoom(scroll.Y)
		}
	}()
}
