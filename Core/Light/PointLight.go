package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
)

type PointLight struct {
	internal.LightPosition `yaml:",inline"`
	internal.LightColor    `yaml:",inline"`
}
