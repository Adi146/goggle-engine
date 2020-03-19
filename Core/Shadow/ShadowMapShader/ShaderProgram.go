package ShadowMapShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	shader_factory_name = "shadowMapShader"

	ua_modelMatrix = "u_modelMatrix"

	ua_directionalLight = "directionalLight"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore
	MaterialShader Material.ShaderProgram
}

func init() {
	Shader.Factory.AddConstructor(shader_factory_name, NewIShaderProgram)
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
	case UniformBuffer.IUniformBuffer:
		switch t := v.GetType(); t {
		case Light.DirectionalLight_ubo_type:
			return program.BindUniform(v, ua_directionalLight)
		default:
			return fmt.Errorf("shadow map shader does not support uniform buffers of type %s", t)
		}
	case Texture.ITexture:
		return nil
	default:
		return fmt.Errorf("shadow map shader does not support type %T", v)
	}
}
