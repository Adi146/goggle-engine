package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/BoundingBox"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Scene"
)

type IMesh interface {
	Scene.IDrawable
	BoundingBox.ICollisionObject

	GetVertexBuffer() ArrayBuffer
	GetVertexArray() VertexArray
	GetIndexBuffer() *IndexBuffer

	GetModelMatrix() GeometryMath.Matrix4x4
	SetModelMatrix(mat GeometryMath.Matrix4x4)
}
