package Shadow

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Light/DirectionalLight"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	modelMatrix_uniformAddress = "u_modelMatrix"

	directionalLightUBO_uniformAddress = "directionalLight"
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
	case *Model.Model:
		return program.BindUniform(v.ModelMatrix, modelMatrix_uniformAddress)
	case *DirectionalLight.UniformBuffer:
		return program.BindUniform(v, directionalLightUBO_uniformAddress)
	case *Model.Material, *Texture.Texture:
		return nil
	default:
		return fmt.Errorf("type %T not supported", v)
	}
}
