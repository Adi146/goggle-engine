package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type InstancedMesh struct {
	VertexBuffer ArrayBuffer
	MatrixBuffer ArrayBuffer
	VertexArray  InstancedVertexArray
	IndexBuffer  *IndexBuffer

	NumInstances int32
}

func NewInstancedMesh(vertices []Vertex, indices []uint32, matrices []GeometryMath.Matrix4x4) *InstancedMesh {
	vbo := NewVertexBuffer(vertices)
	mbo := NewMatrixBuffer(matrices)

	return &InstancedMesh{
		VertexBuffer: vbo,
		MatrixBuffer: mbo,
		VertexArray:  NewInstancedVertexArray(vbo, mbo),
		IndexBuffer:  NewIndexBuffer(indices),
		NumInstances: int32(len(matrices)),
	}
}

func (mesh *InstancedMesh) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	var err Error.ErrorCollection

	err.Push(shader.BindObject(mesh.VertexArray))
	mesh.IndexBuffer.Bind()
	gl.DrawElementsInstanced(gl.TRIANGLES, mesh.IndexBuffer.Length, gl.UNSIGNED_INT, nil, mesh.NumInstances)
	mesh.IndexBuffer.Unbind()
	mesh.VertexBuffer.Unbind()

	return err.Err()
}
