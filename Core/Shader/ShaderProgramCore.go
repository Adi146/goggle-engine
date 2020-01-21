package Shader

import (
	"C"
	"fmt"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

type ShaderProgramCore struct {
	programId      uint32
	vertexShader   *Shader
	fragmentShader *Shader
	isBound        bool
}

func NewShaderProgramFromFiles(vertexShaderFile string, fragmentShaderFile string) (*ShaderProgramCore, error) {
	vertexShader, err := NewShaderFromFile(vertexShaderFile, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := NewShaderFromFile(fragmentShaderFile, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	program, err := NewShaderProgram(vertexShader, fragmentShader)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func NewShaderProgram(vertexShader *Shader, fragmentShader *Shader) (*ShaderProgramCore, error) {
	program := ShaderProgramCore{
		programId:      gl.CreateProgram(),
		vertexShader:   vertexShader,
		fragmentShader: fragmentShader,
		isBound:        false,
	}

	gl.AttachShader(program.programId, vertexShader.shaderId)
	gl.AttachShader(program.programId, fragmentShader.shaderId)
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
	gl.DetachShader(program.programId, program.vertexShader.shaderId)
	gl.DetachShader(program.programId, program.fragmentShader.shaderId)
	program.vertexShader.Destroy()
	program.fragmentShader.Destroy()
}

func (program *ShaderProgramCore) Bind() {
	gl.UseProgram(program.programId)
	program.isBound = true
}

func (program *ShaderProgramCore) Unbind() {
	gl.UseProgram(0)
	program.isBound = false
}

func (program *ShaderProgramCore) BeginDraw() error {
	return nil
}

func (program *ShaderProgramCore) EndDraw() {
}

func (program *ShaderProgramCore) BindUniform(i interface{}, uniformAddress string) error {
	location, err := program.getUniformLocation(uniformAddress)
	if err != nil {
		return err
	}

	switch v := i.(type) {
	case *Matrix.Matrix4x4:
		gl.UniformMatrix4fv(location, 1, false, &v[0][0])
	case *Vector.Vector3:
		gl.Uniform3fv(location, 1, &v[0])
	case float32:
		gl.Uniform1f(location, v)
	case int32:
		gl.Uniform1i(location, v)
	default:
		return fmt.Errorf("type %T not supported", v)
	}

	return nil
}

func (program *ShaderProgramCore) BindTexture(textureSlot uint32, texture *Model.Texture, uniformAddress string) error {
	location, err := program.getUniformLocation(uniformAddress)
	if err != nil {
		return err
	}

	gl.Uniform1i(location, int32(textureSlot))
	gl.ActiveTexture(gl.TEXTURE0 + textureSlot)
	texture.Bind()

	return nil
}

func (program *ShaderProgramCore) getUniformLocation(uniformAddress string) (int32, error) {
	if !program.isBound {
		return -1, fmt.Errorf("shader is not bound")
	}

	location := gl.GetUniformLocation(program.programId, gl.Str(uniformAddress+"\x00"))
	if location == -1 {
		return location, fmt.Errorf("uniform address %s not found", uniformAddress)
	}

	return location, nil
}
