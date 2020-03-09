package Material

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

type Material struct {
	DiffuseBaseColor  GeometryMath.Vector4
	SpecularBaseColor GeometryMath.Vector3
	EmissiveBaseColor GeometryMath.Vector3

	Shininess float32

	Textures []*Texture.Texture
}

func (material *Material) Bind() {
	for _, texture := range material.Textures {
		texture.Bind()
	}
}

func (material *Material) Unbind() {
	for _, texture := range material.Textures {
		texture.Unbind()
	}
}
