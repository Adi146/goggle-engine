package PostProcessing

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Shadow/ShadowMapShader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"

	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

const (
	ua_screenTexture = "u_screenTexture"

	ua_kernelOffset = "u_kernelOffset"
	ua_kernel       = "u_kernel"
)

func init() {
	ShaderFactory.AddType("postProcessing", NewIShaderProgram)
}

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
		return program.BindUniform(&v.ColorTexture, ua_screenTexture)
	case *ShadowMapShader.ShadowMapBuffer:
		return program.BindUniform(&v.ShadowMap, ua_screenTexture)
	case *Kernel:
		return program.bindKernel(v)
	default:
		return fmt.Errorf("post processing shader does not support type %T", v)
	}
}

func (program *ShaderProgram) bindKernel(kernel *Kernel) error {
	var err Error.ErrorCollection

	err.Push(program.BindUniform(kernel.GetOffset(), ua_kernelOffset))
	err.Push(program.BindUniform(kernel.GetKernel(), ua_kernel))

	return err.Err()
}
