package Material

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

const (
	ua_baseColor = ".baseColor"
	ua_shininess = ".shininess"
	ua_uvScale   = ".uvScale"

	ua_color_diffuse  = ua_baseColor + ".diffuse"
	ua_color_specular = ua_baseColor + ".specular"
	ua_color_emissive = ua_baseColor + ".emissive"

	ua_texture_diffuse  = ".textureDiffuse"
	ua_texture_specular = ".textureSpecular"
	ua_texture_emissive = ".textureEmissive"
	ua_texture_normal   = ".textureNormal"

	ua_has_texture_diffuse  = ".hasTextureDiffuse"
	ua_has_texture_specular = ".hasTextureSpecular"
	ua_has_texture_emissive = ".hasTextureEmissive"
	ua_has_texture_normal   = ".hasTextureNormal"

	ua_materials    = "u_materials[%d]"
	ua_blendMap     = "u_blendMap"
	ua_has_blendMap = "u_hasBlendMap"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore

	BindFunctions []func(program *ShaderProgram, material *Material, materialUniformAddress string) error
}

func (program *ShaderProgram) GetUniformAddress(i interface{}) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Material:
		return program.bindSingleMaterial(v)
	case *BlendMaterial:
		return program.bindBlendMaterial(v)
	default:
		return fmt.Errorf("material shader does not support type %T", v)
	}
}

func (program *ShaderProgram) bindSingleMaterial(material *Material) error {
	var err Error.ErrorCollection

	err.Push(program.bindMaterial(material, 0))
	err.Push(program.BindUniform(false, ua_has_blendMap))

	return err.Err()
}

func (program *ShaderProgram) bindMaterial(material *Material, materialIndex int) error {
	var err Error.ErrorCollection
	uaMaterial := fmt.Sprintf(ua_materials, materialIndex)

	for _, bindFunction := range program.BindFunctions {
		err.Push(bindFunction(program, material, uaMaterial))
	}
	err.Push(program.BindUniform(material.UvScale, uaMaterial+ua_uvScale))

	return err.Err()
}

func (program *ShaderProgram) bindBlendMaterial(blendMap *BlendMaterial) error {
	var err Error.ErrorCollection

	for i, material := range blendMap.Materials {
		err.Push(program.bindMaterial(&material, i))
	}
	err.Push(program.BindUniform(&blendMap.BlendTexture, ua_blendMap))
	err.Push(program.BindUniform(true, ua_has_blendMap))

	return err.Err()
}

func BindDiffuse(program *ShaderProgram, material *Material, uaMaterial string) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.DiffuseBaseColor, uaMaterial+ua_color_diffuse))

	exists := material.Textures.Diffuse != nil
	if exists {
		err.Push(program.BindUniform(material.Textures.Diffuse, uaMaterial+ua_texture_diffuse))
	}
	err.Push(program.BindUniform(exists, uaMaterial+ua_has_texture_diffuse))

	return err.Err()
}

func BindSpecular(program *ShaderProgram, material *Material, uaMaterial string) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.SpecularBaseColor, uaMaterial+ua_color_specular))

	exists := material.Textures.Specular != nil
	if exists {
		err.Push(program.BindUniform(material.Textures.Specular, uaMaterial+ua_texture_specular))
	}
	err.Push(program.BindUniform(exists, uaMaterial+ua_has_texture_specular))

	return err.Err()
}

func BindEmissive(program *ShaderProgram, material *Material, uaMaterial string) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(&material.EmissiveBaseColor, uaMaterial+ua_color_emissive))

	exists := material.Textures.Emissive != nil
	if exists {
		err.Push(program.BindUniform(material.Textures.Emissive, uaMaterial+ua_texture_emissive))
	}
	err.Push(program.BindUniform(exists, uaMaterial+ua_has_texture_emissive))

	return err.Err()
}

func BindNormals(program *ShaderProgram, material *Material, uaMaterial string) error {
	var err Error.ErrorCollection

	exists := material.Textures.Normal != nil
	if exists {
		err.Push(program.BindUniform(material.Textures.Normal, uaMaterial+ua_texture_normal))
	}
	err.Push(program.BindUniform(exists, uaMaterial+ua_has_texture_normal))

	return err.Err()
}

func BindShininess(program *ShaderProgram, material *Material, uaMaterial string) error {
	return program.BindUniform(material.Shininess, uaMaterial+ua_shininess)
}
