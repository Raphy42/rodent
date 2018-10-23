package graphic

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/raphy42/rodent/core/assets"
	"github.com/raphy42/rodent/core/gl"
	"github.com/raphy42/rodent/core/system"
	"github.com/raphy42/rodent/core/thread"
)

var Instance = NewGraphic()

type Graphic struct {
	system.Module
	state   *gl.State
	shaders map[string]*gl.Shader
}

func NewGraphic() *Graphic {
	return &Graphic{
		Module:  system.NewModule("graphic", 0),
		state:   gl.NewState(),
		shaders: make(map[string]*gl.Shader),
	}
}

func (g Graphic) PreInit(interface{}) error {
	return nil
}

func (g Graphic) Init() error {
	assets.Instance.RegisterBuilder("shader", g.shaderBuilder)
	return nil
}

func (g Graphic) PostInit() error {
	return nil
}

func (g Graphic) Ticker() func(time.Time) {
	return func(delta time.Time) {

	}
}

func (g *Graphic) shaderBuilder(name string, filenames []string, metadata map[string]string) error {
	root := metadata["root"]
	builder := gl.NewShaderBuilder()

	for _, filename := range filenames {
		base := filepath.Base(filename)
		pack := strings.Split(base, ".")
		if len(pack) != 3 {
			return fmt.Errorf("expected a shader filename like '{name}.{type}.glsl' got %s", filename)
		}
		xtype := pack[1]
		buffer, err := ioutil.ReadFile(filepath.Join(root, filename))
		if err != nil {
			return err
		}
		builder.Source(xtype, string(buffer))
	}

	out := make(chan *gl.Shader, 1)
	thread.Do(func() {
		shader := builder.MustCompile(g.state)
		out <- shader
	})
	g.shaders[name] = <-out
	log.Printf("built shader: %s\n", name)
	return nil
}
