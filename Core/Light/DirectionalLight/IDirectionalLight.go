package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/Light"
)

type IDirectionalLight interface {
	Light.ILight
	Light.IDirectionalLight

	Set(light DirectionalLight)
	Get() DirectionalLight
}
