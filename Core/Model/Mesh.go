package Model

import (
	"github.com/Adi146/assimp"
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Mesh struct {
	vertexBuffer *Buffer.VertexBuffer
	indexBuffer  *Buffer.IndexBuffer
}

func ImportMesh(assimpMesh *assimp.Mesh) (*Mesh, error) {
	assimpVertices := assimpMesh.Vertices()
	assimpNormals := assimpMesh.Normals()
	assimpUVs := assimpMesh.TextureCoords(0)
	assimpFaces := assimpMesh.Faces()
	assimpTangents := assimpMesh.Tangents()

	vertices := make([]Buffer.Vertex, assimpMesh.NumVertices())
	for i := 0; i < assimpMesh.NumVertices(); i++ {
		vertices[i].Position = Vector.Vector3{assimpVertices[i].X(), assimpVertices[i].Y(), assimpVertices[i].Z()}
		vertices[i].Normal = Vector.Vector3{assimpNormals[i].X(), assimpNormals[i].Y(), assimpNormals[i].Z()}
		vertices[i].UV = Vector.Vector2{assimpUVs[i].X(), assimpUVs[i].Y()}
		vertices[i].Tangent = Vector.Vector3{assimpTangents[i].X(), assimpTangents[i].Y(), assimpTangents[i].Z()}
	}

	var indices []uint32
	for _, assimpFace := range assimpFaces {
		indices = append(indices, assimpFace.CopyIndices()...)
	}

	return NewMesh(vertices, Buffer.RegisterVertexBufferAttributes, indices)
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
