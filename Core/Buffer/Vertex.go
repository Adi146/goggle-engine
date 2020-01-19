package Buffer

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/go-gl/gl/v4.1-core/gl"
	"unsafe"
)

type Vertex struct {
	Position Vector.Vector3
	Normal   Vector.Vector3

	UV Vector.Vector2
}

func RegisterVertexBufferAttributes() {
	vertex := Vertex{}

	// position data
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.Position)))
	// normal vector
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.Normal)))
	// Texture coordinates
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, int32(unsafe.Sizeof(vertex)), unsafe.Pointer(unsafe.Offsetof(vertex.UV)))
}
