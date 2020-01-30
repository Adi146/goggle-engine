package YamlFactory

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"reflect"
)

var MatrixFactory = map[string]reflect.Type{
	"translation": reflect.TypeOf((*TranslationConfig)(nil)).Elem(),
	"rotation":    reflect.TypeOf((*RotationConfig)(nil)).Elem(),
	"scale":       reflect.TypeOf((*ScaleConfig)(nil)).Elem(),
	"orthogonal":  reflect.TypeOf((*OrthogonalConfig)(nil)).Elem(),
	"perspective": reflect.TypeOf((*PerspectiveConfig)(nil)).Elem(),
}

type IYamlMatrixConfig interface {
	Decode() *Matrix.Matrix4x4
}

type TranslationConfig Vector.Vector3

func (matrixConfig *TranslationConfig) Decode() *Matrix.Matrix4x4 {
	return Matrix.Translate((*Vector.Vector3)(matrixConfig))
}

type RotationConfig struct {
	Vector *Vector.Vector3 `yaml:"axis"`
	Angle  float32         `yaml:"angle"`
}

func (matrixConfig *RotationConfig) Decode() *Matrix.Matrix4x4 {
	return Matrix.Rotate(Angle.Radians(matrixConfig.Angle), matrixConfig.Vector)
}

type ScaleConfig float32

func (matrixConfig *ScaleConfig) Decode() *Matrix.Matrix4x4 {
	return Matrix.Scale((float32)(*matrixConfig))
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
