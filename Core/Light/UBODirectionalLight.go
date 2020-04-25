package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	shadowMap "github.com/Adi146/goggle-engine/Core/Light/ShadowMapping/DirectionalLight"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer/UniformBufferSection"
	"gopkg.in/yaml.v3"
)

const (
	directionalLight_offset_direction          = 0
	directionalLight_offset_ambient            = 16
	directionalLight_offset_diffuse            = 32
	directionalLight_offset_specular           = 48
	directionalLight_offset_viewProjection     = 64
	directionalLight_offset_distance           = 128
	directionalLight_offset_transitionDistance = 132

	directionalLight_size_section = 144
	directionalLight_ubo_size     = UniformBuffer.ArrayUniformBuffer_offset_elements + UniformBuffer.Num_elements*directionalLight_size_section

	DirectionalLight_fbo_type = "shadowMap_directionalLight"
)

type UBODirectionalLight struct {
	DirectionalLight struct {
		Direction UniformBufferSection.Vector3 `yaml:"direction,flow"`
		Ambient   UniformBufferSection.Vector3 `yaml:"ambient,flow"`
		Diffuse   UniformBufferSection.Vector3 `yaml:"diffuse,flow"`
		Specular  UniformBufferSection.Vector3 `yaml:"specular,flow"`
	} `yaml:"directionalLight"`

	ShadowMap ShadowMapping.ShadowMap `yaml:"shadowMap"`
}

func (light *UBODirectionalLight) ForceUpdate() {
	light.DirectionalLight.Direction.ForceUpdate()
	light.DirectionalLight.Ambient.ForceUpdate()
	light.DirectionalLight.Diffuse.ForceUpdate()
	light.DirectionalLight.Specular.ForceUpdate()
	light.ShadowMap.ForceUpdate()
}

func (light *UBODirectionalLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.DirectionalLight.Direction.SetUniformBuffer(ubo, offset+directionalLight_offset_direction)
	light.DirectionalLight.Ambient.SetUniformBuffer(ubo, offset+directionalLight_offset_ambient)
	light.DirectionalLight.Diffuse.SetUniformBuffer(ubo, offset+directionalLight_offset_diffuse)
	light.DirectionalLight.Specular.SetUniformBuffer(ubo, offset+directionalLight_offset_specular)
	light.ShadowMap.Camera.(UniformBuffer.IUniformBufferSection).SetUniformBuffer(ubo, offset+directionalLight_offset_viewProjection)
	light.ShadowMap.Distance.SetUniformBuffer(ubo, offset+directionalLight_offset_distance)
	light.ShadowMap.TransitionDistance.SetUniformBuffer(ubo, offset+directionalLight_offset_transitionDistance)
}

func (light *UBODirectionalLight) GetSize() int {
	return directionalLight_size_section
}

func (light *UBODirectionalLight) UpdateCamera(scene Scene.IScene, camera Camera.ICamera) {
	light.ShadowMap.Camera.(*shadowMap.Camera).Update(camera, light.DirectionalLight.Direction.Get(), light.ShadowMap.Distance.Get())
}

func (light *UBODirectionalLight) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene, camera Camera.ICamera) error {
	return light.ShadowMap.Draw(shader, invoker, scene, camera)
}

func (light *UBODirectionalLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		Ptr *UniformBuffer.ArrayUniformBufferPtr `yaml:"uniformBuffer"`
	}{
		Ptr: &UniformBuffer.ArrayUniformBufferPtr{
			ArrayUniformBuffer: &UniformBuffer.ArrayUniformBuffer{
				UniformBuffer: &UniformBuffer.UniformBuffer{
					Size: directionalLight_ubo_size,
					Type: ShadowMapping.DirectionalLight_ubo_type,
				},
			},
		},
	}
	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	light.ShadowMap.Camera = &shadowMap.Camera{}
	light.ShadowMap.TextureType = ShadowMapping.ShadowMapDirectionalLight
	light.ShadowMap.TextureConstructor = ShadowMapping.NewShadowMapTexture
	light.ShadowMap.FrameBuffer.Type = DirectionalLight_fbo_type
	light.ShadowMap.UpdateCameraCallback = light.UpdateCamera

	index, err := uboYamlConfig.Ptr.AddElement(light)
	if err != nil {
		return err
	}

	light.ShadowMap.LightIndex = index

	type yamlConfigType UBODirectionalLight
	yamlConfig := (*yamlConfigType)(light)

	return value.Decode(&yamlConfig)
}
