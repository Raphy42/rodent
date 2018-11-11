package input

import "github.com/raphy42/rodent/core/message"

func (ka *KeyboardAction) Is(key Key) bool {
	return ka.Key == key
}

func (ka KeyboardAction) OneOf(keys ...Key) bool {
	for _, key := range keys {
		if ka.Key == key {
			return true
		}
	}
	return false
}

func (ka KeyboardAction) IsPressed() bool {
	return ka.Action == Press
}

func (ka KeyboardAction) IsRepeated() bool {
	return ka.Action == Repeat
}

func (ka KeyboardAction) IsReleased() bool {
	return ka.Action == Release
}

func (ka KeyboardAction) Type() message.Type {
	return message.Keyboard
}
