package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"reflect"
)

const (
	OrthogonalFactoryName = "orthogonal"
)

func init() {
	AddType(OrthogonalFactoryName, reflect.TypeOf((*OrthogonalEntry)(nil)).Elem())
}

type OrthogonalEntry struct {
	Left   float32 `yaml:"left"`
	Right  float32 `yaml:"right"`
	Bottom float32 `yaml:"bottom"`
	Top    float32 `yaml:"top"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (entry *OrthogonalEntry) Decode() *Matrix.Matrix4x4 {
	return Matrix.Orthogonal(
		entry.Left,
		entry.Right,
		entry.Bottom,
		entry.Top,
		entry.Near,
		entry.Far,
	)
}
