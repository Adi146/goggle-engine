package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	shadowMap "github.com/Adi146/goggle-engine/Core/Light/ShadowMapping/SpotLight"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer/UniformBufferSection"
	"gopkg.in/yaml.v3"
)

const (
	spotLight_offset_position       = 0
	spotLight_offset_linear         = 12
	spotLight_offset_quadratic      = 16
	spotLight_offset_ambient        = 32
	spotLight_offset_diffuse        = 48
	spotLight_offset_specular       = 64
	spotLight_offset_direction      = 80
	spotLight_offset_innerCone      = 92
	spotLight_offset_outerCone      = 96
	spotLight_offset_viewProjection = 112

	spotLight_size_section = 176
	spotLight_ubo_size     = UniformBuffer.ArrayUniformBuffer_offset_elements + UniformBuffer.Num_elements*spotLight_size_section

	SpotLight_fbo_type = "shadowMap_spotLight"
)

type SpotLight struct {
	SpotLight struct {
		Position  UniformBufferSection.Vector3 `yaml:"position,flow"`
		Linear    UniformBufferSection.Float   `yaml:"linear,flow"`
		Quadratic UniformBufferSection.Float   `yaml:"quadratic,flow"`
		Ambient   UniformBufferSection.Vector3 `yaml:"ambient,flow"`
		Diffuse   UniformBufferSection.Vector3 `yaml:"diffuse,flow"`
		Specular  UniformBufferSection.Vector3 `yaml:"specular,flow"`
		Direction UniformBufferSection.Vector3 `yaml:"direction,flow"`
		InnerCone UniformBufferSection.Float   `yaml:"innerCone"`
		OuterCone UniformBufferSection.Float   `yaml:"outerCone"`
	} `yaml:"spotLight"`

	ShadowMap ShadowMapping.ShadowMap `yaml:"shadowMap"`
}

func (light *SpotLight) ForceUpdate() {
	light.SpotLight.Position.ForceUpdate()
	light.SpotLight.Linear.ForceUpdate()
	light.SpotLight.Quadratic.ForceUpdate()
	light.SpotLight.Ambient.ForceUpdate()
	light.SpotLight.Diffuse.ForceUpdate()
	light.SpotLight.Specular.ForceUpdate()
	light.SpotLight.Direction.ForceUpdate()
	light.SpotLight.InnerCone.ForceUpdate()
	light.SpotLight.OuterCone.ForceUpdate()
	light.ShadowMap.ForceUpdate()
}

func (light *SpotLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.SpotLight.Position.SetUniformBuffer(ubo, offset+spotLight_offset_position)
	light.SpotLight.Linear.SetUniformBuffer(ubo, offset+spotLight_offset_linear)
	light.SpotLight.Quadratic.SetUniformBuffer(ubo, offset+spotLight_offset_quadratic)
	light.SpotLight.Ambient.SetUniformBuffer(ubo, offset+spotLight_offset_ambient)
	light.SpotLight.Diffuse.SetUniformBuffer(ubo, offset+spotLight_offset_diffuse)
	light.SpotLight.Specular.SetUniformBuffer(ubo, offset+spotLight_offset_specular)
	light.SpotLight.Direction.SetUniformBuffer(ubo, offset+spotLight_offset_direction)
	light.SpotLight.InnerCone.SetUniformBuffer(ubo, offset+spotLight_offset_innerCone)
	light.SpotLight.OuterCone.SetUniformBuffer(ubo, offset+spotLight_offset_outerCone)
	light.ShadowMap.Camera.(UniformBuffer.IUniformBufferSection).SetUniformBuffer(ubo, offset+spotLight_offset_viewProjection)
}

func (light *SpotLight) GetSize() int {
	return spotLight_size_section
}

func (light *SpotLight) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	light.SpotLight.Position.Set(position)
	light.SpotLight.Direction.Set(front)

	light.ShadowMap.Camera.Update(position, front, up)
}

func (light *SpotLight) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene, camera Camera.ICamera) error {
	return light.ShadowMap.Draw(shader, invoker, scene, camera)
}

func (light *SpotLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		Ptr UniformBuffer.ArrayUniformBufferPtr `yaml:"uniformBuffer"`
	}{
		Ptr: UniformBuffer.ArrayUniformBufferPtr{
			ArrayUniformBuffer: &UniformBuffer.ArrayUniformBuffer{
				UniformBuffer: &UniformBuffer.UniformBuffer{
					Size: spotLight_ubo_size,
					Type: ShadowMapping.SpotLight_ubo_type,
				},
			},
		},
	}

	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	light.ShadowMap.Camera = &shadowMap.Camera{}
	light.ShadowMap.TextureType = ShadowMapping.ShadowMapSpotLight
	light.ShadowMap.TextureConstructor = ShadowMapping.NewShadowMapTexture
	light.ShadowMap.FrameBuffer.Type = SpotLight_fbo_type

	index, err := uboYamlConfig.Ptr.AddElement(light)
	if err != nil {
		return err
	}

	light.ShadowMap.LightIndex = index

	type yamlConfigType SpotLight
	yamlConfig := (*yamlConfigType)(light)

	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	light.SpotLight.InnerCone.Set(GeometryMath.Cos(GeometryMath.Radians(light.SpotLight.InnerCone.Get())))
	light.SpotLight.OuterCone.Set(GeometryMath.Cos(GeometryMath.Radians(light.SpotLight.OuterCone.Get())))

	light.ShadowMap.Camera.SetProjection(&GeometryMath.PerspectiveConfig{
		Fovy:   GeometryMath.Degree(light.SpotLight.OuterCone.Get()),
		Aspect: float32(light.ShadowMap.FrameBuffer.Viewport.Width) / float32(light.ShadowMap.FrameBuffer.Viewport.Height),
		Near:   0.1,
		Far:    light.ShadowMap.Distance.Get(),
	})

	return nil
}
