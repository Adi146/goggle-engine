package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"reflect"
)

const (
	ScaleFactoryName = "scale"
)

func init() {
	AddType(ScaleFactoryName, reflect.TypeOf((*ScaleEntry)(nil)).Elem())
}

type ScaleEntry float32

func (entry *ScaleEntry) Decode() *Matrix.Matrix4x4 {
	return Matrix.Scale((float32)(*entry))
}
