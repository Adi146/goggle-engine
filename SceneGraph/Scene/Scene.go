package Scene

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/RenderTarget"
	"github.com/Adi146/goggle-engine/Utils"
)

type Scene struct {
	*RenderTarget.OpenGLRenderTarget
	Root IParentNode
}

func NewScene(renderTarget *RenderTarget.OpenGLRenderTarget) *Scene {
	return &Scene{
		OpenGLRenderTarget: renderTarget,
		Root:               nil,
	}
}

func (scene *Scene) Draw(timeDelta float32) {
	scene.OpenGLRenderTarget.Clear(&Vector.Vector4{0, 0, 0, 1})
	var err Utils.ErrorCollection

	if scene.Root != nil {
		err.Push(scene.Root.TickChildren(timeDelta))
		err.Push(scene.GetActiveShaderProgram().BeginDraw())
		err.Push(scene.Root.DrawChildren())
		scene.GetActiveShaderProgram().EndDraw()
	}
}

func (scene *Scene) SetRoot(node IParentNode) {
	node.setScene(scene)
	scene.Root = node
}
