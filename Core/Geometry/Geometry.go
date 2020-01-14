package Geometry

import (
	"encoding/binary"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/go-gl/gl/v4.1-core/gl"
	"os"
)

type Geometry struct {
	vertexBuffer *VertexBuffer
	indexBuffer  *IndexBuffer

	ModelMatrix *Matrix.Matrix4x4
}

func NewGeometryFromFile(file *os.File) (*Geometry, error) {
	var numVertices uint64
	if err := binary.Read(file, binary.LittleEndian, &numVertices); err != nil {
		return nil, err
	}

	var numIndices uint64
	if err := binary.Read(file, binary.LittleEndian, &numIndices); err != nil {
		return nil, err
	}

	vertices := make([]Vertex, numVertices)
	for i := uint64(0); i < numVertices; i++ {
		if err := binary.Read(file, binary.LittleEndian, &vertices[i].Position); err != nil {
			return nil, err
		}

		if err := binary.Read(file, binary.LittleEndian, &vertices[i].Normal); err != nil {
			return nil, err
		}
	}

	indices := make([]uint32, numIndices)
	for i := uint64(0); i < numIndices; i++ {
		if err := binary.Read(file, binary.LittleEndian, &indices[i]); err != nil {
			return nil, err
		}
	}

	return NewGeometry(vertices, RegisterVertexBufferAttributes, indices)
}

func NewGeometry(vertices []Vertex, vertexBufferAttribFunc func(), indices []uint32) (*Geometry, error) {
	vertexBuffer, err := NewVertexBuffer(vertices, vertexBufferAttribFunc)
	if err != nil {
		return nil, err
	}

	geo := Geometry{
		vertexBuffer: vertexBuffer,
		indexBuffer:  NewIndexBuffer(indices),

		ModelMatrix: Matrix.Identity(),
	}

	return &geo, nil
}

func (geo *Geometry) Draw() {
	geo.vertexBuffer.Bind()
	geo.indexBuffer.Bind()
	gl.DrawElements(gl.TRIANGLES, geo.indexBuffer.length, gl.UNSIGNED_INT, nil)
	geo.indexBuffer.Unbind()
	geo.vertexBuffer.Unbind()
}
