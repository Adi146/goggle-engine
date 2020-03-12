package Skybox

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"

	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	ua_skybox = "u_skybox"

	ua_camera = "camera"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore
}

func NewShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (*ShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &ShaderProgram{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewShaderProgram(vertexShaderFiles, fragmentShaderFiles)
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case Texture.ITexture:
		return program.BindUniform(v, ua_skybox)
	case UniformBuffer.IUniformBuffer:
		switch t := v.GetType(); t {
		case Camera.UBO_type:
			return program.BindUniform(v, ua_camera)
		default:
			return fmt.Errorf("skybox shader does not support uniform buffers of type %s", t)
		}
	default:
		return fmt.Errorf("skybox shader does not support type %T", v)
	}
}
