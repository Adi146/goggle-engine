package BasicShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
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

func (program *BasicShaderProgram) BindObject(i interface{}) []error {
	switch v := i.(type) {
	case *Model.Model:
		return program.bindModel(v)
	case Camera.ICamera:
		return program.bindCamera(v)
	default:
		return []error{fmt.Errorf("type %T not supported", v)}
	}
}

func (program *BasicShaderProgram) bindCamera(camera Camera.ICamera) []error {
	var errors []error
	if err := program.BindUniform(camera.GetProjectionMatrix(), projectionMatrix_uniformAddress); err != nil {
		errors = append(errors, err)
	}
	if err := program.BindUniform(camera.GetViewMatrix(), viewMatrix_uniformAddress); err != nil {
		errors = append(errors, err)
	}
	return errors
}

func (program *BasicShaderProgram) bindModel(model *Model.Model) []error {
	var errors []error
	if err := program.BindUniform(model.ModelMatrix, modelMatrix_uniformAddress); err != nil {
		errors = append(errors, err)
	}
	return errors
}
