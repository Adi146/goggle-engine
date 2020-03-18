package internal

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type LightColor struct {
	Ambient  GeometryMath.Vector3 `yaml:"ambient,flow"`
	Diffuse  GeometryMath.Vector3 `yaml:"diffuse,flow"`
	Specular GeometryMath.Vector3 `yaml:"specular,flow"`
}

func (light *LightColor) GetAmbient() GeometryMath.Vector3 {
	return light.Ambient
}

func (light *LightColor) SetAmbient(color GeometryMath.Vector3) {
	light.Ambient = color
}

func (light *LightColor) GetDiffuse() GeometryMath.Vector3 {
	return light.Diffuse
}

func (light *LightColor) SetDiffuse(color GeometryMath.Vector3) {
	light.Diffuse = color
}

func (light *LightColor) GetSpecular() GeometryMath.Vector3 {
	return light.Specular
}

func (light *LightColor) SetSpecular(color GeometryMath.Vector3) {
	light.Specular = color
}
