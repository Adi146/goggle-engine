package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer/MemoryLayout"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"gopkg.in/yaml.v3"
)

const (
	SpotLight_ubo_binding = 3
	SpotLight_fbo_type    = "shadowMap_spotLight"
)

var (
	SpotLightBuffer Buffer.UniformBuffer
	SpotLightArray  UniformBufferArray
)

type SpotLight struct {
	SpotLight struct {
		Position  GeometryMath.Vector3 `yaml:"position,flow"`
		Linear    float32              `yaml:"linear,flow"`
		Quadratic float32              `yaml:"quadratic,flow"`
		Ambient   GeometryMath.Vector3 `yaml:"ambient,flow"`
		Diffuse   GeometryMath.Vector3 `yaml:"diffuse,flow"`
		Specular  GeometryMath.Vector3 `yaml:"specular,flow"`
		Direction GeometryMath.Vector3 `yaml:"direction,flow"`
		InnerCone float32              `yaml:"innerCone"`
		OuterCone float32              `yaml:"outerCone"`
	} `yaml:"spotLight"`

	ShadowMap ShadowMapping.ShadowMap `yaml:"shadowMap"`
	Buffer.DynamicBufferData
}

func (light *SpotLight) Update(position GeometryMath.Vector3, front GeometryMath.Vector3, up GeometryMath.Vector3) {
	light.SpotLight.Position = position
	light.SpotLight.Direction = front
	light.ShadowMap.Camera.Update(position, front, up)
	light.SetIsSync(false)
}

func (light *SpotLight) UpdateCamera(scene Scene.IScene, camera Camera.ICamera) {
	SpotLightBuffer.Sync()
}

func (light *SpotLight) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene, camera Camera.ICamera) error {
	return light.ShadowMap.Draw(shader, invoker, scene, camera)
}

func (light *SpotLight) GetBufferData() interface{} {
	return struct {
		position             GeometryMath.Vector3
		linear               float32
		quadratic            MemoryLayout.Std140Float32
		ambient              MemoryLayout.Std140Vector3
		diffuse              MemoryLayout.Std140Vector3
		specular             MemoryLayout.Std140Vector3
		direction            GeometryMath.Vector3
		innerCone            float32
		outerCone            MemoryLayout.Std140Float32
		viewProjectionMatrix GeometryMath.Matrix4x4
	}{
		position:             light.SpotLight.Position,
		linear:               light.SpotLight.Linear,
		quadratic:            MemoryLayout.Std140Float32{Float32: light.SpotLight.Quadratic},
		ambient:              MemoryLayout.Std140Vector3{Vector3: light.SpotLight.Ambient},
		diffuse:              MemoryLayout.Std140Vector3{Vector3: light.SpotLight.Diffuse},
		specular:             MemoryLayout.Std140Vector3{Vector3: light.SpotLight.Specular},
		direction:            light.SpotLight.Direction,
		innerCone:            light.SpotLight.InnerCone,
		outerCone:            MemoryLayout.Std140Float32{Float32: light.SpotLight.OuterCone},
		viewProjectionMatrix: light.ShadowMap.Camera.GetProjectionMatrix().Mul(light.ShadowMap.Camera.GetViewMatrix()),
	}
}

func (light *SpotLight) UnmarshalYAML(value *yaml.Node) error {
	if SpotLightBuffer == (Buffer.UniformBuffer{}) {
		SpotLightBuffer = Buffer.NewUniformBuffer(&SpotLightArray, SpotLight_ubo_binding)
	}

	light.ShadowMap.Camera = &Camera.Camera{}
	light.ShadowMap.TextureType = ShadowMapping.ShadowMapSpotLight
	light.ShadowMap.TextureConstructor = ShadowMapping.NewShadowMapTexture
	light.ShadowMap.FrameBuffer.Type = SpotLight_fbo_type
	light.ShadowMap.UpdateCameraCallback = light.UpdateCamera

	light.ShadowMap.LightIndex = SpotLightArray.Add(light)

	type yamlConfigType SpotLight
	yamlConfig := (*yamlConfigType)(light)

	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	light.SpotLight.InnerCone = GeometryMath.Cos(GeometryMath.Radians(light.SpotLight.InnerCone))
	light.SpotLight.OuterCone = GeometryMath.Cos(GeometryMath.Radians(light.SpotLight.OuterCone))

	light.ShadowMap.Camera.SetProjection(&GeometryMath.PerspectiveConfig{
		Fovy:   GeometryMath.Degree(light.SpotLight.OuterCone),
		Aspect: float32(light.ShadowMap.FrameBuffer.Viewport.Width) / float32(light.ShadowMap.FrameBuffer.Viewport.Height),
		Near:   0.1,
		Far:    light.ShadowMap.Distance,
	})

	return nil
}
