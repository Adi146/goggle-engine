package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"reflect"
)

const (
	OrthogonalFactoryName = "orthogonal"
)

func init() {
	AddType(OrthogonalFactoryName, reflect.TypeOf((*OrthogonalConfig)(nil)).Elem())
}

type OrthogonalConfig struct {
	Left   float32 `yaml:"left"`
	Right  float32 `yaml:"right"`
	Bottom float32 `yaml:"bottom"`
	Top    float32 `yaml:"top"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (matrixConfig *OrthogonalConfig) Decode() *Matrix.Matrix4x4 {
	return Matrix.Orthogonal(
		matrixConfig.Left,
		matrixConfig.Right,
		matrixConfig.Bottom,
		matrixConfig.Top,
		matrixConfig.Near,
		matrixConfig.Far,
	)
}
