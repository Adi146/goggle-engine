package PostProcessing

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/Light/ShadowMapping"
	"github.com/Adi146/goggle-engine/Core/Texture"

	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"
)

const (
	shader_factory_name = "postProcessingShader"

	ua_screenTexture = "u_screenTexture"

	ua_kernelOffset = "u_kernelOffset"
	ua_kernel       = "u_kernel"
)

func init() {
	Shader.Factory.AddConstructor(shader_factory_name, NewIShaderProgram)
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
	case FrameBuffer.IFrameBuffer:
		var errors Error.ErrorCollection
		for _, texture := range v.GetTextures() {
			errors.Push(program.BindObject(texture))
		}
		return errors.Err()
	case Texture.ITexture:
		switch t := v.GetType(); t {
		case OffscreenTexture, ShadowMapping.ShadowMap:
			return program.BindUniform(v, ua_screenTexture)
		default:
			return fmt.Errorf("post processing shader does not support texture of type %s", t)
		}
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
