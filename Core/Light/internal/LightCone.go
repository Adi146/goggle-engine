package internal

type LightCone struct {
	InnerCone float32 `yaml:"innerCone"`
	OuterCone float32 `yaml:"outerCone"`
}

func (light *LightCone) GetInnerCone() float32 {
	return light.InnerCone
}

func (light *LightCone) SetInnerCone(val float32) {
	light.InnerCone = val
}

func (light *LightCone) GetOuterCone() float32 {
	return light.OuterCone
}

func (light *LightCone) SetOuterCone(val float32) {
	light.OuterCone = val
}
