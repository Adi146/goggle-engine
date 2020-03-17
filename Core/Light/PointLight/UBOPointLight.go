package PointLight

import "github.com/Adi146/goggle-engine/Core/UniformBuffer"

const (
	ubo_size                    = UniformBuffer.Std140_size_single + UniformBuffer.Num_elements*section_size
	UBO_type UniformBuffer.Type = "pointLight"
)

type UBOPointLight struct {
	UBOSection
}
