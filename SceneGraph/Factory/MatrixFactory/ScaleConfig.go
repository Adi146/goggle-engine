package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"reflect"
)

const (
	ScaleFactoryName = "scale"
)

func init() {
	AddType(ScaleFactoryName, reflect.TypeOf((*ScaleConfig)(nil)).Elem())
}

type ScaleConfig float32

func (matrixConfig *ScaleConfig) Decode() *Matrix.Matrix4x4 {
	return Matrix.Scale((float32)(*matrixConfig))
}
