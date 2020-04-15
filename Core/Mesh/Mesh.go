package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Mesh struct {
	VertexBuffer ArrayBuffer
	VertexArray  VertexArray
	IndexBuffer  *IndexBuffer
	ModelMatrix  GeometryMath.Matrix4x4
}

func NewMesh(vertices []Vertex, indices []uint32) *Mesh {
	vbo := NewVertexBuffer(vertices)

	return &Mesh{
		VertexBuffer: vbo,
		VertexArray:  NewVertexArray(vbo),
		IndexBuffer:  NewIndexBuffer(indices),
		ModelMatrix:  *GeometryMath.Identity(),
	}
}

func (mesh *Mesh) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	var err Error.ErrorCollection

	err.Push(shader.BindObject(&mesh.ModelMatrix))
	err.Push(shader.BindObject(mesh.VertexArray))
	mesh.IndexBuffer.Bind()
	gl.DrawElements(gl.TRIANGLES, mesh.IndexBuffer.Length, gl.UNSIGNED_INT, nil)
	mesh.IndexBuffer.Unbind()
	mesh.VertexBuffer.Unbind()

	return err.Err()
}

func (mesh *Mesh) GetVertexBuffer() ArrayBuffer {
	return mesh.VertexBuffer
}

func (mesh *Mesh) GetVertexArray() VertexArray {
	return mesh.VertexArray
}

func (mesh *Mesh) GetIndexBuffer() *IndexBuffer {
	return mesh.IndexBuffer
}

func (mesh *Mesh) GetModelMatrix() GeometryMath.Matrix4x4 {
	return mesh.ModelMatrix
}

func (mesh *Mesh) SetModelMatrix(mat GeometryMath.Matrix4x4) {
	mesh.ModelMatrix = mat
}
