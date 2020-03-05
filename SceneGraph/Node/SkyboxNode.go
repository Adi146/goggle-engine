package Node

import (
	"github.com/Adi146/goggle-engine/Core/AssetImporter"
	"github.com/Adi146/goggle-engine/Core/Skybox"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/NodeFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Factory/ShaderFactory"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"reflect"
)

const SkyboxNodeFactoryName = "Node.Skybox"
const SkyboxShaderFactoryName = "skybox"

func init() {
	NodeFactory.AddType(SkyboxNodeFactoryName, reflect.TypeOf((*SkyboxNodeConfig)(nil)).Elem())
	ShaderFactory.AddType(SkyboxShaderFactoryName, Skybox.NewIShaderProgram)
}

type SkyboxNodeConfig struct {
	Scene.NodeConfig

	Shader ShaderFactory.Config `yaml:"shader"`
	Textures AssetImporter.CubeMapImportHelper `yaml:"textures"`
}

func (config *SkyboxNodeConfig) Create() (Scene.INode, error) {
	nodeBase, err := config.NodeConfig.Create()
	if err != nil {
		return nil, err
	}

	node := &SkyboxNode{
		INode :nodeBase,
		Config: config,
	}

	cubeMap, result := AssetImporter.ImportCubeMap(config.Textures)
	if result.Errors.Len() > 0 {
		return nil, result.Errors.Err()
	}

	skyBox, err := Skybox.NewSkybox(cubeMap, config.Shader.IShaderProgram)
	if err != nil {
		return nil, err
	}

	node.Skybox = skyBox

	return node, nil
}

type SkyboxNode struct {
	Scene.INode
	*Skybox.Skybox

	Config *SkyboxNodeConfig
}

func (node *SkyboxNode) Draw() error {
	if scene := node.GetScene(); scene != nil {
		scene.OpaqueObjects = append(scene.OpaqueObjects, node.Skybox)
	}

	return nil
}
