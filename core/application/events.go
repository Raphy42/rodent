package application

import "github.com/go-gl/glfw/v3.2/glfw"

type KeyboardEvent struct {
	Key      glfw.Key
	Scancode int
	Action   glfw.Action
	Mods     glfw.ModifierKey
}

func (k *KeyboardEvent) Typename() string {
	return "keyboard"
}

type FramebufferEvent struct {
	Width int
	Height int
}

func (f *FramebufferEvent) Typename() string {
	return "framebuffer"
}