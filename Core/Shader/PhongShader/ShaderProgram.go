package PhongShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light/DirectionalLight"
	"github.com/Adi146/goggle-engine/Core/Light/PointLight"
	"github.com/Adi146/goggle-engine/Core/Light/SpotLight"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shadow"
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

	modelMatrix_uniformAddress       = "u_modelMatrix"
	texture_shadowMap_uniformAddress = "u_shadowMap"

	cameraUBO_uniformAddress           = "Camera"
	directionalLightUBO_uniformAddress = "directionalLight"
	pointLightUBO_uniformAddress       = "pointLight"
	spotLightUBO_uniformAddress        = "spotLight"
)

var (
	textureUniformMap = map[Texture.Type]string{
		Texture.DiffuseTexture:  texture_diffuse_unifromAddress,
		Texture.SpecularTexture: texture_specular_uniformAddress,
		Texture.EmissiveTexture: texture_emissive_uniformAddress,
		Texture.NormalsTexture:  texture_normals_unifromAddress,
	}
	numTextureUniformMap = map[Texture.Type]string{
		Texture.DiffuseTexture:  texture_numDiffuse_uniformAddress,
		Texture.SpecularTexture: texture_numSpecular_uniformAddress,
		Texture.EmissiveTexture: texture_numEmissive_uniformAddress,
		Texture.NormalsTexture:  texture_numNormals_uniformAddress,
	}
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore
}

func NewShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (*ShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &ShaderProgram{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewShaderProgram(vertexShaderFiles, fragmentShaderFiles)
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Model.Material:
		return program.bindMaterial(v)
	case *Model.Model:
		return program.bindModel(v)
	case *Shadow.ShadowMapBuffer:
		program.Bind()
		return program.BindUniform(&v.ShadowMap, texture_shadowMap_uniformAddress)
	case *DirectionalLight.UniformBuffer:
		return program.BindUniform(v, directionalLightUBO_uniformAddress)
	case *PointLight.UniformBuffer:
		return program.BindUniform(v, pointLightUBO_uniformAddress)
	case *SpotLight.UniformBuffer:
		return program.BindUniform(v, spotLightUBO_uniformAddress)
	case *Camera.UniformBuffer:
		return program.BindUniform(v, cameraUBO_uniformAddress)
	default:
		return fmt.Errorf("phong shader does not support type %T", v)
	}
}

func (program *ShaderProgram) bindMaterial(material *Model.Material) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.DiffuseBaseColor, material_diffuseBase_uniformAddress))
	err.Push(program.BindUniform(&material.SpecularBaseColor, material_specularBase_uniformAddress))
	err.Push(program.BindUniform(&material.EmissiveBaseColor, material_emissiveBase_uniformAddress))
	err.Push(program.BindUniform(material.Shininess, material_shineness_uniformAddress))

	textureIndexMap := make(map[Texture.Type]int)
	for _, texture := range material.Textures {
		err.Push(program.BindUniform(texture, fmt.Sprintf(textureUniformMap[texture.Type], textureIndexMap[texture.Type])))
		textureIndexMap[texture.Type] += 1
	}
	for textureType, uniformAddress := range numTextureUniformMap {
		err.Push(program.BindUniform(int32(textureIndexMap[textureType]), uniformAddress))
	}

	return err.Err()
}

func (program *ShaderProgram) bindModel(model *Model.Model) error {
	return program.BindUniform(model.ModelMatrix, modelMatrix_uniformAddress)
}
