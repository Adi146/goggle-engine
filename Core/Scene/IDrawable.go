package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type IDrawable interface {
	Draw(shader Shader.IShaderProgram, invoker IDrawable, scene IScene) error
}
