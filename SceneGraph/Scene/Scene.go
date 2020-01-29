package Scene

import (
	"github.com/Adi146/goggle-engine/Core/RenderTarget"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"github.com/Adi146/goggle-engine/Utils/Log"
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

func (scene *Scene) Tick(timeDelta float32) {
	if scene.Root != nil {
		Log.Error(Error.NewErrorWithFields(scene.Root.TickChildren(timeDelta), scene.Root.GetLogFields()), "tick error")
	}
}

func (scene *Scene) Draw() {
	if scene.Root != nil {
		Log.Error(scene.GetFrameBuffer().GetShaderProgram().BeginDraw(), "begin draw error")
		Log.Error(Error.NewErrorWithFields(scene.Root.DrawChildren(), scene.Root.GetLogFields()), "render error")
		scene.GetFrameBuffer().GetShaderProgram().EndDraw()
	}
}

func (scene *Scene) SetRoot(node IParentNode) {
	node.setScene(scene)
	scene.Root = node
}
