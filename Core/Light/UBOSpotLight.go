package Light

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer/UniformBufferSection"
	"github.com/Adi146/goggle-engine/Utils/Log"
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
	SpotLight_ubo_type     = "spotLight"

	SpotLight_fbo_type = "shadowMap_spotLight"
)

type UBOSpotLight struct {
	Position  UniformBufferSection.Vector3
	Linear    UniformBufferSection.Float
	Quadratic UniformBufferSection.Float
	Ambient   UniformBufferSection.Vector3
	Diffuse   UniformBufferSection.Vector3
	Specular  UniformBufferSection.Vector3
	Direction UniformBufferSection.Vector3
	InnerCone UniformBufferSection.Float
	OuterCone UniformBufferSection.Float

	ShadowMap struct {
		Projection     GeometryMath.Matrix4x4
		ViewProjection UniformBufferSection.Matrix4x4
		Shader         Shader.IShaderProgram
		FrameBuffer    FrameBuffer.FrameBuffer
		Index          int
	}
}

func (light *UBOSpotLight) ForceUpdate() {
	light.Position.ForceUpdate()
	light.Linear.ForceUpdate()
	light.Quadratic.ForceUpdate()
	light.Ambient.ForceUpdate()
	light.Diffuse.ForceUpdate()
	light.Specular.ForceUpdate()
	light.Direction.ForceUpdate()
	light.InnerCone.ForceUpdate()
	light.OuterCone.ForceUpdate()
	light.ShadowMap.ViewProjection.ForceUpdate()
}

func (light *UBOSpotLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.Position.SetUniformBuffer(ubo, offset+spotLight_offset_position)
	light.Linear.SetUniformBuffer(ubo, offset+spotLight_offset_linear)
	light.Quadratic.SetUniformBuffer(ubo, offset+spotLight_offset_quadratic)
	light.Ambient.SetUniformBuffer(ubo, offset+spotLight_offset_ambient)
	light.Diffuse.SetUniformBuffer(ubo, offset+spotLight_offset_diffuse)
	light.Specular.SetUniformBuffer(ubo, offset+spotLight_offset_specular)
	light.Direction.SetUniformBuffer(ubo, offset+spotLight_offset_direction)
	light.InnerCone.SetUniformBuffer(ubo, offset+spotLight_offset_innerCone)
	light.OuterCone.SetUniformBuffer(ubo, offset+spotLight_offset_outerCone)
	light.ShadowMap.ViewProjection.SetUniformBuffer(ubo, offset+spotLight_offset_viewProjection)
}

func (light *UBOSpotLight) GetSize() int {
	return spotLight_size_section
}

func (light *UBOSpotLight) UpdateViewProjection() {
	pos := light.Position.Get()
	dir := light.Direction.Get()
	light.ShadowMap.ViewProjection.Set(light.ShadowMap.Projection.Mul(GeometryMath.LookAt(pos, pos.Add(dir), GeometryMath.Vector3{0, 1, 0})))
}

func (light *UBOSpotLight) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	_, isPointLight := invoker.(*UBOPointLight)
	_, isDirectionalLight := invoker.(*UBODirectionalLight)
	_, isSpotLight := invoker.(*UBOSpotLight)
	if isPointLight || isDirectionalLight || isSpotLight {
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

	if err := light.ShadowMap.Shader.BindObject(int32(light.ShadowMap.Index)); err != nil {
		return err
	}

	return scene.Draw(light.ShadowMap.Shader, light, scene)
}

func (light *UBOSpotLight) UnmarshalYAML(value *yaml.Node) error {
	uboYamlConfig := struct {
		Ptr UniformBuffer.ArrayUniformBufferPtr `yaml:"uniformBuffer"`
	}{
		Ptr: UniformBuffer.ArrayUniformBufferPtr{
			ArrayUniformBuffer: &UniformBuffer.ArrayUniformBuffer{
				UniformBuffer: &UniformBuffer.UniformBuffer{
					Size: spotLight_ubo_size,
					Type: SpotLight_ubo_type,
				},
			},
		},
	}

	if err := value.Decode(&uboYamlConfig); err != nil {
		return err
	}

	type ShadowMap struct {
		Distance    float32                 `yaml:"distance"`
		Shader      Shader.Ptr              `yaml:"shader"`
		FrameBuffer FrameBuffer.FrameBuffer `yaml:"frameBuffer"`
		Shaders     []Shader.Ptr            `yaml:"bindOnShaders"`
	}

	yamlConfig := struct {
		SpotLight `yaml:"spotLight"`
		ShadowMap `yaml:"shadowMap"`
	}{
		SpotLight: SpotLight{
			Position:  light.Position.Get(),
			Linear:    light.Linear.Get(),
			Quadratic: light.Quadratic.Get(),
			Ambient:   light.Ambient.Get(),
			Diffuse:   light.Diffuse.Get(),
			Specular:  light.Specular.Get(),
			Direction: light.Direction.Get(),
		},
		ShadowMap: ShadowMap{
			Shader: Shader.Ptr{
				IShaderProgram: light.ShadowMap.Shader,
			},
			FrameBuffer: FrameBuffer.FrameBuffer{
				Type: SpotLight_fbo_type,
			},
		},
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	texture, err := NewShadowMap(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width, yamlConfig.ShadowMap.FrameBuffer.Viewport.Height, ShadowMapSpotLight)
	if err != nil {
		return err
	}

	yamlConfig.ShadowMap.FrameBuffer.AddDepthAttachment(texture)
	if err := yamlConfig.ShadowMap.FrameBuffer.Finish(); err != nil {
		return err
	}

	light.ShadowMap.Projection = GeometryMath.Perspective(yamlConfig.OuterCone, float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width)/float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Height), 0.1, yamlConfig.ShadowMap.Distance)
	light.Position.Set(yamlConfig.SpotLight.Position)
	light.Linear.Set(yamlConfig.SpotLight.Linear)
	light.Quadratic.Set(yamlConfig.SpotLight.Quadratic)
	light.Ambient.Set(yamlConfig.SpotLight.Ambient)
	light.Diffuse.Set(yamlConfig.SpotLight.Diffuse)
	light.Specular.Set(yamlConfig.SpotLight.Specular)
	light.Direction.Set(yamlConfig.SpotLight.Direction)
	light.InnerCone.Set(GeometryMath.Cos(GeometryMath.Radians(yamlConfig.InnerCone)))
	light.OuterCone.Set(GeometryMath.Cos(GeometryMath.Radians(yamlConfig.OuterCone)))
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
