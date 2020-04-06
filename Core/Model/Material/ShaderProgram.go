package Material

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

const (
	ua_material  = "u_material"
	ua_baseColor = ua_material + ".baseColor"
	ua_shininess = ua_material + ".shininess"

	ua_color_diffuse  = ua_baseColor + ".diffuse"
	ua_color_specular = ua_baseColor + ".specular"
	ua_color_emissive = ua_baseColor + ".emissive"

	ua_texture_diffuse  = ua_material + ".textureDiffuse"
	ua_texture_specular = ua_material + ".textureSpecular"
	ua_texture_emissive = ua_material + ".textureEmissive"
	ua_texture_normal   = ua_material + ".textureNormal"

	ua_has_texture_diffuse  = ua_material + ".hasTextureDiffuse"
	ua_has_texture_specular = ua_material + ".hasTextureSpecular"
	ua_has_texture_emissive = ua_material + ".hasTextureEmissive"
	ua_has_texture_normal   = ua_material + ".hasTextureNormal"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore

	BindFunctions []func(program *ShaderProgram, material *Material) error
}

func (program *ShaderProgram) GetUniformAddress(i interface{}) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Material:
		return program.bindMaterial(v)
	default:
		return fmt.Errorf("material shader does not support type %T", v)
	}
}

func (program *ShaderProgram) bindMaterial(material *Material) error {
	var err Error.ErrorCollection

	for _, bindFunction := range program.BindFunctions {
		err.Push(bindFunction(program, material))
	}

	return err.Err()
}

func BindDiffuse(program *ShaderProgram, material *Material) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.DiffuseBaseColor, ua_color_diffuse))

	exists := material.DiffuseTexture != nil
	if exists {
		err.Push(program.BindUniform(material.DiffuseTexture, ua_texture_diffuse))
	}
	err.Push(program.BindUniform(exists, ua_has_texture_diffuse))

	return err.Err()
}

func BindSpecular(program *ShaderProgram, material *Material) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.SpecularBaseColor, ua_color_specular))

	exists := material.SpecularTexture != nil
	if exists {
		err.Push(program.BindUniform(material.SpecularTexture, ua_texture_specular))
	}
	err.Push(program.BindUniform(exists, ua_has_texture_specular))

	return err.Err()
}

func BindEmissive(program *ShaderProgram, material *Material) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.EmissiveBaseColor, ua_color_emissive))

	exists := material.EmissiveTexture != nil
	if exists {
		err.Push(program.BindUniform(material.EmissiveTexture, ua_texture_emissive))
	}
	err.Push(program.BindUniform(exists, ua_has_texture_emissive))

	return err.Err()
}

func BindNormals(program *ShaderProgram, material *Material) error {
	var err Error.ErrorCollection

	exists := material.NormalTexture != nil
	if exists {
		err.Push(program.BindUniform(material.NormalTexture, ua_texture_normal))
	}
	err.Push(program.BindUniform(exists, ua_has_texture_normal))

	return err.Err()
}

func BindShininess(program *ShaderProgram, material *Material) error {
	return program.BindUniform(material.Shininess, ua_shininess)
}
