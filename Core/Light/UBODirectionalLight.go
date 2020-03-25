package Light

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
)

const (
	directionalLight_offset_direction = 0
	directionalLight_offset_color     = internal.Size_lightDirectionSection
	directionalLight_offset_camera    = internal.Size_lightDirectionSection + internal.Size_lightColorSection

	directionalLight_size_section = internal.Size_lightDirectionSection + internal.Size_lightColorSection + internal.Size_shadowCameraSection

	directionalLight_ubo_size = directionalLight_size_section
	DirectionalLight_ubo_type = "directionalLight"

	DirectionalLight_fbo_type = "shadowMap_directionalLight"
)

type UBODirectionalLight struct {
	internal.LightDirectionSection
	internal.LightColorSection
	ShadowMap struct {
		ShadowCameraSection internal.ShadowCameraSection
		Shader              Shader.IShaderProgram
		FrameBuffer         FrameBuffer.FrameBuffer
	}
}

func (light *UBODirectionalLight) ForceUpdate() {
	light.LightDirectionSection.ForceUpdate()
	light.LightColorSection.ForceUpdate()
	light.ShadowMap.ShadowCameraSection.ForceUpdate()
}

func (light *UBODirectionalLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.LightDirectionSection.SetUniformBuffer(ubo, offset+directionalLight_offset_direction)
	light.LightColorSection.SetUniformBuffer(ubo, offset+directionalLight_offset_color)
	light.ShadowMap.ShadowCameraSection.SetUniformBuffer(ubo, offset+directionalLight_offset_camera)
}

func (light *UBODirectionalLight) GetSize() int {
	return directionalLight_size_section
}

func (light *UBODirectionalLight) SetDirection(val GeometryMath.Vector3) {
	light.LightDirectionSection.SetDirection(val)
	light.ShadowMap.ShadowCameraSection.SetViewMatrix(*GeometryMath.LookAt(val.Invert(), &GeometryMath.Vector3{0, 0, 0}, &GeometryMath.Vector3{0, 1, 0}))
}

func (light *UBODirectionalLight) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
	_, isSpotLight := invoker.(IPointLight)
	_, isDirectionalLight := invoker.(IDirectionalLight)
	if isSpotLight || isDirectionalLight {
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

	return nil
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
		CameraSection internal.ShadowCameraSection `yaml:",inline"`
		Shader        Shader.Ptr                   `yaml:"shader"`
		FrameBuffer   FrameBuffer.FrameBuffer      `yaml:"frameBuffer"`
		Shaders       []Shader.Ptr                 `yaml:"bindOnShaders"`
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
			CameraSection: internal.ShadowCameraSection{
				Camera:        light.ShadowMap.ShadowCameraSection.Camera,
				UniformBuffer: uboYamlConfig.UniformBuffer,
				Offset:        directionalLight_offset_camera,
			},
			Shader: Shader.Ptr{
				IShaderProgram: light.ShadowMap.Shader,
			},
			FrameBuffer: FrameBuffer.FrameBuffer{
				Type: DirectionalLight_fbo_type,
			},
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return nil
	}

	texture, err := internal.NewShadowMap(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width, yamlConfig.ShadowMap.FrameBuffer.Viewport.Height)
	if err != nil {
		return err
	}
	yamlConfig.ShadowMap.FrameBuffer.AddDepthAttachment(texture)
	if err := yamlConfig.ShadowMap.FrameBuffer.Finish(); err != nil {
		return err
	}

	for _, shader := range yamlConfig.Shaders {
		if err := shader.BindObject(texture); err != nil {
			Log.Error(err, "")
		}
	}

	light.LightDirectionSection = yamlConfig.Light.DirectionSection
	light.LightColorSection = yamlConfig.Light.ColorSection
	light.ShadowMap.ShadowCameraSection = yamlConfig.ShadowMap.CameraSection
	light.ShadowMap.Shader = yamlConfig.ShadowMap.Shader.IShaderProgram
	light.ShadowMap.FrameBuffer = yamlConfig.ShadowMap.FrameBuffer

	return nil
}
