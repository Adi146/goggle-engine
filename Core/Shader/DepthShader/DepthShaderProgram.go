package DepthShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

const (
	lightSpaceMatrix_uniformAddress = "u_lightSpaceMatrix"
	modelMatrix_uniformAddress      = "u_modelMatrix"
)

type DepthShaderProgram struct {
	*Shader.ShaderProgramCore
}

func NewDepthShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (*DepthShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &DepthShaderProgram{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewDepthIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewDepthShaderProgram(vertexShaderFiles, fragmentShaderFiles)
}

func (program *DepthShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Model.Model:
		return program.bindModel(v)
	case *Light.DirectionalLight:
		return program.bindDirectionalLight(v)
	default:
		return fmt.Errorf("type %T not supported", v)
	}
}

func (program *DepthShaderProgram) bindDirectionalLight(light *Light.DirectionalLight) error {
	nearPlane := float32(1.0)
	farPlane := float32(7.5)

	lightProjection := Matrix.Orthogonal(-10.0, 10, -10, 10, nearPlane, farPlane)
	lightView := Matrix.LookAt(&Vector.Vector3{-2.0, 4.0, -1.0}, &Vector.Vector3{0.0, 0.0, 0.0}, &Vector.Vector3{0.0, 1.0, 0.0})

	lightSpaceMatrix := lightView.Mul(lightProjection)

	program.BindUniform(lightSpaceMatrix, lightSpaceMatrix_uniformAddress)
}

func (program *DepthShaderProgram) bindModel(model *Model.Model) error {
	return program.BindUniform(model.ModelMatrix, modelMatrix_uniformAddress)
}
