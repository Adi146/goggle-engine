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
	AddType(TranslationFactoryName, reflect.TypeOf((*TranslationEntry)(nil)).Elem())
}

type TranslationEntry Vector.Vector3

func (entry *TranslationEntry) Decode() *Matrix.Matrix4x4 {
	return Matrix.Translate((*Vector.Vector3)(entry))
}
