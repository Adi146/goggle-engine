package ShaderFactory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/FrameBufferFactory"
	"gopkg.in/yaml.v3"
)

type Product struct {
	Shader.IShaderProgram
}

type tmpProduct struct {
	Type            string                      `yaml:"type"`
	VertexShaders   []string                    `yaml:"vertexShaders"`
	FragmentShaders []string                    `yaml:"fragmentShaders"`
	FrameBuffers    []FrameBufferFactory.Config `yaml:"frameBuffers"`
}

func (product *Product) UnmarshalYAML(value *yaml.Node) error {
	var tmpProduct tmpProduct
	if err := value.Decode(&tmpProduct); err != nil {
		return err
	}

	shaderConstructor, ok := typeFactory[tmpProduct.Type]
	if !ok {
		return fmt.Errorf("shader type %s is not in factory", tmpProduct.Type)
	}

	shader, err := shaderConstructor(tmpProduct.VertexShaders, tmpProduct.FragmentShaders)
	if err != nil {
		return err
	}

	for _, fboConfig := range tmpProduct.FrameBuffers {
		if err := shader.BindObject(fboConfig.IFrameBuffer); err != nil {
			return err
		}
	}

	product.IShaderProgram = shader
	return nil
}
