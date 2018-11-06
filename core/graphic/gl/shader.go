package gl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/pkg/errors"
)

type Shader struct {
	handle uint32
	build  func() error
}

type ShaderOptions struct {
	sources map[uint32]string
}

type ShaderOption func(*ShaderOptions)

func addSource(xtype uint32, src string) ShaderOption {
	return func(opt *ShaderOptions) {
		opt.sources[xtype] = src
	}
}

func Vertex(source string) ShaderOption {
	return addSource(gl.VERTEX_SHADER, source)
}

func Fragment(source string) ShaderOption {
	return addSource(gl.FRAGMENT_SHADER, source)
}

func Geometry(source string) ShaderOption {
	return addSource(gl.GEOMETRY_SHADER, source)
}

func compileShader(xtype uint32, source string) (uint32, error) {
	shader := gl.CreateShader(xtype)
	src, free := gl.Strs(source)
	defer free()

	length := int32(len(source))
	gl.ShaderSource(shader, 1, src, &length)
	gl.CompileShader(shader)

	ok := int32(0)
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &ok)
	if ok != gl.TRUE {
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &length)
		buffer := make([]byte, length)
		gl.GetShaderInfoLog(shader, length, nil, &buffer[0])
		return 0, errors.New(string(buffer))
	}
	return shader, nil
}

func createProgram(shaders ...uint32) (uint32, error) {
	program := gl.CreateProgram()
	for _, shader := range shaders {
		gl.AttachShader(program, shader)
	}

	gl.LinkProgram(program)

	ok := int32(0)
	gl.GetProgramiv(program, gl.LINK_STATUS, &ok)
	if ok != gl.TRUE {
		length := int32(0)
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &length)
		buffer := make([]byte, length)
		gl.GetProgramInfoLog(program, length, nil, &buffer[0])
		return 0, errors.New(string(buffer))
	}
	return program, nil
}

func NewShader(options ...ShaderOption) *Shader {
	defaultOpts := ShaderOptions{sources: make(map[uint32]string)}
	for _, opt := range options {
		opt(&defaultOpts)
	}

	s := new(Shader)
	s.build = func() error {
		shaders := make([]uint32, 0)
		for xtype, source := range defaultOpts.sources {
			shader, err := compileShader(xtype, source)
			if err != nil {
				return errors.Wrap(err, source)
			}
			shaders = append(shaders, shader)
		}
		program, err := createProgram(shaders...)
		if err != nil {
			return errors.Wrap(err, "unable to link program")
		}

		s.handle = program

		return nil
	}

	return s
}

func (s Shader) Build() error {
	return s.build()
}

func (s Shader) Bind() {
	gl.UseProgram(s.handle)
}

func (s Shader) uniformLocation(name string) int32 {
	location := gl.GetUniformLocation(s.handle, gl.Str(name+"\x00"))
	return location
}

func (s Shader) SetMat4(name string, mat mgl32.Mat4) {
	s.Bind()
	gl.UniformMatrix4fv(s.uniformLocation(name), 1, false, &mat[0])
}
