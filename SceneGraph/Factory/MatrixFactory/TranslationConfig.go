package MatrixFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"reflect"
)

const (
	TranslationFactoryName = "translation"
)

func init() {
	AddType(TranslationFactoryName, reflect.TypeOf((*TranslationConfig)(nil)).Elem())
}

type TranslationConfig Vector.Vector3

func (matrixConfig *TranslationConfig) Decode() *Matrix.Matrix4x4 {
	return Matrix.Translate((*Vector.Vector3)(matrixConfig))
}
