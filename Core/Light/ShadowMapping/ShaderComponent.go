package ShadowMapping

import (
	"fmt"
	core "github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

const (
	ua_shadowMap = "u_shadowMap"
)

type ShaderComponent struct {
	*Shader.ShaderProgramCore
}

func (program *ShaderComponent) BindObject(i interface{}) error {
	switch v := i.(type) {
	case core.IFrameBuffer:
		var errors Error.ErrorCollection
		for _, texture := range v.GetTextures() {
			errors.Push(program.BindObject(texture))
		}
		return errors.Err()
	case Texture.ITexture:
		return program.BindUniform(v, ua_shadowMap)
	default:
		return fmt.Errorf("shadow shader does not support type %T", v)
	}
}
