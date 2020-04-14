package PhongShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

const (
	shader_factory_name = "phongShader"

	ua_modelMatrix  = "u_modelMatrix"
	ua_normalMatrix = "u_normalMatrix"

	ua_camera           = "camera"
	ua_directionalLight = "directionalLight"
	ua_pointLight       = "pointLight"
	ua_spotLight        = "spotLight"

	ua_directionalLightIsSet = "u_directionalLightIsSet"
	ua_pointLightIsSet       = "u_pointLightIsSet"
	ua_spotLightIsSet        = "u_spotLightIsSet"
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
	case UniformBuffer.IUniformBuffer:
		var err Error.ErrorCollection
		switch t := v.GetType(); t {
		case Light.DirectionalLight_ubo_type:
			err.Push(program.BindUniform(v, ua_directionalLight))
			err.Push(program.BindUniform(true, ua_directionalLightIsSet))
		case Light.PointLight_ubo_type:
			err.Push(program.BindUniform(v, ua_pointLight))
			err.Push(program.BindUniform(true, ua_pointLightIsSet))
		case Light.SpotLight_ubo_type:
			err.Push(program.BindUniform(v, ua_spotLight))
			err.Push(program.BindUniform(true, ua_spotLightIsSet))
		case Camera.UBO_type:
			err.Push(program.BindUniform(v, ua_camera))
		default:
			return fmt.Errorf("phong shader does not support uniform buffers of type %s", t)
		}
		return err.Err()
	default:
		uniformAddress, err := program.GetUniformAddress(v)
		if err != nil {
			return err
		}
		return program.BindUniform(v, uniformAddress)
	}
}
