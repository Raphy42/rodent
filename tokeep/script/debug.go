package script

import (
	"log"

	"github.com/yuin/gopher-lua"
)

type Debug struct {}

func (d *Debug) Load(L *lua.LState) int {
	module := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"print": d.print,
	})
	L.Push(module)
	return 1
}

func (d *Debug) print(L *lua.LState) int {
	lv := L.ToString(1)
	log.Printf("debug.print: %s", lv)
	return 1
}