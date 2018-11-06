package script

import (
	"github.com/yuin/gopher-lua"
)

type Module interface {
	Load(L *lua.LState) int
}
