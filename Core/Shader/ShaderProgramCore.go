package Shader

import (
	"C"
	"fmt"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"strings"

	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/go-gl/gl/v4.3-core/gl"
)

type ShaderProgramCore struct {
	ID              uint32
	vertexShaders   []*Shader
	fragmentShaders []*Shader
	geometryShaders []*Shader
}

func NewShaderProgramFromFiles(vertexShaderFiles []string, fragmentShaderFiles []string, geometryShaderFiles []string) (*ShaderProgramCore, error) {
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

	var geometryShaders []*Shader
	for _, geometryShaderFile := range geometryShaderFiles {
		geometryShader, err := NewShaderFromFile(geometryShaderFile, gl.GEOMETRY_SHADER)
		if err != nil {
			return nil, err
		}
		geometryShaders = append(geometryShaders, geometryShader)
	}

	program, err := NewShaderProgram(vertexShaders, fragmentShaders, geometryShaders)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func NewShaderProgram(vertexShaders []*Shader, fragmentShaders []*Shader, geometryShaders []*Shader) (*ShaderProgramCore, error) {
	program := ShaderProgramCore{
		ID:              gl.CreateProgram(),
		vertexShaders:   vertexShaders,
		fragmentShaders: fragmentShaders,
		geometryShaders: geometryShaders,
	}

	for _, vertexShader := range vertexShaders {
		gl.AttachShader(program.ID, vertexShader.shaderId)
	}
	for _, fragmentShader := range fragmentShaders {
		gl.AttachShader(program.ID, fragmentShader.shaderId)
	}
	for _, geometryShader := range geometryShaders {
		gl.AttachShader(program.ID, geometryShader.shaderId)
	}

	gl.LinkProgram(program.ID)

	var status int32
	gl.GetProgramiv(program.ID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program.ID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program.ID, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	return &program, nil
}

func (program *ShaderProgramCore) Destroy() {
	for _, vertexShader := range program.vertexShaders {
		gl.DetachShader(program.ID, vertexShader.shaderId)
		vertexShader.Destroy()
	}
	for _, fragmentShader := range program.fragmentShaders {
		gl.DetachShader(program.ID, fragmentShader.shaderId)
		fragmentShader.Destroy()
	}
}

func (program *ShaderProgramCore) Bind() {
	gl.UseProgram(program.ID)
}

func (program *ShaderProgramCore) Unbind() {
	gl.UseProgram(0)
}

func (program *ShaderProgramCore) BindUniform(i interface{}, uniformAddress string) error {
	switch v := i.(type) {
	case IUniformBlock:
		index, err := program.getUniformBlockIndex(uniformAddress)
		if err != nil {
			return err
		}
		gl.UniformBlockBinding(program.ID, index, v.GetBinding())
	default:
		location, err := program.getUniformLocation(uniformAddress)
		if err != nil {
			return err
		}

		switch v := i.(type) {
		case *GeometryMath.Matrix4x4:
			gl.ProgramUniformMatrix4fv(program.ID, location, 1, false, &v[0][0])
		case *GeometryMath.Matrix3x3:
			gl.ProgramUniformMatrix3fv(program.ID, location, 1, false, &v[0][0])
		case *GeometryMath.Vector3:
			gl.ProgramUniform3fv(program.ID, location, 1, &v[0])
		case *GeometryMath.Vector4:
			gl.ProgramUniform4fv(program.ID, location, 1, &v[0])
		case float32:
			gl.ProgramUniform1f(program.ID, location, v)
		case []float32:
			gl.ProgramUniform1fv(program.ID, location, int32(len(v)), &v[0])
		case int32:
			gl.ProgramUniform1i(program.ID, location, v)
		case uint32:
			gl.ProgramUniform1i(program.ID, location, int32(v))
		case bool:
			if v {
				gl.ProgramUniform1i(program.ID, location, int32(1))
			} else {
				gl.ProgramUniform1i(program.ID, location, int32(0))
			}
		case Texture.ITexture:
			if err := v.Bind(); err != nil {
				return err
			}

			return program.BindUniform(v.GetUnit().ID, uniformAddress)
		default:
			return fmt.Errorf("type %T not supported", v)
		}
	}

	return nil
}

func (program *ShaderProgramCore) getUniformLocation(uniformAddress string) (int32, error) {
	location := gl.GetUniformLocation(program.ID, gl.Str(uniformAddress+"\x00"))
	if location == -1 {
		return location, fmt.Errorf("uniform address %s not found", uniformAddress)
	}

	return location, nil
}

func (program *ShaderProgramCore) getUniformBlockIndex(uniformAddress string) (uint32, error) {
	index := gl.GetUniformBlockIndex(program.ID, gl.Str(uniformAddress+"\x00"))
	if index == gl.INVALID_INDEX {
		return index, fmt.Errorf("uniform block %s not found", uniformAddress)
	}

	return index, nil
}
