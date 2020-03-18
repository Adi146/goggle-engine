package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
)

type IDirectionalLight interface {
	internal.ILightDirection
	internal.ILightColor
}
