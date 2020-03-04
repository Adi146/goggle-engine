package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"reflect"
)

const (
	PerspectiveFactoryName = "perspective"
)

func init() {
	AddType(PerspectiveFactoryName, reflect.TypeOf((*PerspectiveProduct)(nil)).Elem())
}

type PerspectiveProduct struct {
	Fovy   float32 `yaml:"fovy"`
	Aspect float32 `yaml:"aspect"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (product *PerspectiveProduct) Decode() *Matrix.Matrix4x4 {
	return Matrix.Perspective(
		Angle.Radians(product.Fovy)/2,
		product.Aspect,
		product.Near,
		product.Far,
	)
}
