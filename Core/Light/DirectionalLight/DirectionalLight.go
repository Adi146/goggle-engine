package DirectionalLight

import (
	"math"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type DirectionalLight struct {
	Direction Vector.Vector3 `yaml:"direction,flow"`

	Ambient  Vector.Vector3 `yaml:"ambient,flow"`
	Diffuse  Vector.Vector3 `yaml:"diffuse,flow"`
	Specular Vector.Vector3 `yaml:"specular,flow"`
}

func (light *DirectionalLight) Draw(shader Shader.IShaderProgram) error {
	return shader.BindObject(light)
}

func (light *DirectionalLight) GetPosition() *Vector.Vector3 {
	return light.Direction.MulScalar(-math.MaxFloat32)
}

func (light *DirectionalLight) Set(val DirectionalLight) {
	light = &val
}

func (light *DirectionalLight) Get() DirectionalLight {
	return *light
}

func (light *DirectionalLight) GetDirection() Vector.Vector3 {
	return light.Direction
}

func (light *DirectionalLight) SetDirection(direction Vector.Vector3) {
	light.Direction = direction
}

func (light *DirectionalLight) GetAmbient() Vector.Vector3 {
	return light.Ambient
}

func (light *DirectionalLight) SetAmbient(color Vector.Vector3) {
	light.Ambient = color
}

func (light *DirectionalLight) GetDiffuse() Vector.Vector3 {
	return light.Diffuse
}

func (light *DirectionalLight) SetDiffuse(color Vector.Vector3) {
	light.Diffuse = color
}

func (light *DirectionalLight) GetSpecular() Vector.Vector3 {
	return light.Specular
}

func (light *DirectionalLight) SetSpecular(color Vector.Vector3) {
	light.Specular = color
}
