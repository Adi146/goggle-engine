package Model

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

type Material struct {
	DiffuseBaseColor  Vector.Vector4
	SpecularBaseColor Vector.Vector3
	EmissiveBaseColor Vector.Vector3

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
