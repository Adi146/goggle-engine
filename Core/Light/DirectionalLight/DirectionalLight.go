package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/Light"
)

type DirectionalLight struct {
	Light.LightBase            `yaml:",inline"`
	Light.DirectionalLightBase `yaml:",inline"`
}
