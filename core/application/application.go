package application

import (
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/raphy42/rodent/core/event"
	"github.com/raphy42/rodent/core/system"
	"github.com/raphy42/rodent/core/thread"
)

var (
	Instance = NewApplication()
)

type Config struct {
	Window struct {
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Title     string `json:"title"`
		Resizable bool   `json:"resizable"`
	} `json:"window"`
	OpenGL struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
	} `json:"open_gl"`
}

type Application struct {
	system.Module
	config Config
	window *glfw.Window
}

func NewApplication() *Application {
	app := Application{
		Module: system.NewModule("application", system.HighestPriority-1),
	}

	return &app
}

func (a *Application) PreInit(v interface{}) error {
	config := v.(Config)
	a.config = config

	return nil
}

func (a *Application) Init() error {
	config := a.config

	return thread.DoErr(func() error {
		if err := glfw.Init(); err != nil {
			return err
		}

		glfw.WindowHint(glfw.ContextVersionMajor, config.OpenGL.Major)
		glfw.WindowHint(glfw.ContextVersionMinor, config.OpenGL.Minor)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

		// glfw3 windows are resizable by default
		if !config.Window.Resizable {
			glfw.WindowHint(glfw.Resizable, glfw.False)
		}

		if runtime.GOOS == "darwin" {
			glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		}

		conf := a.config.Window
		window, err := glfw.CreateWindow(conf.Width, conf.Height, conf.Title, nil, nil)
		if err != nil {
			return err
		}
		a.window = window
		a.window.MakeContextCurrent()

		if err := gl.Init(); err != nil {
			return err
		}

		log.Printf("renderer: %s\nversion: %s\n", gl.GoStr(gl.GetString(gl.RENDERER)), gl.GoStr(gl.GetString(gl.VERSION)))

		return nil
	})
}

func (a *Application) PostInit() error {
	event.Instance.RegisterMultiple(actionMap)

	thread.Do(func() {
		a.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			go event.Instance.Call(func(state *event.State) {
				state.Input.Key = key
				state.Input.Action = action
				state.Input.Mods = mods
				state.Input.Scancode = scancode
			})
		})
	})
	return nil
}

func (a *Application) Ticker() func(time.Time) {
	ShouldClose, _ := event.Instance.Hash("application.shutdown")

	return func(delta time.Time) {
		if a.window.ShouldClose() {
			event.Instance.CallHash(ShouldClose)
		}

		// is it really the best place ?
		thread.Do(func() {
			a.window.SwapBuffers()
			glfw.PollEvents()
		})
	}
}
