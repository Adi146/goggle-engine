package Shadow

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shadow/ShadowMapShader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

const (
	ua_shadowMap = "u_shadowMap"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *ShadowMapShader.ShadowMapBuffer:
		return program.BindObject(&v.ShadowMap)
	case Texture.ITexture:
		program.Bind()
		return program.BindUniform(v, ua_shadowMap)
	default:
		return fmt.Errorf("shadow shader does not support type %T", v)
	}
}