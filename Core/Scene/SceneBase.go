package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
	"sort"

	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type SceneBase struct {
	Window        Window.IWindow
	CullFunction  Function.CullFunction
	DepthFunction Function.DepthFunction
	BlendFunction Function.BlendFunction

	CameraPosition     *GeometryMath.Vector3
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

func (scene *SceneBase) Draw(shader Shader.IShaderProgram, invoker IDrawable, origin IScene) error {
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

	if invoker == scene {
		scene.Window.Bind()
		scene.Window.Clear()

		scene.CullFunction.Set()
		scene.DepthFunction.Set()
		scene.BlendFunction.Set()
	}

	var err Error.ErrorCollection
	err.Push(scene.drawPreRenderObjects(shader, invoker, origin))
	err.Push(scene.drawOpaqueObjects(shader, invoker, origin))
	err.Push(scene.drawTransparentObjects(shader, invoker, origin))

	if err.Len() != 0 {
		Log.Error(&err, "render error")
	}

	return err.Err()
}

func (scene *SceneBase) drawPreRenderObjects(shader Shader.IShaderProgram, invoker IDrawable, origin IScene) error {
	var err Error.ErrorCollection

	for i := 0; i < len(scene.preRenderObjects); i++ {
		err.Push(scene.preRenderObjects[i].Draw(shader, invoker, origin))
	}

	return err.Err()
}

func (scene *SceneBase) drawOpaqueObjects(shader Shader.IShaderProgram, invoker IDrawable, origin IScene) error {
	var err Error.ErrorCollection

	for i := 0; i < len(scene.opaqueObjects); i++ {
		err.Push(scene.opaqueObjects[i].Draw(shader, invoker, origin))
	}

	return err.Err()
}

func (scene *SceneBase) drawTransparentObjects(shader Shader.IShaderProgram, invoker IDrawable, origin IScene) error {
	var err Error.ErrorCollection

	transparentDrawables := make([]transparentObject, len(scene.transparentObjects))
	for i, drawable := range scene.transparentObjects {
		transparentDrawables[i] = transparentObject{
			IDrawable:      drawable,
			CameraDistance: scene.CameraPosition.Sub(drawable.GetPosition()).Length(),
		}
	}
	sort.Sort(byDistance(transparentDrawables))
	for _, drawable := range transparentDrawables {
		err.Push(drawable.Draw(shader, invoker, origin))
	}

	return err.Err()
}

func (scene *SceneBase) GetWindow() Window.IWindow {
	return scene.Window
}

func (scene *SceneBase) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		WindowConfig  yaml.Node              `yaml:"window"`
		CullFunction  Function.CullFunction  `yaml:"culling"`
		DepthFunction Function.DepthFunction `yaml:"blend"`
		BlendFunction Function.BlendFunction `yaml:"depthTest"`
	}{
		CullFunction:  scene.CullFunction,
		DepthFunction: scene.DepthFunction,
		BlendFunction: scene.BlendFunction,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	scene.CullFunction = yamlConfig.CullFunction
	scene.DepthFunction = yamlConfig.DepthFunction
	scene.BlendFunction = yamlConfig.BlendFunction

	if scene.Window == nil {
		scene.Window = &Window.SDLWindow{}
	}
	if err := yamlConfig.WindowConfig.Decode(scene.Window); err != nil {
		return err
	}

	return nil
}
