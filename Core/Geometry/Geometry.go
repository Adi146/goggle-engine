package Geometry

import (
	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Geometry struct {
	vertexBuffer *VertexBuffer
	indexBuffer  *IndexBuffer
}

func ImportGeometry(assimpMesh *assimp.Mesh) (*Geometry, error) {
	assimpVertices := assimpMesh.Vertices()
	assimpNormals := assimpMesh.Normals()
	assimpUVs := assimpMesh.TextureCoords(0)
	assimpFaces := assimpMesh.Faces()

	vertices := make([]Vertex, assimpMesh.NumVertices())
	for i := 0; i < assimpMesh.NumVertices(); i++ {
		vertices[i].Position = Vector.Vector3{assimpVertices[i].X(), assimpVertices[i].Y(), assimpVertices[i].Z()}
		vertices[i].Normal = Vector.Vector3{assimpNormals[i].X(), assimpNormals[i].Y(), assimpNormals[i].Z()}
		vertices[i].UV = Vector.Vector2{assimpUVs[i].X(), assimpUVs[i].Y()}
	}

	var indices []uint32
	for _, assimpFace := range assimpFaces {
		indices = append(indices, assimpFace.CopyIndices()...)
	}

	return NewGeometry(vertices, RegisterVertexBufferAttributes, indices)
}

func NewGeometry(vertices []Vertex, vertexBufferAttribFunc func(), indices []uint32) (*Geometry, error) {
	vertexBuffer, err := NewVertexBuffer(vertices, vertexBufferAttribFunc)
	if err != nil {
		return nil, err
	}

	return &Geometry{
		vertexBuffer: vertexBuffer,
		indexBuffer:  NewIndexBuffer(indices),
	}, nil
}

func (geo *Geometry) Draw() {
	geo.vertexBuffer.Bind()
	geo.indexBuffer.Bind()
	gl.DrawElements(gl.TRIANGLES, geo.indexBuffer.length, gl.UNSIGNED_INT, nil)
	geo.indexBuffer.Unbind()
	geo.vertexBuffer.Unbind()
}
