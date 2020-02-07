package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type SceneBase struct {
	activeShaderProgram Shader.IShaderProgram

	keyboardInput Window.IKeyboardInput
	mouseInput    Window.IMouseInput

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

	for _, drawable := range scene.PreRenderObjects {
		err.Push(drawable.Draw(scene.GetActiveShaderProgram()))
	}
	scene.PreRenderObjects = []IDrawable{}
	for _, drawable := range scene.OpaqueObjects {
		err.Push(drawable.Draw(scene.GetActiveShaderProgram()))
	}
	scene.OpaqueObjects = []IDrawable{}
	for _, drawable := range scene.TransparentObjects {
		err.Push(drawable.Draw(scene.GetActiveShaderProgram()))
	}
	scene.TransparentObjects = []IDrawable{}

	return err.Err()
}
