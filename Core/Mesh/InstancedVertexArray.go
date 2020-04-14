package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/go-gl/gl/v4.1-core/gl"
	"unsafe"
)

type InstancedVertexArray struct {
	VertexArray
}

func NewInstancedVertexArray(vbo, mbo ArrayBuffer) InstancedVertexArray {
	matrix := *GeometryMath.Identity()
	vao := InstancedVertexArray{
		VertexArray: NewVertexArray(vbo),
	}

	mbo.Bind()
	defer mbo.Unbind()

	vao.Bind()
	vao.enableInstanceMatrix()

	gl.VertexAttribPointer(index_modelMatrix_1, int32(len(matrix[0])), gl.FLOAT, false, int32(unsafe.Sizeof(matrix)), unsafe.Pointer(0*unsafe.Sizeof(matrix[0])))
	gl.VertexAttribPointer(index_modelMatrix_2, int32(len(matrix[1])), gl.FLOAT, false, int32(unsafe.Sizeof(matrix)), unsafe.Pointer(1*unsafe.Sizeof(matrix[0])))
	gl.VertexAttribPointer(index_modelMatrix_3, int32(len(matrix[2])), gl.FLOAT, false, int32(unsafe.Sizeof(matrix)), unsafe.Pointer(2*unsafe.Sizeof(matrix[0])))
	gl.VertexAttribPointer(index_modelMatrix_4, int32(len(matrix[3])), gl.FLOAT, false, int32(unsafe.Sizeof(matrix)), unsafe.Pointer(3*unsafe.Sizeof(matrix[0])))

	gl.VertexAttribDivisor(index_modelMatrix_1, 1)
	gl.VertexAttribDivisor(index_modelMatrix_2, 1)
	gl.VertexAttribDivisor(index_modelMatrix_3, 1)
	gl.VertexAttribDivisor(index_modelMatrix_4, 1)

	return vao
}

func (vao InstancedVertexArray) enableInstanceMatrix() {
	gl.EnableVertexArrayAttrib(uint32(vao.VertexArray), index_modelMatrix_1)
	gl.EnableVertexArrayAttrib(uint32(vao.VertexArray), index_modelMatrix_2)
	gl.EnableVertexArrayAttrib(uint32(vao.VertexArray), index_modelMatrix_3)
	gl.EnableVertexArrayAttrib(uint32(vao.VertexArray), index_modelMatrix_4)
}

func (vao InstancedVertexArray) disableInstanceMatrix() {
	gl.DisableVertexArrayAttrib(uint32(vao.VertexArray), index_modelMatrix_1)
	gl.DisableVertexArrayAttrib(uint32(vao.VertexArray), index_modelMatrix_2)
	gl.DisableVertexArrayAttrib(uint32(vao.VertexArray), index_modelMatrix_3)
	gl.DisableVertexArrayAttrib(uint32(vao.VertexArray), index_modelMatrix_4)
}
