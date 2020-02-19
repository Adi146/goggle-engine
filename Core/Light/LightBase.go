package Light

import "github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"

type LightBase struct {
	Ambient  Vector.Vector3 `yaml:"ambient,flow"`
	Diffuse  Vector.Vector3 `yaml:"diffuse,flow"`
	Specular Vector.Vector3 `yaml:"specular,flow"`
}

func (light *LightBase) GetAmbient() Vector.Vector3 {
	return light.Ambient
}

func (light *LightBase) SetAmbient(color Vector.Vector3) {
	light.Ambient = color
}

func (light *LightBase) GetDiffuse() Vector.Vector3 {
	return light.Diffuse
}

func (light *LightBase) SetDiffuse(color Vector.Vector3) {
	light.Diffuse = color
}

func (light *LightBase) GetSpecular() Vector.Vector3 {
	return light.Specular
}

func (light *LightBase) SetSpecular(color Vector.Vector3) {
	light.Specular = color
}
