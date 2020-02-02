package BasicQuadShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

const (
	screenTexture_uniformAddress = "u_screenTexture"
)

type BasicQuadShader struct {
	*Shader.ShaderProgramCore
}

func NewBasicQuadShader(vertexShaderFiles []string, fragmentShaderFiles []string) (*BasicQuadShader, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &BasicQuadShader{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewBasicIQuadShader(vertexShaderFiles []string, fragmentShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewBasicQuadShader(vertexShaderFiles, fragmentShaderFiles)
}

func (program *BasicQuadShader) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Model.Texture:
		return program.bindTexture(v)
	default:
		return fmt.Errorf("type %T not supported", v)
	}
}

func (program *BasicQuadShader) bindTexture(texture *Model.Texture) error {
	return program.BindTexture(0, texture, screenTexture_uniformAddress)
}
