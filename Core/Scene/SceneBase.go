package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Texture"
	"sort"

	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

type SceneBase struct {
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

func (scene *SceneBase) AddResult(texture *Texture.Texture) {
}

func (scene *SceneBase) drawPreRenderObjects() error {
	var err Error.ErrorCollection

	for _, drawable := range scene.PreRenderObjects {
		err.Push(drawable.Draw())
	}
	scene.PreRenderObjects = []IDrawable{}

	return err.Err()
}

func (scene *SceneBase) drawOpaqueObjects() error {
	var err Error.ErrorCollection

	for _, drawable := range scene.OpaqueObjects {
		err.Push(drawable.Draw())
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
		err.Push(drawable.Draw())
	}

	scene.TransparentObjects = []IDrawable{}

	return err.Err()
}
