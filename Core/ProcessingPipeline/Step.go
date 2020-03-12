package ProcessingPipeline

import (
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type Step struct {
	FrameBuffer    FrameBuffer.IFrameBuffer
	Scene          Scene.IScene
	EnforcedShader Shader.IShaderProgram
}

func (step *Step) Execute() {
	step.FrameBuffer.Bind()
	step.FrameBuffer.Clear()
	step.Scene.Draw(step.EnforcedShader, nil, nil)
	step.FrameBuffer.Unbind()
}
