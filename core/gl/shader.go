package gl

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	program uint32
}

func (s *Shader) Bind(state *State) {
	gl.UseProgram(s.program)
}

func (s *Shader) Unbind(state *State) {
	gl.UseProgram(0)
}

type ShaderBuilder struct {
	out      *Shader
	geometry bool
	sources  map[string]string
}

func NewShaderBuilder() *ShaderBuilder {
	return &ShaderBuilder{
		sources: make(map[string]string),
	}
}

func (s *ShaderBuilder) Source(xtype string, source string) *ShaderBuilder {
	s.sources[xtype] = source
	return s
}

func (s *ShaderBuilder) Sources(sources map[string]string) *ShaderBuilder {
	s.sources = sources
	return s
}

func (s *ShaderBuilder) Compile(state *State) (*Shader, error) {
	shaders := make([]uint32, 0)

	for xtype, source := range s.sources {
		var shader uint32
		var err error
		switch xtype {
		case "vert":
			shader, err = state.CompileShader(gl.VERTEX_SHADER, source)
		case "frag":
			shader, err = state.CompileShader(gl.FRAGMENT_SHADER, source)
		case "geom":
			shader, err = state.CompileShader(gl.GEOMETRY_SHADER, source)
		default:
			err = fmt.Errorf("invalid shader type: %s", xtype)
		}

		if err != nil {
			return nil, err
		}
		shaders = append(shaders, shader)
	}

	program, err := state.CompileProgram(shaders...)
	if err != nil {
		return nil, err
	}

	return &Shader{program: program}, nil
}

func (s *ShaderBuilder) MustCompile(state *State) *Shader {
	shader, err := s.Compile(state)
	if err != nil {
		panic(err)
	}
	return shader
}
