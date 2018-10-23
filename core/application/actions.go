package application

import "github.com/raphy42/rodent/core/event"

var actionMap = map[string]event.Action {
	"application.shutdown": func(state *event.State) {
		state.Application.Shutdown = true
	},
}
