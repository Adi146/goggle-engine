package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
	"sort"

	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type SceneBase struct {
	Window Window.SDLWindow `yaml:"window"`
	Camera Camera.UBOCamera `yaml:"camera"`

	CullFunction  Function.CullFunction  `yaml:"culling"`
	DepthFunction Function.DepthFunction `yaml:"blend"`
	BlendFunction Function.BlendFunction `yaml:"depthTest"`

	preRenderObjects   []IDrawable
	opaqueObjects      []IDrawable
	transparentObjects []ITransparentDrawable
}

func (scene *SceneBase) Tick(timeDelta float32) {
	scene.Clear()
}

func (scene *SceneBase) AddPreRenderObject(obj IDrawable) {
	scene.preRenderObjects = append(scene.preRenderObjects, obj)
}

func (scene *SceneBase) AddOpaqueObject(obj IDrawable) {
	scene.opaqueObjects = append(scene.opaqueObjects, obj)
}

func (scene *SceneBase) AddTransparentObject(obj ITransparentDrawable) {
	scene.transparentObjects = append(scene.transparentObjects, obj)
}

func (scene *SceneBase) Clear() {
	scene.preRenderObjects = []IDrawable{}
	scene.opaqueObjects = []IDrawable{}
	scene.transparentObjects = []ITransparentDrawable{}
}

func (scene *SceneBase) Draw(shader Shader.IShaderProgram, invoker IDrawable, origin IScene, camera Camera.ICamera) error {
	if shader != nil {
		shader.Bind()
		defer shader.Unbind()
	}
	if invoker == nil {
		invoker = scene
	}
	if origin == nil {
		origin = scene
	}
	if camera == nil {
		camera = scene.GetCamera()
	}

	if invoker == scene {
		scene.Window.Bind()
		scene.Window.Clear()

		scene.CullFunction.Set()
		scene.DepthFunction.Set()
		scene.BlendFunction.Set()
	}

	var err Error.ErrorCollection
	err.Push(scene.drawPreRenderObjects(shader, invoker, origin, camera))
	err.Push(scene.drawOpaqueObjects(shader, invoker, origin, camera))
	err.Push(scene.drawTransparentObjects(shader, invoker, origin, camera))

	if err.Len() != 0 {
		Log.Error(&err, "render error")
	}

	return err.Err()
}

func (scene *SceneBase) drawPreRenderObjects(shader Shader.IShaderProgram, invoker IDrawable, origin IScene, camera Camera.ICamera) error {
	var err Error.ErrorCollection

	for i := 0; i < len(scene.preRenderObjects); i++ {
		err.Push(scene.preRenderObjects[i].Draw(shader, invoker, origin, camera))
	}

	return err.Err()
}

func (scene *SceneBase) drawOpaqueObjects(shader Shader.IShaderProgram, invoker IDrawable, origin IScene, camera Camera.ICamera) error {
	var err Error.ErrorCollection

	for i := 0; i < len(scene.opaqueObjects); i++ {
		err.Push(scene.opaqueObjects[i].Draw(shader, invoker, origin, camera))
	}

	return err.Err()
}

func (scene *SceneBase) drawTransparentObjects(shader Shader.IShaderProgram, invoker IDrawable, origin IScene, camera Camera.ICamera) error {
	var err Error.ErrorCollection
	cameraPosition := scene.Camera.GetPosition()

	transparentDrawables := make([]transparentObject, len(scene.transparentObjects))
	for i, drawable := range scene.transparentObjects {
		transparentDrawables[i] = transparentObject{
			IDrawable:      drawable,
			CameraDistance: cameraPosition.Sub(drawable.GetPosition()).Length(),
		}
	}
	sort.Sort(byDistance(transparentDrawables))
	for _, drawable := range transparentDrawables {
		err.Push(drawable.Draw(shader, invoker, origin, camera))
	}

	return err.Err()
}

func (scene *SceneBase) GetWindow() Window.IWindow {
	return &scene.Window
}

func (scene *SceneBase) GetCamera() Camera.ICamera {
	return &scene.Camera
}

func (scene *SceneBase) UnmarshalYAML(value *yaml.Node) error {
	type yamlConfigType SceneBase
	yamlConfig := (yamlConfigType)(*scene)
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	width, height := yamlConfig.Window.GetSize()

	projectionConfig := yamlConfig.Camera.GetProjection()
	projectionConfig.Aspect = float32(width) / float32(height)
	yamlConfig.Camera.SetProjection(projectionConfig)

	*scene = (SceneBase)(yamlConfig)

	return nil
}
