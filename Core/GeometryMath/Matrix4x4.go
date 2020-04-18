package GeometryMath

type Matrix4x4 [4][4]float32

func (m1 Matrix4x4) Mul(m2 Matrix4x4) Matrix4x4 {
	return Matrix4x4{
		{
			(m1[0][0] * m2[0][0]) + (m1[0][1] * m2[1][0]) + (m1[0][2] * m2[2][0]) + (m1[0][3] * m2[3][0]),
			(m1[0][0] * m2[0][1]) + (m1[0][1] * m2[1][1]) + (m1[0][2] * m2[2][1]) + (m1[0][3] * m2[3][1]),
			(m1[0][0] * m2[0][2]) + (m1[0][1] * m2[1][2]) + (m1[0][2] * m2[2][2]) + (m1[0][3] * m2[3][2]),
			(m1[0][0] * m2[0][3]) + (m1[0][1] * m2[1][3]) + (m1[0][2] * m2[2][3]) + (m1[0][3] * m2[3][3]),
		},
		{
			(m1[1][0] * m2[0][0]) + (m1[1][1] * m2[1][0]) + (m1[1][2] * m2[2][0]) + (m1[1][3] * m2[3][0]),
			(m1[1][0] * m2[0][1]) + (m1[1][1] * m2[1][1]) + (m1[1][2] * m2[2][1]) + (m1[1][3] * m2[3][1]),
			(m1[1][0] * m2[0][2]) + (m1[1][1] * m2[1][2]) + (m1[1][2] * m2[2][2]) + (m1[1][3] * m2[3][2]),
			(m1[1][0] * m2[0][3]) + (m1[1][1] * m2[1][3]) + (m1[1][2] * m2[2][3]) + (m1[1][3] * m2[3][3]),
		},
		{
			(m1[2][0] * m2[0][0]) + (m1[2][1] * m2[1][0]) + (m1[2][2] * m2[2][0]) + (m1[2][3] * m2[3][0]),
			(m1[2][0] * m2[0][1]) + (m1[2][1] * m2[1][1]) + (m1[2][2] * m2[2][1]) + (m1[2][3] * m2[3][1]),
			(m1[2][0] * m2[0][2]) + (m1[2][1] * m2[1][2]) + (m1[2][2] * m2[2][2]) + (m1[2][3] * m2[3][2]),
			(m1[2][0] * m2[0][3]) + (m1[2][1] * m2[1][3]) + (m1[2][2] * m2[2][3]) + (m1[2][3] * m2[3][3]),
		},
		{
			(m1[3][0] * m2[0][0]) + (m1[3][1] * m2[1][0]) + (m1[3][2] * m2[2][0]) + (m1[3][3] * m2[3][0]),
			(m1[3][0] * m2[0][1]) + (m1[3][1] * m2[1][1]) + (m1[3][2] * m2[2][1]) + (m1[3][3] * m2[3][1]),
			(m1[3][0] * m2[0][2]) + (m1[3][1] * m2[1][2]) + (m1[3][2] * m2[2][2]) + (m1[3][3] * m2[3][2]),
			(m1[3][0] * m2[0][3]) + (m1[3][1] * m2[1][3]) + (m1[3][2] * m2[2][3]) + (m1[3][3] * m2[3][3]),
		},
	}
}

func (m1 Matrix4x4) MulScalar(scalar float32) Matrix4x4 {
	return Matrix4x4{
		{m1[0][0] * scalar, m1[0][1] * scalar, m1[0][2] * scalar, m1[0][3] * scalar},
		{m1[1][0] * scalar, m1[1][1] * scalar, m1[1][2] * scalar, m1[1][3] * scalar},
		{m1[2][0] * scalar, m1[2][1] * scalar, m1[2][2] * scalar, m1[2][3] * scalar},
		{m1[3][0] * scalar, m1[3][1] * scalar, m1[3][2] * scalar, m1[3][3] * scalar},
	}
}

func (m1 Matrix4x4) MulVector(v1 Vector3) Vector3 {
	return Vector3{
		(v1.X() * m1[0][0]) + (v1.Y() * m1[0][1]) + (v1.Z() * m1[0][2]) + m1[0][3],
		(v1.X() * m1[1][0]) + (v1.Y() * m1[1][1]) + (v1.Z() * m1[1][2]) + m1[1][3],
		(v1.X() * m1[2][0]) + (v1.Y() * m1[2][1]) + (v1.Z() * m1[2][2]) + m1[2][3],
	}
}

func (m1 Matrix4x4) Transpose() Matrix4x4 {
	return Matrix4x4{
		{m1[0][0], m1[1][0], m1[2][0], m1[3][0]},
		{m1[0][1], m1[1][1], m1[2][1], m1[3][1]},
		{m1[0][2], m1[1][2], m1[2][2], m1[3][2]},
		{m1[0][3], m1[1][3], m1[2][3], m1[3][3]},
	}
}

func (m1 Matrix4x4) Inverse() Matrix4x4 {
	Coef00 := m1[2][2]*m1[3][3] - m1[3][2]*m1[2][3]
	Coef02 := m1[1][2]*m1[3][3] - m1[3][2]*m1[1][3]
	Coef03 := m1[1][2]*m1[2][3] - m1[2][2]*m1[1][3]
	Coef04 := m1[2][1]*m1[3][3] - m1[3][1]*m1[2][3]
	Coef06 := m1[1][1]*m1[3][3] - m1[3][1]*m1[1][3]
	Coef07 := m1[1][1]*m1[2][3] - m1[2][1]*m1[1][3]
	Coef08 := m1[2][1]*m1[3][2] - m1[3][1]*m1[2][2]
	Coef10 := m1[1][1]*m1[3][2] - m1[3][1]*m1[1][2]
	Coef11 := m1[1][1]*m1[2][2] - m1[2][1]*m1[1][2]
	Coef12 := m1[2][0]*m1[3][3] - m1[3][0]*m1[2][3]
	Coef14 := m1[1][0]*m1[3][3] - m1[3][0]*m1[1][3]
	Coef15 := m1[1][0]*m1[2][3] - m1[2][0]*m1[1][3]
	Coef16 := m1[2][0]*m1[3][2] - m1[3][0]*m1[2][2]
	Coef18 := m1[1][0]*m1[3][2] - m1[3][0]*m1[1][2]
	Coef19 := m1[1][0]*m1[2][2] - m1[2][0]*m1[1][2]
	Coef20 := m1[2][0]*m1[3][1] - m1[3][0]*m1[2][1]
	Coef22 := m1[1][0]*m1[3][1] - m1[3][0]*m1[1][1]
	Coef23 := m1[1][0]*m1[2][1] - m1[2][0]*m1[1][1]

	Fac0 := &Vector4{Coef00, Coef00, Coef02, Coef03}
	Fac1 := &Vector4{Coef04, Coef04, Coef06, Coef07}
	Fac2 := &Vector4{Coef08, Coef08, Coef10, Coef11}
	Fac3 := &Vector4{Coef12, Coef12, Coef14, Coef15}
	Fac4 := &Vector4{Coef16, Coef16, Coef18, Coef19}
	Fac5 := &Vector4{Coef20, Coef20, Coef22, Coef23}

	Vec0 := Vector4{m1[1][0], m1[0][0], m1[0][0], m1[0][0]}
	Vec1 := Vector4{m1[1][1], m1[0][1], m1[0][1], m1[0][1]}
	Vec2 := Vector4{m1[1][2], m1[0][2], m1[0][2], m1[0][2]}
	Vec3 := Vector4{m1[1][3], m1[0][3], m1[0][3], m1[0][3]}

	Inv0 := Vec1.MulVector(Fac0).Sub(Vec2.MulVector(Fac1)).Add(Vec3.MulVector(Fac2))
	Inv1 := Vec0.MulVector(Fac0).Sub(Vec2.MulVector(Fac3)).Add(Vec3.MulVector(Fac4))
	Inv2 := Vec0.MulVector(Fac1).Sub(Vec1.MulVector(Fac3)).Add(Vec3.MulVector(Fac5))
	Inv3 := Vec0.MulVector(Fac2).Sub(Vec1.MulVector(Fac4)).Add(Vec2.MulVector(Fac5))

	SignA := &Vector4{1, -1, 1, -1}
	SignB := &Vector4{-1, 1, -1, +1}
	Inverse := Matrix4x4{*Inv0.MulVector(SignA), *Inv1.MulVector(SignB), *Inv2.MulVector(SignA), *Inv3.MulVector(SignB)}

	Row0 := &Vector4{Inverse[0][0], Inverse[1][0], Inverse[2][0], Inverse[3][0]}
	Column0 := Vector4(m1[0])

	Dot0 := Column0.MulVector(Row0)
	Dot1 := (Dot0.X() + Dot0.Y()) + (Dot0.Z() + Dot0.W())

	OneOverDeterminant := 1 / Dot1

	return Inverse.MulScalar(OneOverDeterminant)
}

func (m1 Matrix4x4) Equals(m2 Matrix4x4, threshold float32) bool {
	for i, _ := range m1 {
		for j, _ := range m1[i] {
			if !Equals(m1[i][j], m2[i][j], threshold) {
				return false
			}
		}
	}

	return true
}

func (m1 Matrix4x4) ToMatrix3x3() Matrix3x3 {
	return Matrix3x3{
		{m1[0][0], m1[0][1], m1[0][2]},
		{m1[1][0], m1[1][1], m1[1][2]},
		{m1[2][0], m1[2][1], m1[2][2]},
	}
}
