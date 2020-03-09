package Light

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type LightBase struct {
	Ambient  GeometryMath.Vector3 `yaml:"ambient,flow"`
	Diffuse  GeometryMath.Vector3 `yaml:"diffuse,flow"`
	Specular GeometryMath.Vector3 `yaml:"specular,flow"`
}

func (light *LightBase) GetAmbient() GeometryMath.Vector3 {
	return light.Ambient
}

func (light *LightBase) SetAmbient(color GeometryMath.Vector3) {
	light.Ambient = color
}

func (light *LightBase) GetDiffuse() GeometryMath.Vector3 {
	return light.Diffuse
}

func (light *LightBase) SetDiffuse(color GeometryMath.Vector3) {
	light.Diffuse = color
}

func (light *LightBase) GetSpecular() GeometryMath.Vector3 {
	return light.Specular
}

func (light *LightBase) SetSpecular(color GeometryMath.Vector3) {
	light.Specular = color
}
