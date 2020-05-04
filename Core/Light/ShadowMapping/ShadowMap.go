package ShadowMapping

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	core "github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
)

type ShadowMap struct {
	Camera               Camera.ICamera
	UpdateCameraCallback func(scene Scene.IScene, camera Camera.ICamera)

	Shader      Shader.IShaderProgram
	FrameBuffer FrameBuffer.FrameBuffer
	LightIndex  int

	Distance           float32
	TransitionDistance float32

	TextureConstructor func(width int32, height int32, textureType core.Type) (*core.Texture, error)
	TextureType        core.Type
}

func (shadowMap *ShadowMap) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene, camera Camera.ICamera) error {
	if _, isShadowMap := invoker.(*ShadowMap); isShadowMap {
		return nil
	}

	if shadowMap.UpdateCameraCallback != nil {
		shadowMap.UpdateCameraCallback(scene, camera)
	}

	defer FrameBuffer.GetCurrentFrameBuffer().Bind()
	defer Function.GetCurrentCullFunction().Set()
	defer Function.GetCurrentDepthFunction().Set()
	defer Function.GetCurrentBlendFunction().Set()

	shadowMap.FrameBuffer.Bind()
	Function.Front.Set()
	Function.Less.Set()
	Function.DisabledBlend.Set()

	shadowMap.FrameBuffer.Clear()

	if shader != nil {
		defer shader.Bind()
	}

	if err := shadowMap.Shader.BindObject(int32(shadowMap.LightIndex)); err != nil {
		return err
	}

	return scene.Draw(shadowMap.Shader, shadowMap, scene, shadowMap.Camera)
}

func (shadowMap *ShadowMap) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		Shader             Shader.Ptr               `yaml:"shader"`
		FrameBuffer        *FrameBuffer.FrameBuffer `yaml:"frameBuffer"`
		Shaders            []Shader.Ptr             `yaml:"bindOnShaders"`
		Distance           *float32                 `yaml:"distance"`
		TransitionDistance *float32                 `yaml:"transitionDistance"`
	}{
		Shader: Shader.Ptr{
			IShaderProgram: shadowMap.Shader,
		},
		FrameBuffer:        &shadowMap.FrameBuffer,
		Distance:           &shadowMap.Distance,
		TransitionDistance: &shadowMap.TransitionDistance,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	texture, err := shadowMap.TextureConstructor(yamlConfig.FrameBuffer.Viewport.Width, yamlConfig.FrameBuffer.Viewport.Height, shadowMap.TextureType)
	if err != nil {
		return err
	}
	yamlConfig.FrameBuffer.AddDepthAttachment(texture)
	if err := yamlConfig.FrameBuffer.Finish(); err != nil {
		return err
	}

	shadowMap.Shader = yamlConfig.Shader.IShaderProgram

	for _, shader := range yamlConfig.Shaders {
		uniformAddress, err := shader.GetUniformAddress(texture)
		if err != nil {
			return err
		}
		if err := shader.BindUniform(texture, fmt.Sprintf(uniformAddress, shadowMap.LightIndex)); err != nil {
			Log.Error(err, "")
		}
	}

	return nil
}
