package Factory

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/RenderTarget"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Shader/BasicShader"
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
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

var ShaderFactory = map[string]func(string, string) (Shader.IShaderProgram, error){
	"basic": BasicShader.NewBasicIShaderProgram,
	"phong": PhongShader.NewPhongIShaderProgram,
}

type YamlConfig struct {
	WindowConfig struct {
		Library string    `yaml:"library"`
		Config  yaml.Node `yaml:"config"`
	} `yaml:"window"`
	RenderTargetConfig *RenderTarget.OpenGLRenderTarget `yaml:"renderTarget"`
	RootConfig         YamlNodeConfig                   `yaml:"root"`
	ShaderConfig       struct {
		Type           string `yaml:"type"`
		VertexShader   string `yaml:"vertexShader"`
		FragmentShader string `yaml:"fragmentShader"`
	} `yaml:"shader"`

	RootNode Scene.IParentNode
}

func ReadYamlConfig(file *os.File) (*Scene.Scene, error) {
	config := YamlConfig{}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	window, err := config.UnmarshalWindow()
	if err != nil {
		return nil, err
	}

	config.RenderTargetConfig.Window = window
	if err := config.RenderTargetConfig.Init(); err != nil {
		return nil, err
	}

	shader, err := config.UnmarshalShader()
	if err != nil {
		return nil, err
	}
	config.RenderTargetConfig.SetActiveShaderProgram(shader)

	root, err := config.RootConfig.Unmarshal()
	if err != nil {
		return nil, err
	}

	if rootAsParent, isParent := root.(Scene.IParentNode); isParent {
		config.RootNode = rootAsParent
	} else {
		return nil, fmt.Errorf("root is not a parent node")
	}

	scene := Scene.NewScene(config.RenderTargetConfig)
	scene.SetRoot(config.RootNode)

	return scene, nil
}

func (config *YamlConfig) UnmarshalWindow() (Window.IWindow, error) {
	windowLibrary, ok := WindowFactory[config.WindowConfig.Library]
	if !ok {
		return nil, fmt.Errorf("window library %s is not in factory", config.WindowConfig.Library)
	}

	window := reflect.New(windowLibrary).Interface().(Window.IWindow)

	config.WindowConfig.Config.Decode(window)

	if err := window.Init(); err != nil {
		return nil, err
	}

	return window, nil
}

func (config *YamlConfig) UnmarshalShader() (Shader.IShaderProgram, error) {
	shaderConstructor, ok := ShaderFactory[config.ShaderConfig.Type]
	if !ok {
		return nil, fmt.Errorf("shader type %s is not in factory", config.ShaderConfig.Type)
	}

	return shaderConstructor(config.ShaderConfig.VertexShader, config.ShaderConfig.FragmentShader)
}
