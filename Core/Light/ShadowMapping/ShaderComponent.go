package ShadowMapping

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	ua_shadowMapDirectionalLight = "u_shadowMapDirectionalLights[%d]"
	ua_shadowMapsPointLight      = "u_shadowMapsPointLights[%d]"
	ua_shadowMapsSpotLight       = "u_shadowMapsSpotLights[%d]"
)

type ShaderComponent struct {
	*Shader.ShaderProgramCore
}

func (program *ShaderComponent) GetUniformAddress(i interface{}) (string, error) {
	switch v := i.(type) {
	case Texture.ITexture:
		switch t := v.GetType(); t {
		case ShadowMapDirectionalLight:
			return ua_shadowMapDirectionalLight, nil
		case ShadowMapPointLight:
			return ua_shadowMapsPointLight, nil
		case ShadowMapSpotLight:
			return ua_shadowMapsSpotLight, nil
		default:
			return "", fmt.Errorf("shadow shader does not support textures of type %s", t)
		}
	default:
		return "", fmt.Errorf("shadow shader does not support type %T", v)
	}
}

func (program *ShaderComponent) BindObject(i interface{}) error {
	return fmt.Errorf("use GetUniformAddress and specify index instead")
}
