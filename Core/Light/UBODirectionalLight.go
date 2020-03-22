package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"gopkg.in/yaml.v3"
)

const (
	directionalLight_offset_direction = 0
	directionalLight_offset_color     = 16
	directionalLight_offset_camera    = 64

	directionalLIght_size_section = directionalLight_ubo_size

	directionalLight_ubo_size = 192
	DirectionalLight_ubo_type = "directionalLight"
)

type UBODirectionalLight struct {
	internal.LightDirectionSection
	internal.LightColorSection
	ShadowMap struct {
		CameraSection Camera.CameraSection
		Shader        Shader.IShaderProgram
		FrameBuffer   ShadowMapping.FrameBuffer
	}
}

func (light *UBODirectionalLight) ForceUpdate() {
	light.LightDirectionSection.ForceUpdate()
	light.LightColorSection.ForceUpdate()
	light.ShadowMap.CameraSection.ForceUpdate()
}

func (light *UBODirectionalLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.LightDirectionSection.SetUniformBuffer(ubo, offset+directionalLight_offset_direction)
	light.LightColorSection.SetUniformBuffer(ubo, offset+directionalLight_offset_color)
	light.ShadowMap.CameraSection.SetUniformBuffer(ubo, offset+directionalLight_offset_camera)
}

func (light *UBODirectionalLight) GetSize() int {
	return directionalLIght_size_section
}

func (light *UBODirectionalLight) SetDirection(val GeometryMath.Vector3) {
	light.LightDirectionSection.SetDirection(val)
	light.ShadowMap.CameraSection.SetViewMatrix(*GeometryMath.LookAt(val.Invert(), &GeometryMath.Vector3{0, 0, 0}, &GeometryMath.Vector3{0, 1, 0}))
}

func (light *UBODirectionalLight) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
	if invoker == light {
		return nil
	}
	defer FrameBuffer.GetCurrentFrameBuffer().Bind()
	defer Function.GetCurrentCullFunction().Set()
	defer Function.GetCurrentDepthFunction().Set()
	defer Function.GetCurrentBlendFunction().Set()

	light.ShadowMap.FrameBuffer.Bind()
	Function.Front.Set()
	Function.Less.Set()
	Function.DisabledBlend.Set()

	light.ShadowMap.FrameBuffer.Clear()

	if shader != nil {
		defer shader.Bind()
	}

	return scene.Draw(light.ShadowMap.Shader, light, scene)
}

func (light *UBODirectionalLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		UniformBuffer *UniformBuffer.UniformBuffer `yaml:"uniformBuffer"`
	}{
		UniformBuffer: &UniformBuffer.UniformBuffer{
			Size: directionalLight_ubo_size,
			Type: DirectionalLight_ubo_type,
		},
	}
	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	type Light struct {
		DirectionSection internal.LightDirectionSection `yaml:",inline"`
		ColorSection     internal.LightColorSection     `yaml:",inline"`
	}

	type ShadowMap struct {
		CameraSection Camera.CameraSection      `yaml:",inline"`
		Shader        Shader.Ptr                `yaml:"shader"`
		FrameBuffer   ShadowMapping.FrameBuffer `yaml:"frameBuffer"`
	}

	yamlConfig := struct {
		Light     `yaml:"directionalLight"`
		ShadowMap `yaml:"shadowMap"`
	}{
		Light: Light{
			DirectionSection: internal.LightDirectionSection{
				LightDirection: light.LightDirection,
				UniformBuffer:  uboYamlConfig.UniformBuffer,
				Offset:         directionalLight_offset_direction,
			},
			ColorSection: internal.LightColorSection{
				LightColor:    internal.LightColor{},
				UniformBuffer: uboYamlConfig.UniformBuffer,
				Offset:        directionalLight_offset_color,
			},
		},
		ShadowMap: ShadowMap{
			CameraSection: Camera.CameraSection{
				Camera:        light.ShadowMap.CameraSection.Camera,
				UniformBuffer: uboYamlConfig.UniformBuffer,
				Offset:        directionalLight_offset_camera,
			},
			Shader: Shader.Ptr{
				IShaderProgram: light.ShadowMap.Shader,
			},
			FrameBuffer: light.ShadowMap.FrameBuffer,
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return nil
	}

	light.LightDirectionSection = yamlConfig.Light.DirectionSection
	light.LightColorSection = yamlConfig.Light.ColorSection
	light.ShadowMap.CameraSection = yamlConfig.CameraSection
	light.ShadowMap.Shader = yamlConfig.Shader.IShaderProgram
	light.ShadowMap.FrameBuffer = yamlConfig.FrameBuffer

	return nil
}
