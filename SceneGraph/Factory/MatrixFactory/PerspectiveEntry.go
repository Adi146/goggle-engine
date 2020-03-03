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
	AddType(PerspectiveFactoryName, reflect.TypeOf((*PerspectiveEntry)(nil)).Elem())
}

type PerspectiveEntry struct {
	Fovy   float32 `yaml:"fovy"`
	Aspect float32 `yaml:"aspect"`
	Near   float32 `yaml:"near"`
	Far    float32 `yaml:"far"`
}

func (entry *PerspectiveEntry) Decode() *Matrix.Matrix4x4 {
	return Matrix.Perspective(
		Angle.Radians(entry.Fovy)/2,
		entry.Aspect,
		entry.Near,
		entry.Far,
	)
}
