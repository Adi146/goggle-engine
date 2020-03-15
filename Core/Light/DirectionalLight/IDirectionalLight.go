package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light"
)

type IDirectionalLight interface {
	Light.ILight
	Light.IDirectionalLight
	Camera.ICamera
}
