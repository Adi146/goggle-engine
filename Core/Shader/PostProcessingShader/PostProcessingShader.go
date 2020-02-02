package PostProcessingShader

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
)

const (
	screenTexture_uniformAddress = "u_screenTexture"
)

type PostProcessingShaderProgram struct {
	*Shader.ShaderProgramCore
}

func NewPostProcessingShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (*PostProcessingShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &PostProcessingShaderProgram{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewPostProcessingIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewPostProcessingShaderProgram(vertexShaderFiles, fragmentShaderFiles)
}

func (program *PostProcessingShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *Model.Texture:
		return program.bindTexture(v)
	default:
		return fmt.Errorf("type %T not supported", v)
	}
}

func (program *PostProcessingShaderProgram) bindTexture(texture *Model.Texture) error {
	return program.BindTexture(0, texture, screenTexture_uniformAddress)
}
