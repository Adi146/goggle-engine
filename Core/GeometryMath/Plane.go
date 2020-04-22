package GeometryMath

type Plane3 struct {
	Normal Vector3
	Point  Vector3
	D      float32
}

func NewPlane(normal, point Vector3) Plane3 {
	return Plane3{
		Normal: normal,
		Point:  point,
		D:      -normal.Dot(point),
	}
}

func NewPlaneFromThreePoints(p0, p1, p2 Vector3) Plane3 {
	aux1 := p0.Sub(p1)
	aux2 := p2.Sub(p1)

	normal := aux2.Cross(aux1)
	normal = normal.Normalize()

	return Plane3{
		Normal: normal,
		Point:  p1,
		D:      -normal.Dot(p1),
	}
}

func (plane *Plane3) Distance(p Vector3) float32 {
	return plane.D + plane.Normal.Dot(p)
}
