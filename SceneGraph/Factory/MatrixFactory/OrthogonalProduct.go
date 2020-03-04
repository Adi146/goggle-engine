package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"reflect"
)

const (
	OrthogonalFactoryName = "orthogonal"
)

func init() {
	AddType(OrthogonalFactoryName, reflect.TypeOf((*OrthogonalProduct)(nil)).Elem())
}

type OrthogonalProduct struct {
	Left   float32 `yaml:"left"`
	Right  float32 `yaml:"right"`
	Bottom float32 `yaml:"bottom"`
	Top    float32 `yaml:"top"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (product *OrthogonalProduct) Decode() *Matrix.Matrix4x4 {
	return Matrix.Orthogonal(
		product.Left,
		product.Right,
		product.Bottom,
		product.Top,
		product.Near,
		product.Far,
	)
}
