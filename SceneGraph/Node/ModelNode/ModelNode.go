package ModelNode

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/Core/Model"
	coreScene "github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	_ "github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"github.com/Adi146/goggle-engine/Utils/Log"
	"gopkg.in/yaml.v3"
	"reflect"
)

const ModelNodeFactoryName = "Node.ModelNode"

func init() {
	Scene.NodeFactory.AddType(ModelNodeFactoryName, reflect.TypeOf((*ModelNode)(nil)).Elem())
}

type ModelNode struct {
	Scene.INode
	Model.Model
	IsTransparent bool
	Shader        Shader.IShaderProgram

	MasterMatrix GeometryMath.Matrix4x4
	Slaves       []*ModelSlaveNode
}

func (node *ModelNode) Tick(timeDelta float32) error {
	err := node.INode.Tick(timeDelta)

	node.SetModelMatrix(node.GetGlobalTransformation())
	if instancedMesh, isInstancedMesh := node.Model.IMesh.(*Mesh.InstanceMasterMesh); isInstancedMesh {
		instancedMesh.SetMasterMatrix(node.MasterMatrix)
	}

	if scene := node.GetScene(); scene != nil {
		if node.IsTransparent {
			scene.AddTransparentObject(node)
		} else {
			scene.AddOpaqueObject(node)
		}
	}

	return err
}

func (node *ModelNode) Draw(shader Shader.IShaderProgram, invoker coreScene.IDrawable, scene coreScene.IScene, camera Camera.ICamera) error {
	if shader == nil {
		node.Shader.Bind()
		defer node.Shader.Unbind()

		shader = node.Shader
	}

	return node.Model.Draw(shader, invoker, scene, camera)
}

func (node *ModelNode) SetBase(base Scene.INode) {
	node.INode = base
}

func (node *ModelNode) AddSlave(slave ...*ModelSlaveNode) error {
	var newInstances []*Mesh.InstancedMesh

	switch v := node.IMesh.(type) {
	case *Mesh.Mesh:
		master := Mesh.NewInstanceMasterMesh(v, slaves(slave).GetGlobalTransformations()...)
		master.SetMasterMatrix(node.MasterMatrix)
		newInstances = master.Instances
		node.IMesh = master
	case *Mesh.InstanceMasterMesh:
		newInstances = v.CreateNewInstances(slaves(slave).GetGlobalTransformations()...)
	default:
		return fmt.Errorf("mesh of type %T can not be converted to instance master", v)
	}

	slaves(slave).SetInstancedMeshes(node, newInstances...)
	node.Slaves = append(node.Slaves, slave...)

	Log.Info(fmt.Sprintf("%d new slaves added", len(slave)))

	return nil
}

func (node *ModelNode) UnmarshalYAML(value *yaml.Node) error {
	yamlConfig := struct {
		Model         Model.Model            `yaml:",inline"`
		IsTransparent bool                   `yaml:"isTransparent"`
		Shader        Shader.Ptr             `yaml:"shader"`
		MasterMatrix  GeometryMath.Matrix4x4 `yaml:"masterMatrix"`
	}{
		Model:         node.Model,
		IsTransparent: node.IsTransparent,
		Shader: Shader.Ptr{
			IShaderProgram: node.Shader,
		},
		MasterMatrix: node.MasterMatrix,
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	node.Model = yamlConfig.Model
	node.IsTransparent = yamlConfig.IsTransparent
	node.Shader = yamlConfig.Shader
	node.MasterMatrix = yamlConfig.MasterMatrix
	node.SetModelMatrix(node.GetGlobalTransformation())
	node.EnableFrustumCulling()

	if node.MasterMatrix == (GeometryMath.Matrix4x4{}) {
		node.MasterMatrix = GeometryMath.Identity()
	} else {
		if err := node.AddSlave(); err != nil {
			return err
		}
	}

	return Scene.UnmarshalChildren(value, node)
}
