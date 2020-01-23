package Model

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

type Material struct {
	DiffuseBaseColor  Vector.Vector3
	SpecularBaseColor Vector.Vector3
	EmissiveBaseColor Vector.Vector3

	Shininess float32

	Textures []*Texture
}
