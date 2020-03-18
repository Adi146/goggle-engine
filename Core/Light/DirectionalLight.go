package Light

import "github.com/Adi146/goggle-engine/Core/Light/internal"

type DirectionalLight struct {
	internal.LightDirection `yaml:",inline"`
	internal.LightColor     `yaml:",inline"`
}
