package Buffer

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"reflect"
)

type VertexBuffer struct {
	bufferId uint32
	vao      uint32
}

func sizeof(t reflect.Type) int {
	switch t.Kind() {
	case reflect.Array:
		if s := sizeof(t.Elem()); s >= 0 {
			return s * t.Len()
		}

	case reflect.Struct:
		sum := 0
		for i, n := 0, t.NumField(); i < n; i++ {
			s := sizeof(t.Field(i).Type)
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return int(t.Size())

	case reflect.Slice:
		return sizeof(t.Elem())
	}

	return -1
}

func NewVertexBuffer(vertices interface{}, vertexBufferAttribFunc func()) (*VertexBuffer, error) {
	vertexSlice := reflect.ValueOf(vertices)
	if vertexSlice.Kind() != reflect.Slice {
		return nil, fmt.Errorf("vertices must be a slice")
	}

	buff := VertexBuffer{}

	gl.GenVertexArrays(1, &buff.vao)
	gl.BindVertexArray(buff.vao)

	gl.GenBuffers(1, &buff.bufferId)
	gl.BindBuffer(gl.ARRAY_BUFFER, buff.bufferId)
	gl.BufferData(gl.ARRAY_BUFFER, vertexSlice.Len()*sizeof(vertexSlice.Type()), gl.Ptr(vertices), gl.STATIC_DRAW)

	vertexBufferAttribFunc()

	gl.BindVertexArray(0)

	return &buff, nil
}

func (buff *VertexBuffer) Destroy() {
	gl.DeleteBuffers(1, &buff.bufferId)
}

func (buff *VertexBuffer) Bind() {
	gl.BindVertexArray(buff.vao)
}

func (buff *VertexBuffer) Unbind() {
	gl.BindVertexArray(0)
}
