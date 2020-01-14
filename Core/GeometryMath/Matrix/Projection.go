package Matrix

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

func Orthogonal(left float32, right float32, bottom float32, top float32, near float32, far float32) *Matrix4x4 {
	rml := right - left
	tmb := top - bottom
	fmn := far - near

	return &Matrix4x4{
		{2 / rml, 0, 0, -((right + left) / rml)},
		{0, 2 / tmb, 0, -((top + bottom) / tmb)},
		{0, 0, -2 / fmn, -((far + near) / fmn)},
		{0, 0, 0, 1},
	}
}

func Perspective(fovy float32, aspect float32, near float32, far float32) *Matrix4x4 {
	fmn := far - near
	f := 1 / (GeometryMath.Tan(fovy / 2))

	return &Matrix4x4{
		{aspect * f, 0, 0, 0},
		{0, f, 0, 0},
		{0, 0, -((far + near) / (fmn)), -((2 * far * near) / fmn)},
		{0, 0, -1, 0},
	}
}

func LookAt(eye *Vector.Vector3, center *Vector.Vector3, up *Vector.Vector3) *Matrix4x4 {
	f := center.Sub(eye).Normalize()
	s := f.Cross(up.Normalize()).Normalize()
	u := s.Cross(f)

	return &Matrix4x4{
		{s.X(), s.Y(), s.Z(), -s.Dot(eye)},
		{u.X(), u.Y(), u.Z(), -u.Dot(eye)},
		{-f.X(), -f.Y(), -f.Z(), f.Dot(eye)},
		{0, 0, 0, 1},
	}
}
