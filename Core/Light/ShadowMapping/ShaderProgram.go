package ShadowMapping

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	shader_factory_name = "shadowMapShader"

	ua_modelMatrix = "u_modelMatrix"

	ua_directionalLight = "directionalLight"
	ua_pointLight       = "pointLight"
	ua_spotLight        = "spotLight"
	ua_lightIndex       = "u_lightIndex"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore
	MaterialShader Material.ShaderProgram
}

func init() {
	Shader.Factory.AddConstructor(shader_factory_name, NewIShaderProgram)
}

func NewShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string, geometryShaderFiles []string) (*ShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles, geometryShaderFiles)
	if err != nil {
		return nil, err
	}

	return &ShaderProgram{
		ShaderProgramCore: shaderCore,
		MaterialShader: Material.ShaderProgram{
			ShaderProgramCore: shaderCore,
			BindFunctions: []func(program *Material.ShaderProgram, material *Material.Material, uaMaterial string) error{
				Material.BindDiffuse,
			},
		},
	}, nil
}

func NewIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string, geometryShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewShaderProgram(vertexShaderFiles, fragmentShaderFiles, geometryShaderFiles)
}

func (program *ShaderProgram) GetUniformAddress(i interface{}) (string, error) {
	switch v := i.(type) {
	case *Material.Material, *Material.BlendMaterial:
		return program.MaterialShader.GetUniformAddress(i)
	case *GeometryMath.Matrix4x4:
		return ua_modelMatrix, nil
	case UniformBuffer.IUniformBuffer:
		switch t := v.GetType(); t {
		case Light.DirectionalLight_ubo_type:
			return ua_directionalLight, nil
		case Light.PointLight_ubo_type:
			return ua_pointLight, nil
		case Light.SpotLight_ubo_type:
			return ua_spotLight, nil
		default:
			return "", fmt.Errorf("shadow map shader does not support unifrom buffer of type %s", t)
		}
	case int32:
		return ua_lightIndex, nil
	default:
		return "", fmt.Errorf("shadow map shader does not support type %T", v)
	}
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Material.Material, *Material.BlendMaterial:
		return program.MaterialShader.BindObject(v)
	case Texture.ITexture:
		return nil
	case *GeometryMath.Matrix3x3:
		return nil
	case Mesh.IVertexArray:
		v.Bind()
		v.EnableUVAttribute()
		return nil
	default:
		uniformAddress, err := program.GetUniformAddress(v)
		if err != nil {
			return err
		}
		return program.BindUniform(v, uniformAddress)
	}
}
