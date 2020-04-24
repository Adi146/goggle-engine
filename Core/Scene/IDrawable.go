package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

type IDrawable interface {
	Draw(shader Shader.IShaderProgram, invoker IDrawable, scene IScene, camera Camera.ICamera) error
}
