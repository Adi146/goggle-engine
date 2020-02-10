package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type IDrawable interface {
	Draw(shader Shader.IShaderProgram) error
	GetPosition() *Vector.Vector3
}
