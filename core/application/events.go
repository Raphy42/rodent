package application

import (
	"github.com/raphy42/rodent/core/application/input"
	"github.com/raphy42/rodent/core/message"
)

type KeyboardEvent input.KeyboardAction

func (k KeyboardEvent) Type() message.Type {
	return message.Keyboard
}

type FramebufferEvent struct {
	Width  int
	Height int
}

func (f FramebufferEvent) Type() message.Type {
	return message.FramebufferResize
}

type CursorEvent struct {
	X, Y float32
}

func (c CursorEvent) Type() message.Type {
	return message.Cursor
}
