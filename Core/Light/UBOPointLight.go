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
	"math"
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
	PointLight_ubo_type     = "pointLight"

	PointLight_fbo_type = "shadowMap_pointLight"
)

type UBOPointLight struct {
	Position  UniformBufferSection.Vector3
	Linear    UniformBufferSection.Float
	Quadratic UniformBufferSection.Float
	Ambient   UniformBufferSection.Vector3
	Diffuse   UniformBufferSection.Vector3
	Specular  UniformBufferSection.Vector3
	ShadowMap struct {
		Projection     GeometryMath.Matrix4x4
		ViewProjection [6]UniformBufferSection.Matrix4x4
		Distance       UniformBufferSection.Float
		Shader         Shader.IShaderProgram
		FrameBuffer    FrameBuffer.FrameBuffer
		Index          int
	}
}

func (light *UBOPointLight) ForceUpdate() {
	light.Position.ForceUpdate()
	light.Linear.ForceUpdate()
	light.Quadratic.ForceUpdate()
	light.Ambient.ForceUpdate()
	light.Diffuse.ForceUpdate()
	light.Specular.ForceUpdate()
	for i := range light.ShadowMap.ViewProjection {
		light.ShadowMap.ViewProjection[i].ForceUpdate()
	}
	light.ShadowMap.Distance.ForceUpdate()
}

func (light *UBOPointLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.Position.SetUniformBuffer(ubo, offset+pointLight_offset_position)
	light.Linear.SetUniformBuffer(ubo, offset+pointLight_offset_linear)
	light.Quadratic.SetUniformBuffer(ubo, offset+pointLight_offset_quadratic)
	light.Ambient.SetUniformBuffer(ubo, offset+pointLight_offset_ambient)
	light.Diffuse.SetUniformBuffer(ubo, offset+pointLight_offset_diffuse)
	light.Specular.SetUniformBuffer(ubo, offset+pointLight_offset_specular)
	for i := range light.ShadowMap.ViewProjection {
		light.ShadowMap.ViewProjection[i].SetUniformBuffer(ubo, offset+pointLight_offset_camera+i*UniformBuffer.Std140_size_mat4)
	}
	light.ShadowMap.Distance.SetUniformBuffer(ubo, offset+pointLight_offset_distance)
}

func (light *UBOPointLight) GetSize() int {
	return pointLight_size_section
}

func (light *UBOPointLight) SetPosition(pos GeometryMath.Vector3) {
	light.Position.Set(pos)
	light.ShadowMap.ViewProjection[0].Set(light.ShadowMap.Projection.Mul(GeometryMath.LookAt(pos, pos.Add(GeometryMath.Vector3{1.0, 0.0, 0.0}), GeometryMath.Vector3{0.0, -1.0, 0.0})))
	light.ShadowMap.ViewProjection[1].Set(light.ShadowMap.Projection.Mul(GeometryMath.LookAt(pos, pos.Add(GeometryMath.Vector3{-1.0, 0.0, 0.0}), GeometryMath.Vector3{0.0, -1.0, 0.0})))
	light.ShadowMap.ViewProjection[2].Set(light.ShadowMap.Projection.Mul(GeometryMath.LookAt(pos, pos.Add(GeometryMath.Vector3{0.0, 1.0, 0.0}), GeometryMath.Vector3{0.0, 0.0, 1.0})))
	light.ShadowMap.ViewProjection[3].Set(light.ShadowMap.Projection.Mul(GeometryMath.LookAt(pos, pos.Add(GeometryMath.Vector3{0.0, -1.0, 0.0}), GeometryMath.Vector3{0.0, 0.0, -1.0})))
	light.ShadowMap.ViewProjection[4].Set(light.ShadowMap.Projection.Mul(GeometryMath.LookAt(pos, pos.Add(GeometryMath.Vector3{0.0, 0.0, 1.0}), GeometryMath.Vector3{0.0, -1.0, 0.0})))
	light.ShadowMap.ViewProjection[5].Set(light.ShadowMap.Projection.Mul(GeometryMath.LookAt(pos, pos.Add(GeometryMath.Vector3{0.0, 0.0, -1.0}), GeometryMath.Vector3{0.0, -1.0, 0.0})))
}

func (light *UBOPointLight) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
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

	type ShadowMap struct {
		Distance    float32                 `yaml:"distance"`
		Shader      Shader.Ptr              `yaml:"shader"`
		FrameBuffer FrameBuffer.FrameBuffer `yaml:"frameBuffer"`
		Shaders     []Shader.Ptr            `yaml:"bindOnShaders"`
	}

	yamlConfig := struct {
		PointLight `yaml:"pointLight"`
		ShadowMap  `yaml:"shadowMap"`
	}{
		PointLight: PointLight{
			Position:  light.Position.Get(),
			Linear:    light.Linear.Get(),
			Quadratic: light.Quadratic.Get(),
			Ambient:   light.Ambient.Get(),
			Diffuse:   light.Diffuse.Get(),
			Specular:  light.Specular.Get(),
		},
		ShadowMap: ShadowMap{
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

	texture, err := NewShadowCubeMap(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width, yamlConfig.ShadowMap.FrameBuffer.Viewport.Height, ShadowMapPointLight)
	if err != nil {
		return err
	}

	yamlConfig.ShadowMap.FrameBuffer.AddDepthAttachment(texture)
	if err := yamlConfig.ShadowMap.FrameBuffer.Finish(); err != nil {
		return err
	}

	light.ShadowMap.Projection = GeometryMath.Perspective(math.Pi/2, float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width)/float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Height), 1, yamlConfig.ShadowMap.Distance)
	light.Position.Set(yamlConfig.PointLight.Position)
	light.Linear.Set(yamlConfig.PointLight.Linear)
	light.Quadratic.Set(yamlConfig.PointLight.Quadratic)
	light.Ambient.Set(yamlConfig.PointLight.Ambient)
	light.Diffuse.Set(yamlConfig.PointLight.Diffuse)
	light.Specular.Set(yamlConfig.PointLight.Specular)
	light.ShadowMap.Distance.Set(yamlConfig.ShadowMap.Distance)
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
