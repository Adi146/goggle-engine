package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	shadowMap "github.com/Adi146/goggle-engine/Core/Light/ShadowMapping/DirectionalLight"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer/MemoryLayout"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"gopkg.in/yaml.v3"
)

const (
	DirectionalLight_ubo_binding = 1
	DirectionalLight_fbo_type    = "shadowMap_directionalLight"
)

var (
	DirectionalLightArray  UniformBufferArray
	DirectionalLightBuffer Buffer.UniformBuffer
)

type DirectionalLight struct {
	DirectionalLight struct {
		Direction GeometryMath.Vector3 `yaml:"direction,flow"`
		Ambient   GeometryMath.Vector3 `yaml:"ambient,flow"`
		Diffuse   GeometryMath.Vector3 `yaml:"diffuse,flow"`
		Specular  GeometryMath.Vector3 `yaml:"specular,flow"`
	} `yaml:"directionalLight"`

	ShadowMap ShadowMapping.ShadowMap `yaml:"shadowMap"`
	Buffer.DynamicBufferData
}

func (light *DirectionalLight) Update(direction GeometryMath.Vector3) {
	light.DirectionalLight.Direction = direction
	light.SetIsSync(false)
}

func (light *DirectionalLight) UpdateCamera(scene Scene.IScene, camera Camera.ICamera) {
	light.ShadowMap.Camera.(*shadowMap.Camera).UpdateInternal(camera, light.DirectionalLight.Direction, light.ShadowMap.Distance)
	DirectionalLightBuffer.Sync()
}

func (light *DirectionalLight) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene, camera Camera.ICamera) error {
	return light.ShadowMap.Draw(shader, invoker, scene, camera)
}

func (light *DirectionalLight) GetBufferData() interface{} {
	return struct {
		direction            MemoryLayout.Std140Vector3
		ambient              MemoryLayout.Std140Vector3
		diffuse              MemoryLayout.Std140Vector3
		specular             MemoryLayout.Std140Vector3
		viewProjectionMatrix GeometryMath.Matrix4x4
		distance             float32
		transitionDistance   float32
		padding              [2]MemoryLayout.Padding
	}{
		direction:            MemoryLayout.Std140Vector3{Vector3: light.DirectionalLight.Direction},
		ambient:              MemoryLayout.Std140Vector3{Vector3: light.DirectionalLight.Ambient},
		diffuse:              MemoryLayout.Std140Vector3{Vector3: light.DirectionalLight.Diffuse},
		specular:             MemoryLayout.Std140Vector3{Vector3: light.DirectionalLight.Specular},
		viewProjectionMatrix: light.ShadowMap.Camera.GetProjectionMatrix().Mul(light.ShadowMap.Camera.GetViewMatrix()),
		distance:             light.ShadowMap.Distance,
		transitionDistance:   light.ShadowMap.TransitionDistance,
	}
}

func (light *DirectionalLight) UnmarshalYAML(value *yaml.Node) error {
	if DirectionalLightBuffer == (Buffer.UniformBuffer{}) {
		DirectionalLightBuffer = Buffer.NewUniformBuffer(&DirectionalLightArray, DirectionalLight_ubo_binding)
	}

	light.ShadowMap.Camera = &shadowMap.Camera{}
	light.ShadowMap.TextureType = ShadowMapping.ShadowMapDirectionalLight
	light.ShadowMap.TextureConstructor = ShadowMapping.NewShadowMapTexture
	light.ShadowMap.FrameBuffer.Type = DirectionalLight_fbo_type
	light.ShadowMap.UpdateCameraCallback = light.UpdateCamera

	light.ShadowMap.LightIndex = DirectionalLightArray.Add(light)

	type yamlConfigType DirectionalLight
	yamlConfig := (*yamlConfigType)(light)

	return value.Decode(&yamlConfig)
}
