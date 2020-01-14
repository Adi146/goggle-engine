package PhongShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Geometry"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

const (
	material_diffuse_uniformAddress   = "u_material.diffuse"
	material_specular_uniformAddress  = "u_material.specular"
	material_emissive_uniformAddress  = "u_material.emissive"
	material_shineness_uniformAddress = "u_material.shininess"

	directionalLight_direction_uniformAddress = "u_directionalLight.direction"
	directionalLight_ambient_uniformAddress   = "u_directionalLight.ambient"
	directionalLight_diffuse_uniformAddress   = "u_directionalLight.diffuse"
	directionalLight_specular_uniformAddress  = "u_directionalLight.specular"

	pointLight_position_uniformAddress  = "u_pointLights[%d].position"
	pointLight_ambient_uniformAddress   = "u_pointLights[%d].ambient"
	pointLight_diffuse_uniformAddress   = "u_pointLights[%d].diffuse"
	pointLight_specular_uniformAddress  = "u_pointLights[%d].specular"
	pointLight_linear_uniformAddress    = "u_pointLights[%d].linear"
	pointLight_quadratic_uniformAddress = "u_pointLights[%d].quadratic"
	numPointLights_uniformAddress       = "u_numPointLights"

	modelMatrix_uniformAddress = "u_modelMatrix"

	viewMatrix_uniformAddress       = "u_viewMatrix"
	projectionMatrix_uniformAddress = "u_projectionMatrix"
)

type PhongShaderProgram struct {
	*Shader.ShaderProgramCore

	pointLightIndex int32
}

func NewPhongShaderProgram(vertexShaderFile string, fragmentShaderFile string) (*PhongShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFile, fragmentShaderFile)
	if err != nil {
		return nil, err
	}

	return &PhongShaderProgram{
		ShaderProgramCore: shaderCore,
		pointLightIndex:   0,
	}, nil
}

func NewPhongIShaderProgram(vertexShaderFile string, fragmentShaderFile string) (Shader.IShaderProgram, error) {
	return NewPhongShaderProgram(vertexShaderFile, fragmentShaderFile)
}

func (program *PhongShaderProgram) ResetIndexCounter() {
	program.pointLightIndex = 0
}

func (program *PhongShaderProgram) BindMaterial(material *Model.Material) error {
	if err := program.BindVector3(&material.Diffuse, material_diffuse_uniformAddress); err != nil {
		return err
	}
	if err := program.BindVector3(&material.Specular, material_specular_uniformAddress); err != nil {
		return err
	}
	if err := program.BindVector3(&material.Emissive, material_emissive_uniformAddress); err != nil {
		return err
	}
	if err := program.BindFloat(material.Shininess, material_shineness_uniformAddress); err != nil {
		return err
	}
	return nil
}

func (program *PhongShaderProgram) BindCamera(camera Camera.ICamera) error {
	if err := program.BindMatrix(camera.GetProjectionMatrix(), projectionMatrix_uniformAddress); err != nil {
		return err
	}
	if err := program.BindMatrix(camera.GetViewMatrix(), viewMatrix_uniformAddress); err != nil {
		return err
	}
	return nil
}

func (program *PhongShaderProgram) BindGeometry(geometry *Geometry.Geometry) error {
	if err := program.BindMatrix(geometry.ModelMatrix, modelMatrix_uniformAddress); err != nil {
		return err
	}
	return nil
}

func (program *PhongShaderProgram) BindDirectionalLight(light *Light.DirectionalLight) error {
	if err := program.BindVector3(&light.Direction, directionalLight_direction_uniformAddress); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Ambient, directionalLight_ambient_uniformAddress); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Diffuse, directionalLight_diffuse_uniformAddress); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Specular, directionalLight_specular_uniformAddress); err != nil {
		return err
	}
	return nil
}

func (program *PhongShaderProgram) BindPointLight(light *Light.PointLight) error {
	if err := program.BindVector3(&light.Position, fmt.Sprintf(pointLight_position_uniformAddress, program.pointLightIndex)); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Ambient, fmt.Sprintf(pointLight_ambient_uniformAddress, program.pointLightIndex)); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Diffuse, fmt.Sprintf(pointLight_diffuse_uniformAddress, program.pointLightIndex)); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Specular, fmt.Sprintf(pointLight_specular_uniformAddress, program.pointLightIndex)); err != nil {
		return err
	}
	if err := program.BindFloat(light.Linear, fmt.Sprintf(pointLight_linear_uniformAddress, program.pointLightIndex)); err != nil {
		return err
	}
	if err := program.BindFloat(light.Quadratic, fmt.Sprintf(pointLight_quadratic_uniformAddress, program.pointLightIndex)); err != nil {
		return err
	}

	program.pointLightIndex++
	program.BindInt(program.pointLightIndex, numPointLights_uniformAddress)

	return nil
}
