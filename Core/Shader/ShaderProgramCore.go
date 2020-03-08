package Shader

import (
	"C"
	"fmt"
	"strings"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type ShaderProgramCore struct {
	programId       uint32
	vertexShaders   []*Shader
	fragmentShaders []*Shader
}

func NewShaderProgramFromFiles(vertexShaderFiles []string, fragmentShaderFiles []string) (*ShaderProgramCore, error) {
	var vertexShaders []*Shader
	for _, vertexShaderFile := range vertexShaderFiles {
		vertexShader, err := NewShaderFromFile(vertexShaderFile, gl.VERTEX_SHADER)
		if err != nil {
			return nil, err
		}
		vertexShaders = append(vertexShaders, vertexShader)
	}

	var fragmentShaders []*Shader
	for _, fragmentShaderFile := range fragmentShaderFiles {
		fragmentShader, err := NewShaderFromFile(fragmentShaderFile, gl.FRAGMENT_SHADER)
		if err != nil {
			return nil, err
		}
		fragmentShaders = append(fragmentShaders, fragmentShader)
	}

	program, err := NewShaderProgram(vertexShaders, fragmentShaders)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func NewShaderProgram(vertexShaders []*Shader, fragmentShaders []*Shader) (*ShaderProgramCore, error) {
	program := ShaderProgramCore{
		programId:       gl.CreateProgram(),
		vertexShaders:   vertexShaders,
		fragmentShaders: fragmentShaders,
	}

	for _, vertexShader := range vertexShaders {
		gl.AttachShader(program.programId, vertexShader.shaderId)
	}
	for _, fragmentShader := range fragmentShaders {
		gl.AttachShader(program.programId, fragmentShader.shaderId)
	}

	gl.LinkProgram(program.programId)

	var status int32
	gl.GetProgramiv(program.programId, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program.programId, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program.programId, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	return &program, nil
}

func (program *ShaderProgramCore) Destroy() {
	for _, vertexShader := range program.vertexShaders {
		gl.DetachShader(program.programId, vertexShader.shaderId)
		vertexShader.Destroy()
	}
	for _, fragmentShader := range program.fragmentShaders {
		gl.DetachShader(program.programId, fragmentShader.shaderId)
		fragmentShader.Destroy()
	}
}

func (program *ShaderProgramCore) Bind() {
	gl.UseProgram(program.programId)
}

func (program *ShaderProgramCore) Unbind() {
	gl.UseProgram(0)
}

func (program *ShaderProgramCore) BindUniform(i interface{}, uniformAddress string) error {
	switch v := i.(type) {
	case UniformBuffer.IUniformBuffer:
		index, err := program.getUniformBlockIndex(uniformAddress)
		if err != nil {
			return err
		}
		gl.UniformBlockBinding(program.programId, index, v.GetIndex())
	default:
		var currentProgram int32
		if gl.GetIntegerv(gl.CURRENT_PROGRAM, &currentProgram); uint32(currentProgram) != program.programId {
			return fmt.Errorf("shader is not bound")
		}

		location, err := program.getUniformLocation(uniformAddress)
		if err != nil {
			return err
		}

		switch v := i.(type) {
		case *Matrix.Matrix4x4:
			gl.UniformMatrix4fv(location, 1, false, &v[0][0])
		case *Vector.Vector3:
			gl.Uniform3fv(location, 1, &v[0])
		case *Vector.Vector4:
			gl.Uniform4fv(location, 1, &v[0])
		case float32:
			gl.Uniform1f(location, v)
		case []float32:
			gl.Uniform1fv(location, int32(len(v)), &v[0])
		case int32:
			gl.Uniform1i(location, v)
		default:
			return fmt.Errorf("type %T not supported", v)
		}
	}

	return nil
}

func (program *ShaderProgramCore) BindTexture(texture Texture.ITexture, uniformAddress string, reserve bool) error {
	unit, err := Texture.FindFreeUnit(texture)
	if err != nil {
		return err
	}

	err = program.BindUniform(int32(unit), uniformAddress)
	if err != nil {
		return err
	}

	Texture.BindTexture(texture, unit, reserve)
	return nil
}

func (program *ShaderProgramCore) getUniformLocation(uniformAddress string) (int32, error) {
	location := gl.GetUniformLocation(program.programId, gl.Str(uniformAddress+"\x00"))
	if location == -1 {
		return location, fmt.Errorf("uniform address %s not found", uniformAddress)
	}

	return location, nil
}

func (program *ShaderProgramCore) getUniformBlockIndex(uniformAddress string) (uint32, error) {
	index := gl.GetUniformBlockIndex(program.programId, gl.Str(uniformAddress+"\x00"))
	if index == gl.INVALID_INDEX {
		return index, fmt.Errorf("uniform block %s not found", uniformAddress)
	}

	return index, nil
}
