package Light

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/Camera"
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
	directionalLight_offset_direction          = 0
	directionalLight_offset_ambient            = 16
	directionalLight_offset_diffuse            = 32
	directionalLight_offset_specular           = 48
	directionalLight_offset_viewProjection     = 64
	directionalLight_offset_distance           = 128
	directionalLight_offset_transitionDistance = 132

	directionalLight_size_section = 144
	directionalLight_ubo_size     = directionalLight_size_section
	DirectionalLight_ubo_type     = "directionalLight"

	DirectionalLight_fbo_type = "shadowMap_directionalLight"

	near_plane = 0.1
	offset     = 15
)

type UBODirectionalLight struct {
	Direction UniformBufferSection.Vector3
	Ambient   UniformBufferSection.Vector3
	Diffuse   UniformBufferSection.Vector3
	Specular  UniformBufferSection.Vector3
	ShadowMap struct {
		Projection         GeometryMath.Matrix4x4
		ViewMatrix         GeometryMath.Matrix4x4
		ViewProjection     UniformBufferSection.Matrix4x4
		Distance           UniformBufferSection.Float
		TransitionDistance UniformBufferSection.Float
		Shader             Shader.IShaderProgram
		FrameBuffer        FrameBuffer.FrameBuffer
	}
}

func (light *UBODirectionalLight) ForceUpdate() {
	light.Direction.ForceUpdate()
	light.Ambient.ForceUpdate()
	light.Diffuse.ForceUpdate()
	light.Specular.ForceUpdate()
	light.ShadowMap.ViewProjection.ForceUpdate()
	light.ShadowMap.Distance.ForceUpdate()
	light.ShadowMap.TransitionDistance.ForceUpdate()
}

func (light *UBODirectionalLight) SetUniformBuffer(ubo UniformBuffer.IUniformBuffer, offset int) {
	light.Direction.SetUniformBuffer(ubo, offset+directionalLight_offset_direction)
	light.Ambient.SetUniformBuffer(ubo, offset+directionalLight_offset_ambient)
	light.Diffuse.SetUniformBuffer(ubo, offset+directionalLight_offset_diffuse)
	light.Specular.SetUniformBuffer(ubo, offset+directionalLight_offset_specular)
	light.ShadowMap.ViewProjection.SetUniformBuffer(ubo, offset+directionalLight_offset_viewProjection)
	light.ShadowMap.Distance.SetUniformBuffer(ubo, offset+directionalLight_offset_distance)
	light.ShadowMap.TransitionDistance.SetUniformBuffer(ubo, offset+directionalLight_offset_transitionDistance)
}

func (light *UBODirectionalLight) GetSize() int {
	return directionalLight_size_section
}

func (light *UBODirectionalLight) SetDirection(val GeometryMath.Vector3) {
	light.Direction.Set(val)
}

func (light *UBODirectionalLight) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	_, isPointLight := invoker.(*UBOPointLight)
	_, isDirectionalLight := invoker.(*UBODirectionalLight)
	_, isSpotLight := invoker.(*UBOSpotLight)
	if isPointLight || isDirectionalLight || isSpotLight {
		return nil
	}

	light.updateShadowCamera(scene)

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

func (light *UBODirectionalLight) updateShadowCamera(scene Scene.IScene) {
	boundingBox, center := light.calcCameraFrustumBoundingBox(scene)
	direction := light.Direction.Get()

	light.ShadowMap.Projection = GeometryMath.Orthographic(boundingBox.Min.X(), boundingBox.Max.X(), boundingBox.Min.Y(), boundingBox.Max.Y(), boundingBox.Min.Z(), boundingBox.Max.Z())
	light.ShadowMap.ViewMatrix = GeometryMath.LookAt(center.Add(direction.Invert()), center, GeometryMath.Vector3{0, 1, 0})
	light.ShadowMap.ViewProjection.Set(light.ShadowMap.Projection.Mul(light.ShadowMap.ViewMatrix))
}

func (light *UBODirectionalLight) calcCameraFrustumBoundingBox(scene Scene.IScene) (BoundingVolume.AABB, GeometryMath.Vector3) {
	direction := light.Direction.Get()
	tmpViewMatrix := GeometryMath.LookAt(direction.Invert(), GeometryMath.Vector3{0, 0, 0}, GeometryMath.Vector3{0, 1, 0})

	frustumPoints := light.calcCameraFrustumPoints(scene.GetCamera())
	for i := range frustumPoints {
		frustumPoints[i] = tmpViewMatrix.MulVector(frustumPoints[i])
	}
	boundingBox := BoundingVolume.NewAABB(frustumPoints[:])
	boundingBox.Max[2] += offset

	return boundingBox, tmpViewMatrix.Inverse().MulVector(boundingBox.GetCenter())
}

func (light *UBODirectionalLight) calcCameraFrustumPoints(camera Camera.ICamera) [8]GeometryMath.Vector3 {
	projectionConfig := camera.GetProjection()
	position := camera.GetPosition()

	farHeight := light.ShadowMap.Distance.Get() * GeometryMath.Tan(GeometryMath.Radians(projectionConfig.Fovy*0.5))
	nearHeight := float32(near_plane) * GeometryMath.Tan(projectionConfig.Fovy*0.5)
	farWidth := farHeight * projectionConfig.Aspect
	nearWidth := nearHeight * projectionConfig.Aspect

	front := camera.GetFront()
	up := camera.GetUp()
	right := front.Cross(up)
	down := up.Invert()
	left := right.Invert()

	centerFar := position.Add(front.MulScalar(light.ShadowMap.Distance.Get()))
	centerNear := position.Add(front.MulScalar(near_plane))

	farTop := centerFar.Add(up.MulScalar(farHeight))
	farBottom := centerFar.Add(down.MulScalar(farHeight))
	nearTop := centerNear.Add(up.MulScalar(nearHeight))
	nearBottom := centerNear.Add(down.MulScalar(nearHeight))

	return [8]GeometryMath.Vector3{
		farTop.Add(right.MulScalar(farWidth)),
		farTop.Add(left.MulScalar(farWidth)),
		farBottom.Add(right.MulScalar(farWidth)),
		farBottom.Add(left.MulScalar(farWidth)),
		nearTop.Add(right.MulScalar(nearWidth)),
		nearTop.Add(left.MulScalar(nearWidth)),
		nearBottom.Add(right.MulScalar(nearWidth)),
		nearBottom.Add(left.MulScalar(nearWidth)),
	}
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

	type ShadowMap struct {
		Distance           float32                 `yaml:"distance"`
		TransitionDistance float32                 `yaml:"transitionDistance"`
		Shader             Shader.Ptr              `yaml:"shader"`
		FrameBuffer        FrameBuffer.FrameBuffer `yaml:"frameBuffer"`
		Shaders            []Shader.Ptr            `yaml:"bindOnShaders"`
	}

	yamlConfig := struct {
		DirectionalLight `yaml:"directionalLight"`
		ShadowMap        `yaml:"shadowMap"`
	}{
		ShadowMap: ShadowMap{
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

	texture, err := NewShadowMap(yamlConfig.ShadowMap.FrameBuffer.Viewport.Width, yamlConfig.ShadowMap.FrameBuffer.Viewport.Height, ShadowMapDirectionalLight)
	if err != nil {
		return err
	}
	yamlConfig.ShadowMap.FrameBuffer.AddDepthAttachment(texture)
	if err := yamlConfig.ShadowMap.FrameBuffer.Finish(); err != nil {
		return err
	}

	light.Direction.Set(yamlConfig.DirectionalLight.Direction)
	light.Ambient.Set(yamlConfig.DirectionalLight.Ambient)
	light.Diffuse.Set(yamlConfig.DirectionalLight.Diffuse)
	light.Specular.Set(yamlConfig.DirectionalLight.Specular)
	light.ShadowMap.Distance.Set(yamlConfig.ShadowMap.Distance)
	light.ShadowMap.TransitionDistance.Set(yamlConfig.ShadowMap.TransitionDistance)
	light.ShadowMap.Shader = yamlConfig.ShadowMap.Shader.IShaderProgram
	light.ShadowMap.FrameBuffer = yamlConfig.ShadowMap.FrameBuffer

	light.SetUniformBuffer(uboYamlConfig.UniformBuffer, 0)

	for _, shader := range yamlConfig.Shaders {
		if err := shader.BindObject(texture); err != nil {
			Log.Error(err, "")
		}
	}

	return nil
}
