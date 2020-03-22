package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/Scene"
)

type IDirectionalLight interface {
	internal.ILightDirection
	internal.ILightColor
	Scene.IDrawable
}
