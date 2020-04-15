package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/go-gl/gl/v4.1-core/gl"
	"unsafe"
)

func NewInstancedVertexArray(vao VertexArray, mbo ArrayBuffer) VertexArray {
	matrix := GeometryMath.Matrix4x4{}

	mbo.Bind()
	defer mbo.Unbind()

	vao.Bind()
	gl.EnableVertexArrayAttrib(uint32(vao), index_modelMatrix_1)
	gl.EnableVertexArrayAttrib(uint32(vao), index_modelMatrix_2)
	gl.EnableVertexArrayAttrib(uint32(vao), index_modelMatrix_3)
	gl.EnableVertexArrayAttrib(uint32(vao), index_modelMatrix_4)

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
