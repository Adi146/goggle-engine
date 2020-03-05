package Log

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"github.com/go-gl/gl/v3.3-core/gl"
	"unsafe"
)

func EnableDebugLogging() {
	gl.Enable(gl.DEBUG_OUTPUT)
	gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS)
	gl.DebugMessageCallback(openGLDebugCallback, nil)
}

func openGLDebugCallback(source uint32, gltype uint32, id uint32, severity uint32, length int32, message string, userParam unsafe.Pointer) {
	Log.Error(fmt.Errorf("[OpenGL Error] %s", message), "OpenGL Error")
}
