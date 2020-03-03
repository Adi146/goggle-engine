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
	AddType(PerspectiveFactoryName, reflect.TypeOf((*PerspectiveConfig)(nil)).Elem())
}

type PerspectiveConfig struct {
	Fovy   float32 `yaml:"fovy"`
	Aspect float32 `yaml:"aspect"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (matrixConfig *PerspectiveConfig) Decode() *Matrix.Matrix4x4 {
	return Matrix.Perspective(
		Angle.Radians(matrixConfig.Fovy)/2,
		matrixConfig.Aspect,
		matrixConfig.Near,
		matrixConfig.Far,
	)
}
