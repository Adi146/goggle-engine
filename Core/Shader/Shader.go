package Shader

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	Source   string
	Type     uint32
	shaderId uint32
}

func NewShaderFromSource(source string, shaderType uint32) (*Shader, error) {
	shader := &Shader{
		Source: source,
		Type:   shaderType,
	}
	if err := shader.compile(); err != nil {
		return nil, err
	}

	return shader, nil
}

func NewShaderFromFile(filename string, shaderType uint32) (*Shader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	source, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return NewShaderFromSource(string(source)+"\x00", shaderType)
}

func (shader *Shader) Destroy() {
	gl.DeleteShader(shader.shaderId)
}

func (shader *Shader) compile() error {
	shader.shaderId = gl.CreateShader(shader.Type)

	src, free := gl.Strs(shader.Source)
	gl.ShaderSource(shader.shaderId, 1, src, nil)
	free()
	gl.CompileShader(shader.shaderId)

	var status int32
	gl.GetShaderiv(shader.shaderId, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader.shaderId, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader.shaderId, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to compile %v: %v", shader.Source, log)
	}

	return nil
}
