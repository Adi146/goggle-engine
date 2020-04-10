package PhongShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	shader_factory_name = "phongShader"

	ua_modelMatrix  = "u_modelMatrix"
	ua_normalMatrix = "u_normalMatrix"

	ua_camera           = "camera"
	ua_directionalLight = "directionalLight"
	ua_pointLight       = "pointLight"
	ua_spotLight        = "spotLight"
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
	case *GeometryMath.Matrix3x3:
		return ua_normalMatrix, nil
	case Texture.ITexture:
		return program.ShadowShader.GetUniformAddress(v)
	case UniformBuffer.IUniformBuffer:
		switch t := v.GetType(); t {
		case Light.DirectionalLight_ubo_type:
			return ua_directionalLight, nil
		case Light.PointLight_ubo_type:
			return ua_pointLight, nil
		case Light.SpotLight_ubo_type:
			return ua_spotLight, nil
		case Camera.UBO_type:
			return ua_camera, nil
		default:
			return "", fmt.Errorf("phong shader does not support uniform buffers of type %s", t)
		}
	default:
		return "", fmt.Errorf("phong shader does not support type %T", v)
	}
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Material.Material, *Material.BlendMaterial:
		return program.MaterialShader.BindObject(v)
	case Texture.ITexture:
		return program.ShadowShader.BindObject(v)
	default:
		uniformAddress, err := program.GetUniformAddress(v)
		if err != nil {
			return err
		}
		return program.BindUniform(v, uniformAddress)
	}
}
