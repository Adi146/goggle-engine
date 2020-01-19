package Shader

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

type IShaderProgram interface {
	Bind()
	Unbind()

	Destroy()

	ResetIndexCounter()

	BindCamera(camera Camera.ICamera) error

	BindModel(mode *Model.Model) error
	BindMaterial(material *Model.Material) error

	BindDirectionalLight(light *Light.DirectionalLight) error
	BindPointLight(light *Light.PointLight) error
	BindSpotLight(light *Light.SpotLight) error

	BindTexture(textureSlot uint32, texture *Texture.Texture, uniformAddress string) error
	BindMatrix(modelMatrix *Matrix.Matrix4x4, uniformAddress string) error
}
