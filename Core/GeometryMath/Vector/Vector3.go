package Vector

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Vector3 [3]float32

func (v1 *Vector3) X() float32 {
	return v1[0]
}

func (v1 *Vector3) Y() float32 {
	return v1[1]
}

func (v1 *Vector3) Z() float32 {
	return v1[2]
}

func (v1 *Vector3) Invert() *Vector3 {
	return v1.MulScalar(-1)
}

func (v1 *Vector3) MulScalar(scalar float32) *Vector3 {
	return &Vector3{
		v1.X() * scalar,
		v1.Y() * scalar,
		v1.Z() * scalar,
	}
}

func (v1 *Vector3) MulVector(v2 *Vector3) *Vector3 {
	return &Vector3{
		v1.X() * v2.X(),
		v1.Y() * v2.Y(),
		v1.Z() * v2.Z(),
	}
}

func (v1 *Vector3) Add(v2 *Vector3) *Vector3 {
	return &Vector3{
		v1.X() + v2.X(),
		v1.Y() + v2.Y(),
		v1.Z() + v2.Z(),
	}
}

func (v1 *Vector3) Sub(v2 *Vector3) *Vector3 {
	return &Vector3{
		v1.X() - v2.X(),
		v1.Y() - v2.Y(),
		v1.Z() - v2.Z(),
	}
}

func (v1 *Vector3) Cross(v2 *Vector3) *Vector3 {
	return &Vector3{
		v1.Y()*v2.Z() - v1.Z()*v2.Y(),
		v1.Z()*v2.X() - v1.X()*v2.Z(),
		v1.X()*v2.Y() - v1.Y()*v2.X(),
	}
}

func (v1 *Vector3) Dot(v2 *Vector3) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}

func (v1 *Vector3) Length() float32 {
	return GeometryMath.Sqrt(GeometryMath.Pow(v1.X(), 2) + GeometryMath.Pow(v1.Y(), 2) + GeometryMath.Pow(v1.Z(), 2))
}

func (v1 *Vector3) Normalize() *Vector3 {
	length := v1.Length()
	inv := 1.0 / length
	return &Vector3{
		v1.X() * inv,
		v1.Y() * inv,
		v1.Z() * inv,
	}
}
