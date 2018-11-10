package application

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/raphy42/rodent/core/message"
)

type KeyboardEvent struct {
	Key      glfw.Key
	Scancode int
	Action   glfw.Action
	Mods     glfw.ModifierKey
}

func (k KeyboardEvent) Type() message.Type {
	return message.Keyboard
}

type FramebufferEvent struct {
	Width  int
	Height int
}

func (f FramebufferEvent) Type() message.Type {
	return message.Framebuffer
}

type CursorEvent struct {
	X, Y float32
}

func (c CursorEvent) Type() message.Type {
	return message.Cursor
}
