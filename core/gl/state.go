package gl

import (
	"fmt"
	"hash/fnv"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type State struct {
	shaders map[uint32]uint32
}

func NewState() *State {
	return &State{
		shaders: make(map[uint32]uint32),
	}
}

func (s *State) CompileShader(xtype uint32, source string) (uint32, error) {
	hasher := fnv.New32a()
	hasher.Write([]byte(source))
	hash := hasher.Sum32()

	if shader, ok := s.shaders[hash]; ok {
		return shader, nil
	}

	shader := gl.CreateShader(xtype)
	if shader == 0 {
		return 0, fmt.Errorf("gl.CreateShader error")
	}
	src, free := gl.Strs(source)
	defer free()

	length := int32(len(source))
	gl.ShaderSource(shader, 1, src, &length)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status != gl.TRUE {
		var logLength int32
		gl.GetShaderiv(shader, gl.COMPILE_STATUS, &logLength)

		infoLog := make([]byte, logLength)
		gl.GetShaderInfoLog(shader, logLength, nil, &infoLog[0])

		return 0, CompileError{Log: string(infoLog), Source: source, Type: xtype}
	}

	return shader, nil
}

func (s *State) CompileProgram(shaders ...uint32) (uint32, error) {
	program := gl.CreateProgram()
	for _, shader := range shaders {
		gl.AttachShader(program, shader)
	}

	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status != gl.TRUE {
		var logLength int32
		gl.GetProgramiv(program, gl.COMPILE_STATUS, &logLength)

		infoLog := make([]byte, logLength)
		gl.GetProgramInfoLog(program, logLength, nil, &infoLog[0])

		return 0, CompileError{Log: string(infoLog), Type: gl.PROGRAM}
	}

	return program, nil
}
