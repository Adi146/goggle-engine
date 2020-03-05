package Model

import (
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type Mesh struct {
	vertexBuffer *Buffer.VertexBuffer
	indexBuffer  *Buffer.IndexBuffer
}

func NewMesh(vertices []Buffer.Vertex, vertexBufferAttribFunc func(), indices []uint32) (*Mesh, error) {
	vertexBuffer, err := Buffer.NewVertexBuffer(vertices, vertexBufferAttribFunc)
	if err != nil {
		return nil, err
	}

	return &Mesh{
		vertexBuffer: vertexBuffer,
		indexBuffer:  Buffer.NewIndexBuffer(indices),
	}, nil
}

func (mesh *Mesh) Draw() {
	mesh.vertexBuffer.Bind()
	mesh.indexBuffer.Bind()
	gl.DrawElements(gl.TRIANGLES, mesh.indexBuffer.Length, gl.UNSIGNED_INT, nil)
	mesh.indexBuffer.Unbind()
	mesh.vertexBuffer.Unbind()
}
