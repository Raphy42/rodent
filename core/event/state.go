package event

import "github.com/go-gl/glfw/v3.2/glfw"

type State struct {
	Application struct {
		Ready bool
		Shutdown bool
	}
	Input struct {
		Context int
		Action glfw.Action
		Key glfw.Key
		Scancode int
		Mods glfw.ModifierKey
	}
}

func (s *State) Apply(action Action) {
	action(s)
}

