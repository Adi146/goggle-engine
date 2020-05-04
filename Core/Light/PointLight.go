package Light

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	shadowMap "github.com/Adi146/goggle-engine/Core/Light/ShadowMapping/PointLight"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer/MemoryLayout"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"gopkg.in/yaml.v3"
)

const (
	PointLight_ubo_binding = 2
	PointLight_fbo_type    = "shadowMap_pointLight"
)

var (
	PointLightBuffer Buffer.UniformBuffer
	PointLightArray  UniformBufferArray
)

type PointLight struct {
	PointLight struct {
		Position  GeometryMath.Vector3 `yaml:"position,flow"`
		Linear    float32              `yaml:"linear,flow"`
		Quadratic float32              `yaml:"quadratic,flow"`
		Ambient   GeometryMath.Vector3 `yaml:"ambient,flow"`
		Diffuse   GeometryMath.Vector3 `yaml:"diffuse,flow"`
		Specular  GeometryMath.Vector3 `yaml:"specular,flow"`
	} `yaml:"pointLight"`

	ShadowMap ShadowMapping.ShadowMap `yaml:"shadowMap"`
	Buffer.DynamicBufferData
}

type std140PointLight struct {
	position             GeometryMath.Vector3
	linear               float32
	quadratic            MemoryLayout.Std140Float32
	ambient              MemoryLayout.Std140Vector3
	diffuse              MemoryLayout.Std140Vector3
	specular             MemoryLayout.Std140Vector3
	viewProjectionMatrix [6]GeometryMath.Matrix4x4
	distance             float32
	padding              [3]MemoryLayout.Padding
}

func (light *PointLight) Update(position GeometryMath.Vector3) {
	light.PointLight.Position = position
	light.ShadowMap.Camera.(*shadowMap.Camera).Update(position, GeometryMath.Vector3{0, 0, 1}, GeometryMath.Vector3{0, 1, 0})
	light.SetIsSync(false)
}

func (light *PointLight) UpdateCamera(scene Scene.IScene, camera Camera.ICamera) {
	PointLightBuffer.Sync()
}

func (light *PointLight) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene, camera Camera.ICamera) error {
	return light.ShadowMap.Draw(shader, invoker, scene, camera)
}

func (light *PointLight) GetBufferData() interface{} {
	return std140PointLight{
		position:             light.PointLight.Position,
		linear:               light.PointLight.Linear,
		quadratic:            MemoryLayout.Std140Float32{Float32: light.PointLight.Quadratic},
		ambient:              MemoryLayout.Std140Vector3{Vector3: light.PointLight.Ambient},
		diffuse:              MemoryLayout.Std140Vector3{Vector3: light.PointLight.Diffuse},
		specular:             MemoryLayout.Std140Vector3{Vector3: light.PointLight.Specular},
		viewProjectionMatrix: light.ShadowMap.Camera.(*shadowMap.Camera).ViewProjectionMatrices,
		distance:             light.ShadowMap.Distance,
	}
}

func (light *PointLight) UnmarshalYAML(value *yaml.Node) error {
	if PointLightBuffer == (Buffer.UniformBuffer{}) {
		PointLightBuffer = Buffer.NewUniformBuffer(&PointLightArray, PointLight_ubo_binding)
	}

	light.ShadowMap.Camera = &shadowMap.Camera{}
	light.ShadowMap.TextureType = ShadowMapping.ShadowMapPointLight
	light.ShadowMap.TextureConstructor = ShadowMapping.NewShadowCubeMapTexture
	light.ShadowMap.FrameBuffer.Type = PointLight_fbo_type
	light.ShadowMap.UpdateCameraCallback = light.UpdateCamera

	light.ShadowMap.LightIndex = PointLightArray.Add(light)

	type yamlConfigType PointLight
	yamlConfig := (*yamlConfigType)(light)

	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	light.ShadowMap.Camera.SetProjection(&GeometryMath.PerspectiveConfig{
		Fovy:   90,
		Aspect: float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width) / float32(yamlConfig.ShadowMap.FrameBuffer.Viewport.Height),
		Near:   0.1,
		Far:    yamlConfig.ShadowMap.Distance,
	})

	return nil
}
