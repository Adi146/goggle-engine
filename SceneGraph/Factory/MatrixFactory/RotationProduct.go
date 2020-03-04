package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"reflect"
)

const (
	RotationFactoryName = "rotation"
)

func init() {
	AddType(RotationFactoryName, reflect.TypeOf((*RotationProduct)(nil)).Elem())
}

type RotationProduct struct {
	Vector *Vector.Vector3 `yaml:"axis"`
	Angle  float32         `yaml:"angle"`
}

func (product *RotationProduct) Decode() *Matrix.Matrix4x4 {
	return Matrix.Rotate(Angle.Radians(product.Angle), product.Vector)
}
