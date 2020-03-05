package Angle

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"math"
)

type EulerAngles struct {
	Yaw                float32
	Pitch              float32
	Roll               float32
	GimbalLockDetected bool
}

func Radians(degree float32) float32 {
	return degree * math.Pi / 180
}

func Degree(radians float32) float32 {
	return radians * 180 / math.Pi
}

func ExtractFromMatrix(m *Matrix.Matrix4x4) []EulerAngles {
	if m[2][0] != 1 && m[2][0] != -1 {
		yaw1 := -GeometryMath.ASin(m[2][0])
		yaw2 := math.Pi - yaw1
		pitch1 := GeometryMath.ATan2(m[2][1]/GeometryMath.Cos(yaw1), m[2][2]/GeometryMath.Cos(yaw1))
		pitch2 := GeometryMath.ATan2(m[2][1]/GeometryMath.Cos(yaw2), m[2][2]/GeometryMath.Cos(yaw2))
		roll1 := GeometryMath.ATan2(m[1][0]/GeometryMath.Cos(yaw1), m[0][0]/GeometryMath.Cos(yaw1))
		roll2 := GeometryMath.ATan2(m[1][0]/GeometryMath.Cos(yaw2), m[0][0]/GeometryMath.Cos(yaw2))
		return []EulerAngles{
			{
				Yaw:                yaw1,
				Pitch:              pitch1,
				Roll:               roll1,
				GimbalLockDetected: false,
			},
			{
				Yaw:                yaw2,
				Pitch:              pitch2,
				Roll:               roll2,
				GimbalLockDetected: false,
			},
		}
	} else {
		roll := float32(0)
		if m[2][0] == -1 {
			yaw := float32(math.Pi) / 2
			pitch := roll + GeometryMath.ATan2(m[0][1], m[0][2])
			return []EulerAngles{
				{
					Yaw:                yaw,
					Pitch:              pitch,
					Roll:               roll,
					GimbalLockDetected: true,
				},
			}
		} else {
			yaw := -float32(math.Pi) / 2
			pitch := -roll + GeometryMath.ATan2(-m[0][1], -m[0][2])
			return []EulerAngles{
				{
					Yaw:                yaw,
					Pitch:              pitch,
					Roll:               roll,
					GimbalLockDetected: true,
				},
			}
		}
	}
}
