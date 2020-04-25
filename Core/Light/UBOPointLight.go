package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	shadowMap "github.com/Adi146/goggle-engine/Core/Light/ShadowMapping/PointLight"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer/UniformBufferSection"
	"gopkg.in/yaml.v3"
)

const (
	pointLight_offset_position  = 0
	pointLight_offset_linear    = 12
	pointLight_offset_quadratic = 16
	pointLight_offset_ambient   = 32
	pointLight_offset_diffuse   = 48
	pointLight_offset_specular  = 64
	pointLight_offset_camera    = 80
	pointLight_offset_distance  = 464

	pointLight_size_section = 480
	pointLight_ubo_size     = UniformBuffer.ArrayUniformBuffer_offset_elements + UniformBuffer.Num_elements*pointLight_size_section

	PointLight_fbo_type = "shadowMap_pointLight"
)

type UBOPointLight struct {
	PointLight struct {
		Position  UniformBufferSection.Vector3 `yaml:"position,flow"`
		Linear    UniformBufferSection.Float   `yaml:"linear,flow"`
		Quadratic UniformBufferSection.Float   `yaml:"quadratic,flow"`
		Ambient   UniformBufferSection.Vector3 `yaml:"ambient,flow"`
		Diffuse   UniformBufferSection.Vector3 `yaml:"diffuse,flow"`
		Specular  UniformBufferSection.Vector3 `yaml:"specular,flow"`
	} `yaml:"pointLight"`

	ShadowMap ShadowMapping.ShadowMap `yaml:"shadowMap"`
}

func (light *UBOPointLight) ForceUpdate() {
	light.PointLight.Position.ForceUpdate()
	light.PointLight.Linear.ForceUpdate()
	light.PointLight.Quadratic.ForceUpdate()
	light.PointLight.Ambient.ForceUpdate()
	light.PointLight.Diffuse.ForceUpdate()
	light.PointLight.Specular.ForceUpdate()
	light.ShadowMap.ForceUpdate()
}

func (light *UBOPointLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.PointLight.Position.SetUniformBuffer(ubo, offset+pointLight_offset_position)
	light.PointLight.Linear.SetUniformBuffer(ubo, offset+pointLight_offset_linear)
	light.PointLight.Quadratic.SetUniformBuffer(ubo, offset+pointLight_offset_quadratic)
	light.PointLight.Ambient.SetUniformBuffer(ubo, offset+pointLight_offset_ambient)
	light.PointLight.Diffuse.SetUniformBuffer(ubo, offset+pointLight_offset_diffuse)
	light.PointLight.Specular.SetUniformBuffer(ubo, offset+pointLight_offset_specular)
	light.ShadowMap.Camera.(UniformBuffer.IUniformBufferSection).SetUniformBuffer(ubo, offset+pointLight_offset_camera)
	light.ShadowMap.Distance.SetUniformBuffer(ubo, offset+pointLight_offset_distance)
}

func (light *UBOPointLight) GetSize() int {
	return pointLight_size_section
}

func (light *UBOPointLight) UpdateCamera(scene Scene.IScene, camera Camera.ICamera) {
	light.ShadowMap.Camera.(*shadowMap.Camera).Update(light.PointLight.Position.Get())
}

func (light *UBOPointLight) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene, camera Camera.ICamera) error {
	return light.ShadowMap.Draw(shader, invoker, scene, camera)
}

func (light *UBOPointLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		Ptr UniformBuffer.ArrayUniformBufferPtr `yaml:"uniformBuffer"`
	}{
		Ptr: UniformBuffer.ArrayUniformBufferPtr{
			ArrayUniformBuffer: &UniformBuffer.ArrayUniformBuffer{
				UniformBuffer: &UniformBuffer.UniformBuffer{
					Size: pointLight_ubo_size,
					Type: ShadowMapping.PointLight_ubo_type,
				},
			},
		},
	}
	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	light.ShadowMap.Camera = &shadowMap.Camera{}
	light.ShadowMap.TextureType = ShadowMapping.ShadowMapPointLight
	light.ShadowMap.TextureConstructor = ShadowMapping.NewShadowCubeMapTexture
	light.ShadowMap.FrameBuffer.Type = PointLight_fbo_type
	light.ShadowMap.UpdateCameraCallback = light.UpdateCamera

	index, err := uboYamlConfig.Ptr.AddElement(light)
	if err != nil {
		return err
	}

	light.ShadowMap.LightIndex = index

	type yamlConfigType UBOPointLight
	yamlConfig := (*yamlConfigType)(light)

	err = value.Decode(&yamlConfig)
	if err != nil {
		return err
	}

	light.ShadowMap.Camera.SetProjection(&GeometryMath.PerspectiveConfig{
		Fovy:   90,
		Aspect: float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width) / float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Height),
		Near:   0.1,
		Far:    yamlConfig.ShadowMap.Distance.Get(),
	})

	return nil
}
