package Light

import (
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	"github.com/Adi146/goggle-engine/Core/Scene"
)

type IPointLight interface {
	Scene.IDrawable
	internal.ILightPosition
	internal.ILightColor
}
