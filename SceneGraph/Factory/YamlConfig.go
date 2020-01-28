package Factory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/FrameBuffer"
	"github.com/Adi146/goggle-engine/Core/RenderTarget"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shader/DepthShader"
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
	"github.com/Adi146/goggle-engine/Core/Shadow"
	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
)

var NodeFactory = map[string]reflect.Type{
	"Scene.ChildBaseNode":        reflect.TypeOf((*Scene.ChildNodeBase)(nil)).Elem(),
	"Scene.ParentBaseNode":       reflect.TypeOf((*Scene.ParentNodeBase)(nil)).Elem(),
	"Scene.IntermediateNodeBase": reflect.TypeOf((*Scene.IntermediateNodeBase)(nil)).Elem(),
}

var MatrixFactory = map[string]reflect.Type{
	"translation": reflect.TypeOf((*TranslationConfig)(nil)).Elem(),
	"rotation":    reflect.TypeOf((*RotationConfig)(nil)).Elem(),
	"scale":       reflect.TypeOf((*ScaleConfig)(nil)).Elem(),
	"orthogonal":  reflect.TypeOf((*OrthogonalConfig)(nil)).Elem(),
	"perspective": reflect.TypeOf((*PerspectiveConfig)(nil)).Elem(),
}

var WindowFactory = map[string]reflect.Type{
	"sdl": reflect.TypeOf((*Window.SDLWindow)(nil)).Elem(),
}

var ShaderFactory = map[string]func([]string, []string) (Shader.IShaderProgram, error){
	//"basic": BasicShader.NewBasicIShaderProgram,
	"phong": PhongShader.NewPhongIShaderProgram,
	"depth": DepthShader.NewDepthIShaderProgram,
}

var FrameBufferFactory = map[string]reflect.Type{
	"depthBuffer": reflect.TypeOf((*Shadow.DepthBuffer)(nil)).Elem(),
	"sdlWindow":   reflect.TypeOf((*Window.SDLWindow)(nil)).Elem(),
}

type yamlConfig struct {
	RenderTargetConfig *RenderTarget.OpenGLRenderTarget `yaml:"renderTarget"`
	RootConfig         YamlNodeConfig                   `yaml:"root"`
	FrameBuffersConfig []struct {
		Type   string    `yaml:"type"`
		Config yaml.Node `yaml:"config"`
		Shader string    `yaml:"shader"`
	} `yaml:"frameBuffers"`
	ShadersConfig map[string]YamlShaderConfig `yaml:"shaders"`

	shaders map[string]Shader.IShaderProgram
}

func ReadYamlConfig(file *os.File) (*Config, error) {
	config := yamlConfig{
		shaders: map[string]Shader.IShaderProgram{},
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	frameBuffers, err := config.UnmarshalFrameBuffers()
	if err != nil {
		return nil, err
	}

	if err := config.RenderTargetConfig.Init(); err != nil {
		return nil, err
	}

	root, err := config.RootConfig.Unmarshal("root")
	if err != nil {
		return nil, err
	}

	scene := Scene.NewScene(config.RenderTargetConfig)
	if rootAsParent, isParent := root.(Scene.IParentNode); isParent {
		scene.SetRoot(rootAsParent)
	} else {
		return nil, fmt.Errorf("root is not a parent node")
	}

	return &Config{
		Scene:        scene,
		FrameBuffers: frameBuffers,
	}, nil
}

func (config *yamlConfig) UnmarshalFrameBuffers() ([]FrameBuffer.IFrameBuffer, error) {
	var frameBuffers []FrameBuffer.IFrameBuffer

	for _, frameBufferConfig := range config.FrameBuffersConfig {
		frameBufferType, ok := FrameBufferFactory[frameBufferConfig.Type]
		if !ok {
			return nil, fmt.Errorf("framebuffer type %s is not in factory", frameBufferConfig.Type)
		}

		frameBuffer := reflect.New(frameBufferType).Interface().(FrameBuffer.IFrameBuffer)

		frameBufferConfig.Config.Decode(frameBuffer)

		if err := frameBuffer.Init(); err != nil {
			return nil, err
		}

		shader, ok := config.shaders[frameBufferConfig.Shader]
		if !ok {
			shaderConfig := config.ShadersConfig[frameBufferConfig.Shader]
			newShader, err := shaderConfig.Unmarshal()
			if err != nil {
				return nil, err
			}
			config.shaders[frameBufferConfig.Shader] = newShader
			shader = newShader
		}
		frameBuffer.SetShaderProgram(shader)

		frameBuffers = append(frameBuffers, frameBuffer)
	}

	return frameBuffers, nil
}
