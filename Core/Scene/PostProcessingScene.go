package Scene

import (
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Model"
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

type PostProcessingScene struct {
	SceneBase
	quad *Model.Mesh
}

func (scene *PostProcessingScene) Init() error {
	quad, err := Model.NewMesh(quadVertices, Buffer.RegisterVertexBufferAttributes, quadIndices)
	if err != nil {
		return err
	}
	scene.quad = quad

	return nil
}

func (scene *PostProcessingScene) Tick(timeDelta float32) {
}

func (scene *PostProcessingScene) Draw() {
	scene.quad.Draw()
}
