package ShadowMapping

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	ua_shadowMap = "u_shadowMap"
)

type ShaderComponent struct {
	*Shader.ShaderProgramCore
}

func (program *ShaderComponent) BindObject(i interface{}) error {
	switch v := i.(type) {
	case Texture.ITexture:
		return program.BindUniform(v, ua_shadowMap)
	default:
		return fmt.Errorf("shadow shader does not support type %T", v)
	}
}
