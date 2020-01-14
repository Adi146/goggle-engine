package BasicShader

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Geometry"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

const (
	modelMatrix_uniformAddress = "u_modelMatrix"

	viewMatrix_uniformAddress       = "u_viewMatrix"
	projectionMatrix_uniformAddress = "u_projectionMatrix"
)

type BasicShaderProgram struct {
	*Shader.ShaderProgramCore
}

func NewBasicShaderProgram(vertexShaderFile string, fragmentShaderFile string) (*BasicShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFile, fragmentShaderFile)
	if err != nil {
		return nil, err
	}

	return &BasicShaderProgram{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewBasicIShaderProgram(vertexShaderFile string, fragmentShaderFile string) (Shader.IShaderProgram, error) {
	return NewBasicShaderProgram(vertexShaderFile, fragmentShaderFile)
}

func (program *BasicShaderProgram) ResetIndexCounter() {
}

func (program *BasicShaderProgram) BindMaterial(material *Model.Material) error {
	return nil
}

func (program *BasicShaderProgram) BindCamera(camera Camera.ICamera) error {
	if err := program.BindMatrix(camera.GetProjectionMatrix(), projectionMatrix_uniformAddress); err != nil {
		return err
	}
	if err := program.BindMatrix(camera.GetViewMatrix(), viewMatrix_uniformAddress); err != nil {
		return err
	}
	return nil
}

func (program *BasicShaderProgram) BindGeometry(geometry *Geometry.Geometry) error {
	if err := program.BindMatrix(geometry.ModelMatrix, modelMatrix_uniformAddress); err != nil {
		return err
	}
	return nil
}

func (program *BasicShaderProgram) BindDirectionalLight(light *Light.DirectionalLight) error {
	return nil
}

func (program *BasicShaderProgram) BindPointLight(light *Light.PointLight) error {
	return nil
}

func (program *BasicShaderProgram) BindSpotLight(light *Light.SpotLight) error {
	return nil
}
