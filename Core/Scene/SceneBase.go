package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"sort"

	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type SceneBase struct {
	keyboardInput Window.IKeyboardInput
	mouseInput    Window.IMouseInput

	CameraPosition     *GeometryMath.Vector3
	PreRenderObjects   []IDrawable
	OpaqueObjects      []IDrawable
	TransparentObjects []IDrawable
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

func (scene *SceneBase) Draw(step *ProcessingPipelineStep) error {
	var err Error.ErrorCollection

	err.Push(scene.drawPreRenderObjects(step))
	err.Push(scene.drawOpaqueObjects(step))
	err.Push(scene.drawTransparentObjects(step))

	return err.Err()
}

func (scene *SceneBase) drawPreRenderObjects(step *ProcessingPipelineStep) error {
	var err Error.ErrorCollection

	for _, drawable := range scene.PreRenderObjects {
		err.Push(drawable.Draw(step))
	}
	scene.PreRenderObjects = []IDrawable{}

	return err.Err()
}

func (scene *SceneBase) drawOpaqueObjects(step *ProcessingPipelineStep) error {
	var err Error.ErrorCollection

	for _, drawable := range scene.OpaqueObjects {
		err.Push(drawable.Draw(step))
	}
	scene.OpaqueObjects = []IDrawable{}

	return err.Err()
}

func (scene *SceneBase) drawTransparentObjects(step *ProcessingPipelineStep) error {
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
		err.Push(drawable.Draw(step))
	}

	scene.TransparentObjects = []IDrawable{}

	return err.Err()
}
