package Model

import (
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/VertexBuffer"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Mesh struct {
	vertexBuffer *VertexBuffer.VertexBuffer
	indexBuffer  *VertexBuffer.IndexBuffer
}

func NewMesh(vertices []VertexBuffer.Vertex, vertexBufferAttribFunc func(), indices []uint32) (*Mesh, error) {
	vertexBuffer, err := VertexBuffer.NewVertexBuffer(vertices, vertexBufferAttribFunc)
	if err != nil {
		return nil, err
	}

	return &Mesh{
		vertexBuffer: vertexBuffer,
		indexBuffer:  VertexBuffer.NewIndexBuffer(indices),
	}, nil
}

func (mesh *Mesh) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	mesh.vertexBuffer.Bind()
	mesh.indexBuffer.Bind()
	gl.DrawElements(gl.TRIANGLES, mesh.indexBuffer.Length, gl.UNSIGNED_INT, nil)
	mesh.indexBuffer.Unbind()
	mesh.vertexBuffer.Unbind()

	return nil
}
