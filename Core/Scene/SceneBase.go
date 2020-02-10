package Scene

import (
	"sort"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type SceneBase struct {
	activeShaderProgram Shader.IShaderProgram

	keyboardInput Window.IKeyboardInput
	mouseInput    Window.IMouseInput

	CameraPosition     *Vector.Vector3
	PreRenderObjects   []IDrawable
	OpaqueObjects      []IDrawable
	TransparentObjects []IDrawable
}

func (scene *SceneBase) Init() error {
	return nil
}

func (scene *SceneBase) SetActiveShaderProgram(shaderProgram Shader.IShaderProgram) {
	scene.activeShaderProgram = shaderProgram
}

func (scene *SceneBase) GetActiveShaderProgram() Shader.IShaderProgram {
	return scene.activeShaderProgram
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

func (scene *SceneBase) Draw() error {
	var err Error.ErrorCollection

	err.Push(scene.drawPreRenderObjects())
	err.Push(scene.drawOpaqueObjects())
	err.Push(scene.drawTransparentObjects())

	return err.Err()
}

func (scene *SceneBase) drawPreRenderObjects() error {
	var err Error.ErrorCollection

	for _, drawable := range scene.PreRenderObjects {
		err.Push(drawable.Draw(scene.GetActiveShaderProgram()))
	}
	scene.PreRenderObjects = []IDrawable{}

	return err.Err()
}

func (scene *SceneBase) drawOpaqueObjects() error {
	var err Error.ErrorCollection

	for _, drawable := range scene.OpaqueObjects {
		err.Push(drawable.Draw(scene.GetActiveShaderProgram()))
	}
	scene.OpaqueObjects = []IDrawable{}

	return err.Err()
}

func (scene *SceneBase) drawTransparentObjects() error {
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
		err.Push(drawable.Draw(scene.GetActiveShaderProgram()))
	}

	scene.TransparentObjects = []IDrawable{}

	return err.Err()
}
