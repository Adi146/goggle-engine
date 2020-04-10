package Model

import (
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/VertexBuffer"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Mesh struct {
	VertexBuffer *VertexBuffer.VertexBuffer
	IndexBuffer  *VertexBuffer.IndexBuffer
}

func NewMesh(vertices []VertexBuffer.Vertex, indices []uint32) (*Mesh, error) {
	vertexBuffer, err := VertexBuffer.NewVertexBuffer(vertices)
	if err != nil {
		return nil, err
	}

	return &Mesh{
		VertexBuffer: vertexBuffer,
		IndexBuffer:  VertexBuffer.NewIndexBuffer(indices),
	}, nil
}

func (mesh *Mesh) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	var err Error.ErrorCollection

	err.Push(shader.BindObject(mesh.VertexBuffer))
	mesh.IndexBuffer.Bind()
	gl.DrawElements(gl.TRIANGLES, mesh.IndexBuffer.Length, gl.UNSIGNED_INT, nil)
	mesh.IndexBuffer.Unbind()
	mesh.VertexBuffer.Unbind()

	return err.Err()
}
