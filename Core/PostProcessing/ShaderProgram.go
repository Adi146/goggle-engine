package PostProcessing

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Shadow"

	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

const (
	screenTexture_uniformAddress = "u_screenTexture"

	kernelOffset_uniformAddress = "u_kernelOffset"
	kernel_uniformAddress       = "u_kernel"
)

type ShaderProgram struct {
	*Shader.ShaderProgramCore
}

func NewShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (*ShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles)
	if err != nil {
		return nil, err
	}

	return &ShaderProgram{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewShaderProgram(vertexShaderFiles, fragmentShaderFiles)
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	switch v := i.(type) {
	case *FrameBuffer.OffScreenBuffer:
		program.Bind()
		return program.BindTexture(&v.ColorTexture, screenTexture_uniformAddress, true)
	case *Shadow.ShadowMapBuffer:
		program.Bind()
		return program.BindTexture(&v.ShadowMap, screenTexture_uniformAddress, true)
	case *Kernel:
		return program.bindKernel(v)
	default:
		return fmt.Errorf("post processing shader dows not support type %T", v)
	}
}

func (program *ShaderProgram) bindKernel(kernel *Kernel) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(kernel.GetOffset(), kernelOffset_uniformAddress))
	err.Push(program.BindUniform(kernel.GetKernel(), kernel_uniformAddress))

	return err.Err()
}
