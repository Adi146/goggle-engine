package ShadowMapShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Light/DirectionalLight"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	ua_modelMatrix = "u_modelMatrix"

	ua_directionalLight = "directionalLight"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore
	MaterialShader Material.ShaderProgram
}

func NewShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (*ShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &ShaderProgram{
		ShaderProgramCore: shaderCore,
		MaterialShader: Material.ShaderProgram{
			ShaderProgramCore: shaderCore,
			BindFunctions: []func(program *Material.ShaderProgram, material *Material.Material) error{
				Material.BindDiffuse,
			},
		},
	}, nil
}

func NewIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewShaderProgram(vertexShaderFiles, fragmentShaderFiles)
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Material.Material:
		return program.MaterialShader.BindObject(v)
	case *Model.Model:
		return program.BindUniform(v.ModelMatrix, ua_modelMatrix)
	case *DirectionalLight.UniformBuffer:
		return program.BindUniform(v, ua_directionalLight)
	case *Texture.Texture:
		return nil
	default:
		return fmt.Errorf("type %T not supported", v)
	}
}
