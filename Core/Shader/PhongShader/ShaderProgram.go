package PhongShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light/DirectionalLight"
	"github.com/Adi146/goggle-engine/Core/Light/PointLight"
	"github.com/Adi146/goggle-engine/Core/Light/SpotLight"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shadow"
	"github.com/Adi146/goggle-engine/Core/Shadow/ShadowMapShader"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
)

const (
	ua_modelMatrix = "u_modelMatrix"

	ua_camera           = "camera"
	ua_directionalLight = "directionalLight"
	ua_pointLight       = "pointLight"
	ua_spotLight        = "spotLight"
)

func init() {
	ShaderFactory.AddType("phong", NewIShaderProgram)
}

type ShaderProgram struct {
	*Shader.ShaderProgramCore
	ShadowShader   Shadow.ShaderProgram
	MaterialShader Material.ShaderProgram
}

func NewShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (*ShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &ShaderProgram{
		ShaderProgramCore: shaderCore,
		ShadowShader: Shadow.ShaderProgram{
			ShaderProgramCore: shaderCore,
		},
		MaterialShader: Material.ShaderProgram{
			ShaderProgramCore: shaderCore,
			BindFunctions: []func(program *Material.ShaderProgram, material *Material.Material) error{
				Material.BindDiffuse,
				Material.BindSpecular,
				Material.BindEmissive,
				Material.BindNormals,
				Material.BindShininess,
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
	case *ShadowMapShader.ShadowMapBuffer:
		return program.ShadowShader.BindObject(v)
	case UniformBuffer.IUniformBuffer:
		switch t := v.GetType(); t {
		case DirectionalLight.UBO_type:
			return program.BindUniform(v, ua_directionalLight)
		case PointLight.UBO_type:
			return program.BindUniform(v, ua_pointLight)
		case SpotLight.UBO_type:
			return program.BindUniform(v, ua_spotLight)
		case Camera.UBO_type:
			return program.BindUniform(v, ua_camera)
		default:
			return fmt.Errorf("phong shader does not support uniform buffers of type %s", t)
		}
	default:
		return fmt.Errorf("phong shader does not support type %T", v)
	}
}
