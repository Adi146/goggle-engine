package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
)

type IPointLight interface {
	internal.ILightPosition
	internal.ILightColor
}
