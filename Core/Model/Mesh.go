package Model

import (
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/go-gl/gl/v4.1-core/gl"
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

func (mesh *Mesh) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	mesh.vertexBuffer.Bind()
	mesh.indexBuffer.Bind()
	gl.DrawElements(gl.TRIANGLES, mesh.indexBuffer.Length, gl.UNSIGNED_INT, nil)
	mesh.indexBuffer.Unbind()
	mesh.vertexBuffer.Unbind()

	return nil
}
