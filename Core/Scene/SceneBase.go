package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Window"
)

type SceneBase struct {
	activeShaderProgram Shader.IShaderProgram

	keyboardInput Window.IKeyboardInput
	mouseInput    Window.IMouseInput
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