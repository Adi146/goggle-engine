package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
)

const (
	ubo_size                    = 192
	UBO_type UniformBuffer.Type = "directionalLight"
)

type UBODirectionalLight struct {
	UBOSection
	CameraSection Camera.UBOSection
}
