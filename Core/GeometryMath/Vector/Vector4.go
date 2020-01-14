package Vector

import "github.com/Adi146/goggle-engine/Core/GeometryMath"

type Vector4 [4]float32

func (v1 *Vector4) X() float32 {
	return v1[0]
}

func (v1 *Vector4) Y() float32 {
	return v1[1]
}

func (v1 *Vector4) Z() float32 {
	return v1[2]
}

func (v1 *Vector4) W() float32 {
	return v1[3]
}

func (v1 *Vector4) Invert() *Vector4 {
	return v1.MulScalar(-1)
}

func (v1 *Vector4) MulScalar(scalar float32) *Vector4 {
	return &Vector4{
		v1.X() * scalar,
		v1.Y() * scalar,
		v1.Z() * scalar,
		v1.W() * scalar,
	}
}

func (v1 *Vector4) MulVector(v2 *Vector4) *Vector4 {
	return &Vector4{
		v1.X() * v2.X(),
		v1.Y() * v2.Y(),
		v1.Z() * v2.Z(),
		v1.W() * v2.W(),
	}
}

func (v1 *Vector4) Add(v2 *Vector4) *Vector4 {
	return &Vector4{
		v1.X() + v2.X(),
		v1.Y() + v2.Y(),
		v1.Z() + v2.Z(),
		v1.W() + v2.W(),
	}
}

func (v1 *Vector4) Sub(v2 *Vector4) *Vector4 {
	return &Vector4{
		v1.X() - v2.X(),
		v1.Y() - v2.Y(),
		v1.Z() - v2.Z(),
		v1.W() - v2.W(),
	}
}

func (v1 *Vector4) Dot(v2 *Vector4) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z() + v1.W()*v2.W()
}

func (v1 *Vector4) Length() float32 {
	return GeometryMath.Sqrt(GeometryMath.Pow(v1.X(), 2) + GeometryMath.Pow(v1.Y(), 2) + GeometryMath.Pow(v1.Z(), 2) + GeometryMath.Pow(v1.W(), 2))
}

func (v1 *Vector4) Normalize() *Vector4 {
	length := v1.Length()
	inv := 1.0 / length
	return &Vector4{
		v1.X() * inv,
		v1.Y() * inv,
		v1.Z() * inv,
		v1.W() * inv,
	}
}
