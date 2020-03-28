package Camera

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Camera struct {
	ProjectionMatrix GeometryMath.Matrix4x4 `yaml:"projection"`
	ViewMatrix       GeometryMath.Matrix4x4 `yaml:"view"`
}
