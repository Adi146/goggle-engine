package PhongShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Light/DirectionalLight"
	"github.com/Adi146/goggle-engine/Core/Light/PointLight"

	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

const (
	material_diffuseBase_uniformAddress  = "u_material.baseColor.diffuse"
	material_specularBase_uniformAddress = "u_material.baseColor.specular"
	material_emissiveBase_uniformAddress = "u_material.baseColor.emissive"
	material_shineness_uniformAddress    = "u_material.shininess"

	texture_diffuse_unifromAddress     = "u_material.texturesDiffuse[%d]"
	texture_specular_uniformAddress    = "u_material.texturesSpecular[%d]"
	texture_emissive_uniformAddress    = "u_material.texturesEmissive[%d]"
	texture_normals_unifromAddress     = "u_material.texturesNormals[%d]"
	texture_numDiffuse_uniformAddress  = "u_material.numTextureDiffuse"
	texture_numSpecular_uniformAddress = "u_material.numTextureSpecular"
	texture_numEmissive_uniformAddress = "u_material.numTextureEmissive"
	texture_numNormals_uniformAddress  = "u_material.numTextureNormals"

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

	cameraUBO_uniformAddress           = "Camera"
	directionalLightUBO_uniformAddress = "directionalLight"
	pointLightUBO_uniformAddrress      = "pointLight"
)

var (
	textureUniformMap = map[Texture.TextureType]string{
		Texture.DiffuseTexture:  texture_diffuse_unifromAddress,
		Texture.SpecularTexture: texture_specular_uniformAddress,
		Texture.EmissiveTexture: texture_emissive_uniformAddress,
		Texture.NormalsTexture:  texture_normals_unifromAddress,
	}
	numTextureUniformMap = map[Texture.TextureType]string{
		Texture.DiffuseTexture:  texture_numDiffuse_uniformAddress,
		Texture.SpecularTexture: texture_numSpecular_uniformAddress,
		Texture.EmissiveTexture: texture_numEmissive_uniformAddress,
		Texture.NormalsTexture:  texture_numNormals_uniformAddress,
	}
)

type PhongShaderProgram struct {
	*Shader.ShaderProgramCore

	spotLightIndex  int32
}

func NewPhongShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (*PhongShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &PhongShaderProgram{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewPhongIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewPhongShaderProgram(vertexShaderFiles, fragmentShaderFiles)
}

func (program *PhongShaderProgram) BeginDraw() error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(program.spotLightIndex, numSpotLights_uniformAddress))

	return err.Err()
}

func (program *PhongShaderProgram) EndDraw() {
	program.spotLightIndex = 0
}

func (program *PhongShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Model.Material:
		return program.bindMaterial(v)
	case *Model.Model:
		return program.bindModel(v)
	case *DirectionalLight.UniformBuffer:
		return program.BindUniform(v, directionalLightUBO_uniformAddress)
	case *PointLight.UniformBuffer:
		return program.BindUniform(v, pointLightUBO_uniformAddrress)
	case *Light.SpotLight:
		return program.bindSpotLight(v)
	case *Camera.UniformBuffer:
		return program.BindUniform(v, cameraUBO_uniformAddress)
	default:
		return fmt.Errorf("type %T not supported", v)
	}
}

func (program *PhongShaderProgram) bindMaterial(material *Model.Material) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.DiffuseBaseColor, material_diffuseBase_uniformAddress))
	err.Push(program.BindUniform(&material.SpecularBaseColor, material_specularBase_uniformAddress))
	err.Push(program.BindUniform(&material.EmissiveBaseColor, material_emissiveBase_uniformAddress))
	err.Push(program.BindUniform(material.Shininess, material_shineness_uniformAddress))

	textureIndexMap := make(map[Texture.TextureType]int)
	for i, texture := range material.Textures {
		err.Push(program.BindTexture(uint32(i), texture, fmt.Sprintf(textureUniformMap[texture.TextureType], textureIndexMap[texture.TextureType])))
		textureIndexMap[texture.TextureType] += 1
	}
	for textureType, uniformAddress := range numTextureUniformMap {
		err.Push(program.BindUniform(int32(textureIndexMap[textureType]), uniformAddress))
	}

	return err.Err()
}

func (program *PhongShaderProgram) bindModel(model *Model.Model) error {
	return program.BindUniform(model.ModelMatrix, modelMatrix_uniformAddress)
}

func (program *PhongShaderProgram) bindSpotLight(light *Light.SpotLight) error {
	var err Error.ErrorCollection

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
