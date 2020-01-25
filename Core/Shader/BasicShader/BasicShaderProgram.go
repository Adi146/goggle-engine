package BasicShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
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

func (program *BasicShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Model.Model:
		return program.bindModel(v)
	case Camera.ICamera:
		return program.bindCamera(v)
	default:
		return fmt.Errorf("type %T not supported", v)
	}
}

func (program *BasicShaderProgram) bindCamera(camera Camera.ICamera) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(camera.GetProjectionMatrix(), projectionMatrix_uniformAddress))
	err.Push(program.BindUniform(camera.GetViewMatrix(), viewMatrix_uniformAddress))

	return err.Err()
}

func (program *BasicShaderProgram) bindModel(model *Model.Model) error {
	return program.BindUniform(model.ModelMatrix, modelMatrix_uniformAddress)
}
