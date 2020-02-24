package PostProcessing

import (
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Model"
	sceneCore "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

var (
	quadVertices = []Buffer.Vertex{
		{
			Position: Vector.Vector3{-1.0, 1.0, 0.0},
			UV:       Vector.Vector2{0.0, 1.0},
		},
		{
			Position: Vector.Vector3{-1.0, -1.0, 0.0},
			UV:       Vector.Vector2{0.0, 0.0},
		},
		{
			Position: Vector.Vector3{1.0, -1.0, 0.0},
			UV:       Vector.Vector2{1.0, 0.0},
		},
		{
			Position: Vector.Vector3{1.0, 1.0, 0.0},
			UV:       Vector.Vector2{1.0, 1.0},
		},
	}
	quadIndices = []uint32{
		0, 1, 2,
		0, 2, 3,
	}
)

type Scene struct {
	sceneCore.SceneBase
	quad *Model.Mesh

	Kernel *Kernel `yaml:",inline"`
	Shader Shader.IShaderProgram
}

func (scene *Scene) Init() error {
	quad, err := Model.NewMesh(quadVertices, Buffer.RegisterVertexBufferAttributes, quadIndices)
	if err != nil {
		return err
	}
	scene.quad = quad

	return nil
}

func (scene *Scene) Tick(timeDelta float32) {
}

func (scene *Scene) Draw() {
	scene.Shader.Bind()
	scene.Shader.BindObject(scene.Kernel)
	scene.quad.Draw()
}

func (scene *Scene) AddResult(result *Texture.Texture) {
	scene.Shader.Bind()
	scene.Shader.BindObject(result)
}
