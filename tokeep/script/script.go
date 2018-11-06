package script

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/raphy42/rodent/tokeep/assets"
	"github.com/raphy42/rodent/core/system"
	"github.com/yuin/gopher-lua"
)

var Instance = NewScriptSystem()

type ScriptSystem struct {
	system.Module
	L *lua.LState
	scripts map[string]string
	modules map[string]Module
}

func (s *ScriptSystem) PreInit(v interface{}) error {
	s.L = lua.NewState()
	for _, pair := range []struct {
		n string
		f lua.LGFunction
	}{
		{lua.LoadLibName, lua.OpenPackage}, // Must be first
		{lua.BaseLibName, lua.OpenBase},
		{lua.TabLibName, lua.OpenTable},
	} {
		if err := s.L.CallByParam(lua.P{
			Fn:      s.L.NewFunction(pair.f),
			NRet:    0,
			Protect: true,
		}, lua.LString(pair.n)); err != nil {
			return err
		}
	}
	return nil
}

func (s *ScriptSystem) Init() error {
	debug := &Debug{}
	s.RegisterModule("debug", debug)

	assets.Instance.RegisterBuilder("script", s.scriptBuilder)

	return nil
}

func (s ScriptSystem) PostInit() error {
	for name, module := range s.modules {
		s.L.PreloadModule(name, module.Load)
		log.Printf("loaded module '%s'\n", name)
	}

	for _, source := range s.scripts {
		s.L.DoString(source)
	}
	return nil
}

func (s ScriptSystem) Ticker() func(time.Time) {
	return func(delta time.Time) {
		
	}
}

func (s ScriptSystem) Shutdown() error {
	s.L.Close()
	return nil
}

func (s *ScriptSystem) scriptBuilder(name string, filenames []string, metadata map[string]string) error {
	root := metadata["root"]
	filename := filepath.Join(root, filenames[0])
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	source := string(buffer)
	s.scripts[name] = source
	return nil
}

func (s *ScriptSystem) RegisterModule(name string, module Module) {
	s.modules[name] = module
}

func NewScriptSystem() *ScriptSystem {
	return &ScriptSystem{
		Module: system.NewModule("script", system.LowestPriority),
		scripts: make(map[string]string),
		modules: make(map[string]Module),
	}
}
