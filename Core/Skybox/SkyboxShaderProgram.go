package Skybox

import (
	"fmt"

	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	skybox_uniformAddress = "u_skybox"

	cameraUBO_uniformAddress = "Camera"
)

type SkyboxShaderProgram struct {
	*Shader.ShaderProgramCore
}

func NewSkyboxShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (*SkyboxShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &SkyboxShaderProgram{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewSkyboxIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewSkyboxShaderProgram(vertexShaderFiles, fragmentShaderFiles)
}

func (program *SkyboxShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Texture.CubeMap:
		return program.bindSkybox(v)
	case *Camera.UniformBuffer:
		return program.BindUniform(v, cameraUBO_uniformAddress)
	default:
		return fmt.Errorf("type %T not supported", v)
	}
}

func (program *SkyboxShaderProgram) bindSkybox(skybox *Texture.CubeMap) error {
	return program.BindTexture(0, skybox, skybox_uniformAddress)
}
