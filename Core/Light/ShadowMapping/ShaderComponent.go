package ShadowMapping

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	ua_shadowMapDirectionalLight = "u_shadowMapDirectionalLight"
	ua_shadowMapsPointLight      = "u_shadowMapsPointLight[%d]"
)

type ShaderComponent struct {
	*Shader.ShaderProgramCore
}

func (program *ShaderComponent) GetUniformAddress(i interface{}) (string, error) {
	switch v := i.(type) {
	case Texture.ITexture:
		switch t := v.GetType(); t {
		case internal.ShadowMapDirectionalLight:
			return ua_shadowMapDirectionalLight, nil
		case internal.ShadowMapPointLight:
			return ua_shadowMapsPointLight, nil
		default:
			return "", fmt.Errorf("shadow shader does not support textures of type %s", t)
		}
	default:
		return "", fmt.Errorf("shadow shader does not support type %T", v)
	}
}

func (program *ShaderComponent) BindObject(i interface{}) error {
	uniformAddress, err := program.GetUniformAddress(i)
	if err != nil {
		return err
	}

	switch v := i.(type) {
	case Texture.ITexture:
		switch t := v.GetType(); t {
		case internal.ShadowMapDirectionalLight:
			return program.BindUniform(v, uniformAddress)
		case internal.ShadowMapPointLight:
			return fmt.Errorf("use GetUniformAddress and specify index instead")
		}
	}

	return err
}
