package PostProcessing

import (
	"fmt"
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

func NewShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string, geometryShaderFiles []string) (*ShaderProgram, error) {
	shaderCore, err := Shader.NewShaderProgramFromFiles(vertexShaderFiles, fragmentShaderFiles, geometryShaderFiles)
	if err != nil {
		return nil, err
	}

	return &ShaderProgram{
		ShaderProgramCore: shaderCore,
	}, nil
}

func NewIShaderProgram(vertexShaderFiles []string, fragmentShaderFiles []string, geometryShaderFiles []string) (Shader.IShaderProgram, error) {
	return NewShaderProgram(vertexShaderFiles, fragmentShaderFiles, geometryShaderFiles)
}

func (program *ShaderProgram) GetUniformAddress(i interface{}) (string, error) {
	switch v := i.(type) {
	case Texture.ITexture:
		switch t := v.GetType(); t {
		case OffscreenTexture:
			return ua_screenTexture, nil
		default:
			return "", fmt.Errorf("post processing shader does not support texture of type %s", t)
		}
	case float32:
		return ua_kernelOffset, nil
	case []float32:
		return ua_kernel, nil
	default:
		return "", fmt.Errorf("post processing shader does not support type %T", v)
	}
}

func (program *ShaderProgram) BindObject(i interface{}) error {
	kernel, isKernel := i.(*Kernel)
	if isKernel {
		var err Error.ErrorCollection
		err.Push(program.BindObject(kernel.GetOffset()))
		err.Push(program.BindObject(kernel.GetKernel()))
		return err.Err()
	}

	uniformAddress, err := program.GetUniformAddress(i)
	if err != nil {
		return err
	}
	return program.BindUniform(i, uniformAddress)
}
