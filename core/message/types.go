package message

// _ go:generate go get golang.org/x/tools/cmd/stringer
//go:generate stringer -type=Type

type Type uint32

const (
	Keyboard Type = iota
	FramebufferResize
	Cursor
	Camera
)
