package application

type windowOptions struct {
	Width int `json:"width"`
	Height int `json:"height"`
	Resizable bool `json:"resizable"`
	GLMajor int `json:"gl_major"`
	GLMinor int `json:"gl_minor"`
}

type Options struct {
	Window windowOptions `json:"window"`
}

var (
	defaultOptions = Options{
		Window:windowOptions{
			Width: 1200,
			Height: 800,
		},
	}
)

type Option func (*Options)

func WindowSize(width, height int) Option {
	return func(options *Options) {
		options.Window.Height = height
		options.Window.Width = width
	}
}

func Resizable(v bool) Option {
	return func(options *Options) {
		options.Window.Resizable = v
	}
}

func GLVersion(major, minor int) Option {
	return func(options *Options) {
		options.Window.GLMajor = major
		options.Window.GLMinor = minor
	}
}