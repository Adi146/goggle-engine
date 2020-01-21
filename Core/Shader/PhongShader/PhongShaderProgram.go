package PhongShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils"
)

const (
	material_diffuseBase_uniformAddress  = "u_material.baseColor.diffuse"
	material_specularBase_uniformAddress = "u_material.baseColor.specular"
	material_emissiveBase_uniformAddress = "u_material.baseColor.emissive"
	material_shineness_uniformAddress    = "u_material.shininess"

	texture_diffuse_unifromAddress    = "u_material.texturesDiffuse[%d]"
	texture_normals_unifromAddress    = "u_material.texturesNormals[%d]"
	texture_numDiffuse_uniformAddress = "u_material.numTextureDiffuse"
	texture_numNormals_uniformAddress = "u_material.numTextureNormals"

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

	spotLight_position_uniformAddress  = "u_spotLights[%d].position"
	spotLight_direction_uniformAddress = "u_spotLights[%d].direction"
	spotLight_innerCone_uniformAddress = "u_spotLights[%d].innerCone"
	spotLight_outerCone_uniformAddress = "u_spotLights[%d].outerCone"
	spotLight_ambient_uniformAddress   = "u_spotLights[%d].ambient"
	spotLight_diffuse_uniformAddress   = "u_spotLights[%d].diffuse"
	spotLight_specular_uniformAddress  = "u_spotLights[%d].specular"
	spotLight_linear_uniformAddress    = "u_spotLights[%d].linear"
	spotLight_quadratic_uniformAddress = "u_spotLights[%d].quadratic"
	numSpotLights_uniformAddress       = "u_numSpotLights"

	modelMatrix_uniformAddress = "u_modelMatrix"

	viewMatrix_uniformAddress       = "u_viewMatrix"
	projectionMatrix_uniformAddress = "u_projectionMatrix"
)

type PhongShaderProgram struct {
	*Shader.ShaderProgramCore

	pointLightIndex int32
	spotLightIndex  int32
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

func (program *PhongShaderProgram) BeginDraw() error {
	var err Utils.ErrorCollection

	err.Push(program.BindUniform(program.pointLightIndex, numPointLights_uniformAddress))
	err.Push(program.BindUniform(program.spotLightIndex, numSpotLights_uniformAddress))

	return err.Err()
}

func (program *PhongShaderProgram) EndDraw() {
	program.pointLightIndex = 0
	program.spotLightIndex = 0
}

func (program *PhongShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Model.Material:
		return program.bindMaterial(v)
	case *Model.Model:
		return program.bindModel(v)
	case Camera.ICamera:
		return program.bindCamera(v)
	case *Light.DirectionalLight:
		return program.bindDirectionalLight(v)
	case *Light.PointLight:
		return program.bindPointLight(v)
	case *Light.SpotLight:
		return program.bindSpotLight(v)
	default:
		return fmt.Errorf("type %T not supported", v)
	}
}

func (program *PhongShaderProgram) bindMaterial(material *Model.Material) error {
	var err Utils.ErrorCollection

	err.Push(program.BindUniform(&material.DiffuseBaseColor, material_diffuseBase_uniformAddress))
	err.Push(program.BindUniform(&material.SpecularBaseColor, material_specularBase_uniformAddress))
	err.Push(program.BindUniform(&material.EmissiveBaseColor, material_emissiveBase_uniformAddress))
	err.Push(program.BindUniform(material.Shininess, material_shineness_uniformAddress))

	textureIndex := 0
	for i, texture := range material.DiffuseTextures {
		err.Push(program.BindTexture(uint32(textureIndex), texture, fmt.Sprintf(texture_diffuse_unifromAddress, i)))
		textureIndex++
	}
	err.Push(program.BindUniform(int32(len(material.DiffuseTextures)), texture_numDiffuse_uniformAddress))
	for i, texture := range material.NormalTextures {
		err.Push(program.BindTexture(uint32(textureIndex), texture, fmt.Sprintf(texture_normals_unifromAddress, i)))
		textureIndex++
	}
	err.Push(program.BindUniform(int32(len(material.NormalTextures)), texture_numNormals_uniformAddress))

	return err.Err()
}

func (program *PhongShaderProgram) bindCamera(camera Camera.ICamera) error {
	var err Utils.ErrorCollection

	err.Push(program.BindUniform(camera.GetProjectionMatrix(), projectionMatrix_uniformAddress))
	err.Push(program.BindUniform(camera.GetViewMatrix(), viewMatrix_uniformAddress))

	return err.Err()
}

func (program *PhongShaderProgram) bindModel(model *Model.Model) error {
	return program.BindUniform(model.ModelMatrix, modelMatrix_uniformAddress)
}

func (program *PhongShaderProgram) bindDirectionalLight(light *Light.DirectionalLight) error {
	var err Utils.ErrorCollection

	err.Push(program.BindUniform(&light.Direction, directionalLight_direction_uniformAddress))
	err.Push(program.BindUniform(&light.Ambient, directionalLight_ambient_uniformAddress))
	err.Push(program.BindUniform(&light.Diffuse, directionalLight_diffuse_uniformAddress))
	err.Push(program.BindUniform(&light.Specular, directionalLight_specular_uniformAddress))

	return err.Err()
}

func (program *PhongShaderProgram) bindPointLight(light *Light.PointLight) error {
	var err Utils.ErrorCollection

	err.Push(program.BindUniform(&light.Position, fmt.Sprintf(pointLight_position_uniformAddress, program.pointLightIndex)))
	err.Push(program.BindUniform(&light.Ambient, fmt.Sprintf(pointLight_ambient_uniformAddress, program.pointLightIndex)))
	err.Push(program.BindUniform(&light.Diffuse, fmt.Sprintf(pointLight_diffuse_uniformAddress, program.pointLightIndex)))
	err.Push(program.BindUniform(&light.Specular, fmt.Sprintf(pointLight_specular_uniformAddress, program.pointLightIndex)))
	err.Push(program.BindUniform(light.Linear, fmt.Sprintf(pointLight_linear_uniformAddress, program.pointLightIndex)))
	err.Push(program.BindUniform(light.Quadratic, fmt.Sprintf(pointLight_quadratic_uniformAddress, program.pointLightIndex)))

	program.pointLightIndex++
	return err.Err()
}

func (program *PhongShaderProgram) bindSpotLight(light *Light.SpotLight) error {
	var err Utils.ErrorCollection

	err.Push(program.BindUniform(&light.Position, fmt.Sprintf(spotLight_position_uniformAddress, program.spotLightIndex)))
	err.Push(program.BindUniform(&light.Direction, fmt.Sprintf(spotLight_direction_uniformAddress, program.spotLightIndex)))
	err.Push(program.BindUniform(light.InnerCone, fmt.Sprintf(spotLight_innerCone_uniformAddress, program.spotLightIndex)))
	err.Push(program.BindUniform(light.OuterCone, fmt.Sprintf(spotLight_outerCone_uniformAddress, program.spotLightIndex)))
	err.Push(program.BindUniform(&light.Ambient, fmt.Sprintf(spotLight_ambient_uniformAddress, program.spotLightIndex)))
	err.Push(program.BindUniform(&light.Diffuse, fmt.Sprintf(spotLight_diffuse_uniformAddress, program.spotLightIndex)))
	err.Push(program.BindUniform(&light.Specular, fmt.Sprintf(spotLight_specular_uniformAddress, program.spotLightIndex)))
	err.Push(program.BindUniform(light.Linear, fmt.Sprintf(spotLight_linear_uniformAddress, program.spotLightIndex)))
	err.Push(program.BindUniform(light.Quadratic, fmt.Sprintf(spotLight_quadratic_uniformAddress, program.spotLightIndex)))

	program.spotLightIndex++
	return err.Err()
}
