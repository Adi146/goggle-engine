package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"reflect"
)

const (
	ScaleFactoryName = "scale"
)

func init() {
	AddType(ScaleFactoryName, reflect.TypeOf((*ScaleProduct)(nil)).Elem())
}

type ScaleProduct float32

func (product *ScaleProduct) Decode() *Matrix.Matrix4x4 {
	return Matrix.Scale((float32)(*product))
}
