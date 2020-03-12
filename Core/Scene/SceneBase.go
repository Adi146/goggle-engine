package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"sort"

	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type SceneBase struct {
	keyboardInput Window.IKeyboardInput
	mouseInput    Window.IMouseInput

	CameraPosition     *GeometryMath.Vector3
	PreRenderObjects   []IPreRenderObject
	OpaqueObjects      []IDrawable
	TransparentObjects []ITransparentDrawable
}

func (scene *SceneBase) Init() error {
	return nil
}

func (scene *SceneBase) GetKeyboardInput() Window.IKeyboardInput {
	return scene.keyboardInput
}

func (scene *SceneBase) SetKeyboardInput(input Window.IKeyboardInput) {
	scene.keyboardInput = input
}

func (scene *SceneBase) GetMouseInput() Window.IMouseInput {
	return scene.mouseInput
}

func (scene *SceneBase) SetMouseInput(input Window.IMouseInput) {
	scene.mouseInput = input
}

func (scene *SceneBase) Tick(timeDelta float32) {
	scene.PreRenderObjects = []IPreRenderObject{}
	scene.OpaqueObjects = []IDrawable{}
	scene.TransparentObjects = []ITransparentDrawable{}
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

	sort.Sort(byPriority(scene.PreRenderObjects))
	for _, drawable := range scene.PreRenderObjects {
		err.Push(drawable.Draw(shader, invoker, origin))
	}

	return err.Err()
}

func (scene *SceneBase) drawOpaqueObjects(shader Shader.IShaderProgram, invoker IDrawable, origin IScene) error {
	var err Error.ErrorCollection

	for _, drawable := range scene.OpaqueObjects {
		err.Push(drawable.Draw(shader, invoker, origin))
	}

	return err.Err()
}

func (scene *SceneBase) drawTransparentObjects(shader Shader.IShaderProgram, invoker IDrawable, origin IScene) error {
	var err Error.ErrorCollection

	transparentDrawables := make([]transparentObject, len(scene.TransparentObjects))
	for i, drawable := range scene.TransparentObjects {
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
