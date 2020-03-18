package internal

type ILightCone interface {
	GetInnerCone() float32
	SetInnerCone(val float32)

	GetOuterCone() float32
	SetOuterCone(val float32)
}
