package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
)

type SpotLight struct {
	internal.LightPosition  `yaml:",inline"`
	internal.LightColor     `yaml:",inline"`
	internal.LightDirection `yaml:",inline"`
	internal.LightCone      `yaml:",inline"`
}
