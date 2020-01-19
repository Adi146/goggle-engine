package PhongShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
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

func (program *PhongShaderProgram) ResetIndexCounter() {
	program.pointLightIndex = 0
	program.spotLightIndex = 0
}

func (program *PhongShaderProgram) BindMaterial(material *Model.Material) error {
	if err := program.BindVector3(&material.DiffuseBaseColor, material_diffuseBase_uniformAddress); err != nil {
		return err
	}
	if err := program.BindVector3(&material.SpecularBaseColor, material_specularBase_uniformAddress); err != nil {
		return err
	}
	if err := program.BindVector3(&material.EmissiveBaseColor, material_emissiveBase_uniformAddress); err != nil {
		return err
	}
	if err := program.BindFloat(material.Shininess, material_shineness_uniformAddress); err != nil {
		return err
	}

	for i, texture := range material.DiffuseTextures {
		if err := program.BindTexture(uint32(i), texture, fmt.Sprintf(texture_diffuse_unifromAddress, i)); err != nil {
			return err
		}
	}

	if err := program.BindInt(int32(len(material.DiffuseTextures)), texture_numDiffuse_uniformAddress); err != nil {
		return err
	}

	/*if err := program.BindTexture(0, material.NormalTextures[0], texture_normals_unifromAddress); err != nil{
		return err
	}*/
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

func (program *PhongShaderProgram) BindModel(model *Model.Model) error {
	if err := program.BindMatrix(model.ModelMatrix, modelMatrix_uniformAddress); err != nil {
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

func (program *PhongShaderProgram) BindSpotLight(light *Light.SpotLight) error {
	if err := program.BindVector3(&light.Position, fmt.Sprintf(spotLight_position_uniformAddress, program.spotLightIndex)); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Direction, fmt.Sprintf(spotLight_direction_uniformAddress, program.spotLightIndex)); err != nil {
		return err
	}
	if err := program.BindFloat(light.InnerCone, fmt.Sprintf(spotLight_innerCone_uniformAddress, program.spotLightIndex)); err != nil {
		return err
	}
	if err := program.BindFloat(light.OuterCone, fmt.Sprintf(spotLight_outerCone_uniformAddress, program.spotLightIndex)); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Ambient, fmt.Sprintf(spotLight_ambient_uniformAddress, program.spotLightIndex)); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Diffuse, fmt.Sprintf(spotLight_diffuse_uniformAddress, program.spotLightIndex)); err != nil {
		return err
	}
	if err := program.BindVector3(&light.Specular, fmt.Sprintf(spotLight_specular_uniformAddress, program.spotLightIndex)); err != nil {
		return err
	}
	if err := program.BindFloat(light.Linear, fmt.Sprintf(spotLight_linear_uniformAddress, program.spotLightIndex)); err != nil {
		return err
	}
	if err := program.BindFloat(light.Quadratic, fmt.Sprintf(spotLight_quadratic_uniformAddress, program.spotLightIndex)); err != nil {
		return err
	}

	program.spotLightIndex++
	program.BindInt(program.spotLightIndex, numSpotLights_uniformAddress)

	return nil
}
