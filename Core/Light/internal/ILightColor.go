package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type ILightColor interface {
	GetAmbient() GeometryMath.Vector3
	SetAmbient(color GeometryMath.Vector3)

	GetDiffuse() GeometryMath.Vector3
	SetDiffuse(color GeometryMath.Vector3)

	GetSpecular() GeometryMath.Vector3
	SetSpecular(color GeometryMath.Vector3)
}
