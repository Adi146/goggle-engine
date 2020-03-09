package Material

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

const (
	ua_material  = "u_material"
	ua_baseColor = ua_material + ".baseColor"
	ua_shininess = ua_material + ".shininess"

	ua_color_diffuse  = ua_baseColor + ".diffuse"
	ua_color_specular = ua_baseColor + ".specular"
	ua_color_emissive = ua_baseColor + ".emissive"

	ua_textures_diffuse  = ua_material + ".texturesDiffuse[%d]"
	ua_textures_specular = ua_material + ".texturesSpecular[%d]"
	ua_textures_emissive = ua_material + ".texturesEmissive[%d]"
	ua_textures_normals  = ua_material + ".texturesNormals[%d]"

	ua_num_textures_diffuse  = ua_material + ".numTextureDiffuse"
	ua_num_textures_specular = ua_material + ".numTextureSpecular"
	ua_num_textures_emissive = ua_material + ".numTextureEmissive"
	ua_num_textures_normals  = ua_material + ".numTextureNormals"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore

	BindFunctions []func(program *ShaderProgram, material *Material) error
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
	var i int32
	for _, texture := range material.Textures {
		if texture.Type == Texture.DiffuseTexture {
			err.Push(program.BindUniform(texture, fmt.Sprintf(ua_textures_diffuse, i)))
			i++
		}
	}

	err.Push(program.BindUniform(i, ua_num_textures_diffuse))

	return err.Err()
}

func BindSpecular(program *ShaderProgram, material *Material) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.SpecularBaseColor, ua_color_specular))
	var i int32
	for _, texture := range material.Textures {
		if texture.Type == Texture.SpecularTexture {
			err.Push(program.BindUniform(texture, fmt.Sprintf(ua_textures_specular, i)))
			i++
		}
	}

	err.Push(program.BindUniform(i, ua_num_textures_specular))

	return err.Err()
}

func BindEmissive(program *ShaderProgram, material *Material) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.EmissiveBaseColor, ua_color_emissive))
	var i int32
	for _, texture := range material.Textures {
		if texture.Type == Texture.EmissiveTexture {
			err.Push(program.BindUniform(texture, fmt.Sprintf(ua_textures_emissive, i)))
			i++
		}
	}

	err.Push(program.BindUniform(i, ua_num_textures_emissive))

	return err.Err()
}

func BindNormals(program *ShaderProgram, material *Material) error {
	var err Error.ErrorCollection

	var i int32
	for _, texture := range material.Textures {
		if texture.Type == Texture.NormalsTexture {
			err.Push(program.BindUniform(texture, fmt.Sprintf(ua_textures_normals, i)))
			i++
		}
	}

	err.Push(program.BindUniform(i, ua_num_textures_normals))

	return err.Err()
}

func BindShininess(program *ShaderProgram, material *Material) error {
	return program.BindUniform(material.Shininess, ua_shininess)
}
