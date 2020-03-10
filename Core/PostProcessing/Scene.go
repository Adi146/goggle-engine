package PostProcessing

import (
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Model"
	sceneCore "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

var (
	quadVertices = []Buffer.Vertex{
		{
			Position: GeometryMath.Vector3{-1.0, 1.0, 0.0},
			UV:       GeometryMath.Vector2{0.0, 1.0},
		},
		{
			Position: GeometryMath.Vector3{-1.0, -1.0, 0.0},
			UV:       GeometryMath.Vector2{0.0, 0.0},
		},
		{
			Position: GeometryMath.Vector3{1.0, -1.0, 0.0},
			UV:       GeometryMath.Vector2{1.0, 0.0},
		},
		{
			Position: GeometryMath.Vector3{1.0, 1.0, 0.0},
			UV:       GeometryMath.Vector2{1.0, 1.0},
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
	scene.SceneBase.Tick(timeDelta)
	scene.OpaqueObjects = append(scene.OpaqueObjects, scene.quad)
}

func (scene *Scene) Draw(shader Shader.IShaderProgram) {
	if shader == nil {
		shader = scene.Shader
	}

	shader.BindObject(scene.Kernel)

	scene.SceneBase.Draw(shader)
}
