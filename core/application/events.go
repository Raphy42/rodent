package application

import (
	"github.com/raphy42/rodent/core/message"
)

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

type ScrollEvent struct {
	X, Y float32
}

func (s ScrollEvent) Type() message.Type {
	return message.Scroll
}
