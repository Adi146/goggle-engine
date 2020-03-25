package Light

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light/internal"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
	"math"
)

const (
	pointLight_offset_position = 0
	pointLight_offset_color    = internal.Size_lightPositionSection
	pointLight_offset_camera   = internal.Size_lightPositionSection + internal.Size_lightColorSection
	pointLight_offset_distance = internal.Size_lightPositionSection + internal.Size_lightColorSection + 6*internal.Size_shadowCameraSection

	pointLight_size_section = internal.Size_lightPositionSection + internal.Size_lightColorSection + 6*internal.Size_shadowCameraSection + internal.Size_shadowDistanceSection

	pointLight_ubo_size                    = UniformBuffer.Std140_size_single + UniformBuffer.Num_elements*pointLight_size_section
	PointLight_ubo_type UniformBuffer.Type = "pointLight"

	PointLight_fbo_type = "shadowMap_pointLight"
)

type UBOPointLight struct {
	internal.LightPositionSection
	internal.LightColorSection
	ShadowMap struct {
		CameraSections  [6]internal.ShadowCameraSection
		DistanceSection internal.ShadowDistanceSection
		Shader          Shader.IShaderProgram
		FrameBuffer     FrameBuffer.FrameBuffer
		Index           int
	}
}

func (light *UBOPointLight) ForceUpdate() {
	light.LightPositionSection.ForceUpdate()
	light.LightColorSection.ForceUpdate()
	for i := range light.ShadowMap.CameraSections {
		light.ShadowMap.CameraSections[i].ForceUpdate()
	}
	light.ShadowMap.DistanceSection.ForceUpdate()
}

func (light *UBOPointLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.LightPositionSection.SetUniformBuffer(ubo, offset+pointLight_offset_position)
	light.LightColorSection.SetUniformBuffer(ubo, offset+pointLight_offset_color)
	for i := range light.ShadowMap.CameraSections {
		light.ShadowMap.CameraSections[i].SetUniformBuffer(ubo, offset+pointLight_offset_camera+i*internal.Size_shadowCameraSection)
	}
	light.ShadowMap.DistanceSection.SetUniformBuffer(ubo, offset+pointLight_offset_distance)
}

func (light *UBOPointLight) GetSize() int {
	return pointLight_size_section
}

func (light *UBOPointLight) SetPosition(pos GeometryMath.Vector3) {
	light.LightPositionSection.SetPosition(pos)
	light.ShadowMap.CameraSections[0].SetViewMatrix(*GeometryMath.LookAt(&pos, pos.Add(&GeometryMath.Vector3{1.0, 0.0, 0.0}), &GeometryMath.Vector3{0.0, -1.0, 0.0}))
	light.ShadowMap.CameraSections[1].SetViewMatrix(*GeometryMath.LookAt(&pos, pos.Add(&GeometryMath.Vector3{-1.0, 0.0, 0.0}), &GeometryMath.Vector3{0.0, -1.0, 0.0}))
	light.ShadowMap.CameraSections[2].SetViewMatrix(*GeometryMath.LookAt(&pos, pos.Add(&GeometryMath.Vector3{0.0, 1.0, 0.0}), &GeometryMath.Vector3{0.0, 0.0, 1.0}))
	light.ShadowMap.CameraSections[3].SetViewMatrix(*GeometryMath.LookAt(&pos, pos.Add(&GeometryMath.Vector3{0.0, -1.0, 0.0}), &GeometryMath.Vector3{0.0, 0.0, -1.0}))
	light.ShadowMap.CameraSections[4].SetViewMatrix(*GeometryMath.LookAt(&pos, pos.Add(&GeometryMath.Vector3{0.0, 0.0, 1.0}), &GeometryMath.Vector3{0.0, -1.0, 0.0}))
	light.ShadowMap.CameraSections[5].SetViewMatrix(*GeometryMath.LookAt(&pos, pos.Add(&GeometryMath.Vector3{0.0, 0.0, -1.0}), &GeometryMath.Vector3{0.0, -1.0, 0.0}))
}

func (light *UBOPointLight) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene) error {
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

	light.ShadowMap.Shader.BindObject(int32(light.ShadowMap.Index))
	return scene.Draw(light.ShadowMap.Shader, light, scene)
}

func (light *UBOPointLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		Ptr UniformBuffer.ArrayUniformBufferPtr `yaml:"uniformBuffer"`
	}{
		Ptr: UniformBuffer.ArrayUniformBufferPtr{
			ArrayUniformBuffer: &UniformBuffer.ArrayUniformBuffer{
				UniformBuffer: &UniformBuffer.UniformBuffer{
					Size: pointLight_ubo_size,
					Type: PointLight_ubo_type,
				},
			},
		},
	}
	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	type Light struct {
		PositionSection internal.LightPositionSection `yaml:",inline"`
		ColorSection    internal.LightColorSection    `yaml:",inline"`
	}

	type ShadowMap struct {
		Distance    internal.ShadowDistanceSection `yaml:"distance"`
		Shader      Shader.Ptr                     `yaml:"shader"`
		FrameBuffer FrameBuffer.FrameBuffer        `yaml:"frameBuffer"`
		Shaders     []Shader.Ptr                   `yaml:"bindOnShaders"`
	}

	yamlConfig := struct {
		Light     `yaml:"pointLight"`
		ShadowMap `yaml:"shadowMap"`
	}{
		Light: Light{
			PositionSection: internal.LightPositionSection{
				LightPosition: light.LightPositionSection.LightPosition,
				Offset:        pointLight_offset_position,
			},
			ColorSection: internal.LightColorSection{
				LightColor: light.LightColorSection.LightColor,
				Offset:     pointLight_offset_color,
			},
		},
		ShadowMap: ShadowMap{
			Distance: internal.ShadowDistanceSection{
				Distance: light.ShadowMap.DistanceSection.Distance,
				Offset:   pointLight_offset_distance,
			},
			Shader: Shader.Ptr{
				IShaderProgram: light.ShadowMap.Shader,
			},
			FrameBuffer: FrameBuffer.FrameBuffer{
				Type: PointLight_fbo_type,
			},
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	texture, err := internal.NewShadowCubeMap(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width, yamlConfig.ShadowMap.FrameBuffer.Viewport.Height)
	if err != nil {
		return err
	}

	yamlConfig.ShadowMap.FrameBuffer.AddDepthAttachment(texture)
	if err := yamlConfig.ShadowMap.FrameBuffer.Finish(); err != nil {
		return err
	}

	light.LightPositionSection = yamlConfig.Light.PositionSection
	light.LightColorSection = yamlConfig.Light.ColorSection
	for i := range light.ShadowMap.CameraSections {
		light.ShadowMap.CameraSections[i] = internal.ShadowCameraSection{
			Camera: Camera.Camera{
				ProjectionMatrix: *GeometryMath.Perspective(math.Pi/2, float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width)/float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Height), 1, yamlConfig.ShadowMap.Distance.Distance),
			},
			Offset: pointLight_offset_camera + i*internal.Size_shadowCameraSection,
		}
	}
	light.ShadowMap.DistanceSection = yamlConfig.ShadowMap.Distance
	light.ShadowMap.Shader = yamlConfig.ShadowMap.Shader.IShaderProgram
	light.ShadowMap.FrameBuffer = yamlConfig.ShadowMap.FrameBuffer
	index, err := uboYamlConfig.Ptr.AddElement(light)
	if err != nil {
		return err
	}
	light.ShadowMap.Index = index

	for _, shader := range yamlConfig.Shaders {
		uniformAddress, err := shader.GetUniformAddress(texture)
		if err != nil {
			return err
		}
		if err := shader.BindUniform(texture, fmt.Sprintf(uniformAddress, index)); err != nil {
			Log.Error(err, "")
		}
	}

	return nil
}
