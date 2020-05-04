package Buffer

import "github.com/go-gl/gl/all-core/gl"

func newStaticBuffer(bufferType uint32, data interface{}) Buffer {
	return newBuffer(bufferType, data, gl.STATIC_DRAW)

}
