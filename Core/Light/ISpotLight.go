package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
)

type ISpotLight interface {
	internal.ILightPosition
	internal.ILightColor
	internal.ILightDirection
	internal.ILightCone
}
