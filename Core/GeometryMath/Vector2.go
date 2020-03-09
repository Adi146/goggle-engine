package GeometryMath

type Vector2 [2]float32

func (v1 *Vector2) X() float32 {
	return v1[0]
}

func (v1 *Vector2) Y() float32 {
	return v1[1]
}

func (v1 *Vector2) Invert() *Vector2 {
	return v1.MulScalar(-1)
}

func (v1 *Vector2) MulScalar(scalar float32) *Vector2 {
	return &Vector2{
		v1.X() * scalar,
		v1.Y() * scalar,
	}
}

func (v1 *Vector2) MulVector(v2 *Vector2) *Vector2 {
	return &Vector2{
		v1.X() * v2.X(),
		v1.Y() * v2.Y(),
	}
}

func (v1 *Vector2) Add(v2 *Vector2) *Vector2 {
	return &Vector2{
		v1.X() + v2.X(),
		v1.Y() + v2.Y(),
	}
}

func (v1 *Vector2) Sub(v2 *Vector2) *Vector2 {
	return &Vector2{
		v1.X() - v2.X(),
		v1.Y() - v2.Y(),
	}
}

func (v1 *Vector2) Dot(v2 *Vector2) float32 {
	return v1.X()*v2.X() + v1.Y()*v2.Y()
}

func (v1 *Vector2) Length() float32 {
	return Sqrt(Pow(v1.X(), 2) + Pow(v1.Y(), 2))
}

func (v1 *Vector2) Normalize() *Vector2 {
	length := v1.Length()
	inv := 1.0 / length
	return &Vector2{
		v1.X() * inv,
		v1.Y() * inv,
	}
}
