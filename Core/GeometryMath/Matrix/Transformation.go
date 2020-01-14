package Matrix

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
)

func Identity() *Matrix4x4 {
	return &Matrix4x4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func Scale(v float32) *Matrix4x4 {
	return &Matrix4x4{
		{v, 0, 0, 0},
		{0, v, 0, 0},
		{0, 0, v, 0},
		{0, 0, 0, 1},
	}
}

func Rotate(a1 float32, v1 *Vector.Vector3) *Matrix4x4 {
	cosAlpha := GeometryMath.Cos(a1)
	sinAlpha := GeometryMath.Sin(a1)

	return &Matrix4x4{
		{cosAlpha + GeometryMath.Pow(v1.X(), 2)*(1-cosAlpha), v1.X()*v1.Y()*(1-cosAlpha) - v1.Z()*sinAlpha, v1.X()*v1.Z()*(1-cosAlpha) + v1.Y()*sinAlpha, 0},
		{v1.Y()*v1.X()*(1-cosAlpha) + v1.Z()*sinAlpha, cosAlpha + GeometryMath.Pow(v1.Y(), 2)*(1-cosAlpha), v1.Y()*v1.Z()*(1-cosAlpha) - v1.X()*sinAlpha, 0},
		{v1.Z()*v1.X()*(1-cosAlpha) - v1.Y()*sinAlpha, v1.Z()*v1.Y()*(1-cosAlpha) + v1.X()*sinAlpha, cosAlpha + GeometryMath.Pow(v1.Z(), 2)*(1-cosAlpha), 0},
		{0, 0, 0, 1},
	}
}

func RotateX(a1 float32) *Matrix4x4 {
	cosAlpha := GeometryMath.Cos(a1)
	sinAlpha := GeometryMath.Sin(a1)

	return &Matrix4x4{
		{1, 0, 0, 0},
		{0, cosAlpha, -sinAlpha, 0},
		{0, sinAlpha, cosAlpha, 0},
		{0, 0, 0, 1},
	}
}

func RotateY(a1 float32) *Matrix4x4 {
	cosAlpha := GeometryMath.Cos(a1)
	sinAlpha := GeometryMath.Sin(a1)

	return &Matrix4x4{
		{cosAlpha, 0, sinAlpha, 0},
		{0, 1, 0, 0},
		{-sinAlpha, 0, cosAlpha, 0},
		{0, 0, 0, 1},
	}
}

func RotateZ(a1 float32) *Matrix4x4 {
	cosAlpha := GeometryMath.Cos(a1)
	sinAlpha := GeometryMath.Sin(a1)

	return &Matrix4x4{
		{cosAlpha, -sinAlpha, 0, 0},
		{sinAlpha, cosAlpha, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func Translate(v1 *Vector.Vector3) *Matrix4x4 {
	return &Matrix4x4{
		{1, 0, 0, v1.X()},
		{0, 1, 0, v1.Y()},
		{0, 0, 1, v1.Z()},
		{0, 0, 0, 1},
	}
}
