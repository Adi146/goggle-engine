package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light"
)

type DirectionalLight struct {
	Light.LightBase            `yaml:",inline"`
	Light.DirectionalLightBase `yaml:",inline"`
	Camera.Camera
}

func (light *DirectionalLight) Set(val DirectionalLight) {
	*light = val
}

func (light *DirectionalLight) Get() DirectionalLight {
	return *light
}
