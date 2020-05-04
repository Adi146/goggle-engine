package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"github.com/go-gl/gl/v4.3-core/gl"
)

type PrimitiveType uint32

type Mesh struct {
	VertexBuffer              Buffer.Buffer
	VertexArray               VertexArray
	IndexBuffer               *IndexBuffer
	ModelMatrix               GeometryMath.Matrix4x4
	boundingVolume            BoundingVolume.IBoundingVolume
	TransformedBoundingVolume BoundingVolume.IBoundingVolume

	PrimitiveType  PrimitiveType
	FrustumCulling bool
}

func NewMesh(vertices []Vertex, indices []uint32, boundingVolume func(vertices []GeometryMath.Vector3) BoundingVolume.IBoundingVolume) *Mesh {
	vbo := Buffer.NewStaticArrayBuffer(&vertices)

	mesh := Mesh{
		VertexBuffer:  vbo,
		VertexArray:   NewVertexArray(vbo),
		IndexBuffer:   NewIndexBuffer(indices),
		ModelMatrix:   GeometryMath.Identity(),
		PrimitiveType: gl.TRIANGLES,
	}

	if boundingVolume != nil {
		mesh.boundingVolume = boundingVolume(Vertices(vertices).GetPositions())
		mesh.TransformedBoundingVolume = mesh.boundingVolume.Transform(mesh.ModelMatrix)
	}

	return &mesh
}

func (mesh *Mesh) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene, camera Camera.ICamera) error {
	var err Error.ErrorCollection

	if !mesh.FrustumCulling || (mesh.FrustumCulling && camera.GetFrustum().Contains(mesh.GetBoundingVolume())) {
		err.Push(shader.BindObject(&mesh.ModelMatrix))
		err.Push(shader.BindObject(mesh.VertexArray))
		mesh.IndexBuffer.Bind()
		gl.DrawElements(uint32(mesh.PrimitiveType), mesh.IndexBuffer.Length, gl.UNSIGNED_INT, nil)
		mesh.IndexBuffer.Unbind()
	}

	return err.Err()
}

func (mesh *Mesh) GetVertexBuffer() Buffer.Buffer {
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
	mesh.TransformedBoundingVolume = mesh.boundingVolume.Transform(mesh.GetModelMatrix())
}

func (mesh *Mesh) GetBoundingVolume() BoundingVolume.IBoundingVolume {
	return mesh.TransformedBoundingVolume
}

func (mesh *Mesh) GetPrimitiveType() PrimitiveType {
	return mesh.PrimitiveType
}

func (mesh *Mesh) EnableFrustumCulling() {
	mesh.FrustumCulling = true
}

func (mesh *Mesh) DisableFrustumCulling() {
	mesh.FrustumCulling = false
}
