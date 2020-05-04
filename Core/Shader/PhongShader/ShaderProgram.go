package PhongShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	shader_factory_name = "phongShader"

	ua_modelMatrix = "u_modelMatrix"
)

func init() {
	Shader.Factory.AddConstructor(shader_factory_name, NewIShaderProgram)
}

type ShaderProgram struct {
	*Shader.ShaderProgramCore
	ShadowShader   ShadowMapping.ShaderComponent
	MaterialShader Material.ShaderProgram
}

func NewShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string, geometryShaderFiles []string) (*ShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles, geometryShaderFiles)
	if err != nil {
		return nil, err
	}

	return &ShaderProgram{
		ShaderProgramCore: shaderCore,
		ShadowShader: ShadowMapping.ShaderComponent{
			ShaderProgramCore: shaderCore,
		},
		MaterialShader: Material.ShaderProgram{
			ShaderProgramCore: shaderCore,
			BindFunctions: []func(program *Material.ShaderProgram, material *Material.Material, uaMaterial string) error{
				Material.BindDiffuse,
				Material.BindSpecular,
				Material.BindEmissive,
				Material.BindNormals,
				Material.BindShininess,
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
		return program.MaterialShader.GetUniformAddress(v)
	case *GeometryMath.Matrix4x4:
		return ua_modelMatrix, nil
	case Texture.ITexture:
		return program.ShadowShader.GetUniformAddress(v)
	default:
		return "", fmt.Errorf("GetUniformAddress of phong shader does not support type %T, try to use BindObject instead", v)
	}
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Material.Material, *Material.BlendMaterial:
		return program.MaterialShader.BindObject(v)
	case Texture.ITexture:
		return program.ShadowShader.BindObject(v)
	case Mesh.IVertexArray:
		v.Bind()
		v.EnableUVAttribute()
		v.EnableNormalAttribute()
		v.EnableTangentAttribute()
		v.EnableBiTangentAttribute()
		return nil
	default:
		uniformAddress, err := program.GetUniformAddress(v)
		if err != nil {
			return err
		}
		return program.BindUniform(v, uniformAddress)
	}
}
