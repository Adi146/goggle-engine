package Skybox

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"

	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	shader_factory_name = "skyboxShader"

	ua_skybox = "u_skybox"

	ua_camera = "camera"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore
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
	}, nil
}

func NewIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string, geometryShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewShaderProgram(vertexShaderFiles, fragmentShaderFiles, geometryShaderFiles)
}

func (program *ShaderProgram) GetUniformAddress(i interface{}) (string, error) {
	switch v := i.(type) {
	case Texture.ITexture:
		return ua_skybox, nil
	case UniformBuffer.IUniformBuffer:
		switch t := v.GetType(); t {
		case Camera.UBO_type:
			return ua_camera, nil
		default:
			return "", fmt.Errorf("skybox shader does not support uniform buffers of type %s", t)
		}
	default:
		return "", fmt.Errorf("skybox shader does not support type %T", v)
	}
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case Mesh.VertexArray:
		v.Bind()
		return nil
	case *GeometryMath.Matrix4x4:
		return nil
	default:
		uniformAddress, err := program.GetUniformAddress(i)
		if err != nil {
			return err
		}
		return program.BindUniform(i, uniformAddress)
	}
}
