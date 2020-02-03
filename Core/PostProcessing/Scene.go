package PostProcessing

import (
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Model"
	sceneCore "github.com/Adi146/goggle-engine/Core/Scene"
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
	if shader := scene.GetActiveShaderProgram(); shader != nil {
		if scene.Kernel != nil {
			shader.BindObject(scene.Kernel)
		}
		scene.quad.Draw()
	}
}