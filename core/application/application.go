package application

import (
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/pkg/errors"
	"github.com/raphy42/rodent/core/application/input"
	"github.com/raphy42/rodent/core/graphic/gl"
	"github.com/raphy42/rodent/core/logger"
	"github.com/raphy42/rodent/core/message"
	"go.uber.org/zap"
)

var log = logger.New()

type Application struct {
	window  *glfw.Window
	options windowOptions
	title   chan string
}

func New(options ...Option) *Application {
	config := defaultOptions
	for _, option := range options {
		option(&config)
	}
	return &Application{
		options: config.Window,
		title:   make(chan string, 1),
	}
}

func coerceFlag(v bool) int {
	if v {
		return 1
	}
	return 0
}

func (a *Application) Init() error {
	if err := glfw.Init(); err != nil {
		return err
	}

	log.Info("glfw initialised")

	config := a.options

	glfw.WindowHint(glfw.ContextVersionMajor, config.GLMajor)
	glfw.WindowHint(glfw.ContextVersionMinor, config.GLMinor)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable, coerceFlag(config.Resizable))

	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	}

	window, err := glfw.CreateWindow(config.Width, config.Height, "default", nil, nil)
	if err != nil {
		return err
	}

	log.Debug("window created",
		zap.Int("width", config.Width),
		zap.Int("height", config.Height),
	)

	a.window = window
	a.window.MakeContextCurrent()

	glfw.SwapInterval(-1)

	if err := gl.Init(); err != nil {
		return errors.Wrapf(err, "opengl bootstrap version:%d.%d", config.GLMinor, config.GLMajor)
	}
	log.Info("OpenGL initialised",
		zap.String("renderer", gl.Renderer()),
		zap.String("version", gl.Version()),
	)

	return nil
}

func (a *Application) Dispose() {
	glfw.Terminate()
}

func (a *Application) Tick(delta time.Time) {
	glfw.PollEvents()
	a.window.SwapBuffers()
}

func (a *Application) ShouldShutdown() bool {
	return a.window.ShouldClose()
}

func (a *Application) RegisterEvents(bus message.Bus) {
	keyboardEvents := bus.Publish(message.Keyboard)
	a.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		go func() {
			keyboardEvents <- &input.KeyboardAction{
				Key: input.Key(key), Scancode: input.Scancode(scancode),
				Action: input.Action(action), Mods: input.Mods(mods),
			}
		}()
	})

	framebufferEvents := bus.Publish(message.FramebufferResize)
	a.window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		go func() {
			framebufferEvents <- &FramebufferEvent{Width: width, Height: height}
		}()
	})

	cursorEvents := bus.Publish(message.Cursor)
	a.window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		go func() {
			cursorEvents <- &CursorEvent{X: float32(xpos), Y: float32(ypos)}
		}()
	})

	scrollEvents := bus.Publish(message.Scroll)
	a.window.SetScrollCallback(func(w *glfw.Window, xoff float64, yoff float64) {
		go func() {
			scrollEvents <- &ScrollEvent{X: float32(xoff), Y: float32(yoff)}
		}()
	})
}
