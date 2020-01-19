package main

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/RenderTarget"
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/Examples/SceneGraph"
	"github.com/Adi146/goggle-engine/SceneGraph/Node"
	"github.com/Adi146/goggle-engine/SceneGraph/Node/LightNode"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"github.com/Adi146/goggle-engine/UI/Control"
	"runtime"

	_ "github.com/ftrvxmtrx/tga"
)

const (
	width  = 500
	height = 500

	modelFile      = "Models/fern.fbx"
	vertexShader   = "../Core/Shader/PhongShader/phong.vert"
	fragmentShader = "../Core/Shader/PhongShader/phong.frag"
)

func main() {
	runtime.LockOSThread()

	window := &Window.SDLWindow{
		Title:  "SceneGraph Example",
		Width:  width,
		Height: height,
		Type:   "window",
	}
	if err := window.Init(); err != nil {
		panic(err)
	}

	openGLRenderTarget := &RenderTarget.OpenGLRenderTarget{
		Window:       window,
		Culling:      true,
		DepthTest:    true,
		DebugLogging: false,
	}
	if err := openGLRenderTarget.Init(); err != nil {
		panic(err)
	}

	shaderProgram, err := PhongShader.NewPhongShaderProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
	openGLRenderTarget.SetActiveShaderProgram(shaderProgram)

	controlNode := &Control.WASDControl{
		IIntermediateNode:   Scene.NewIntermediateNodeBase(),
		KeyboardSensitivity: 10,
		MouseSensitivity:    0.5,
	}
	controlNode.Translate(&Vector.Vector3{0, 0, 5})

	cameraNode := &Node.CameraNode{
		IChildNode:   Scene.NewChildNodeBase(),
		Camera:       Camera.NewCameraPerspective(Angle.Radians(90), float32(width), float32(height)),
		InitialFront: &Vector.Vector3{0, 0, -1},
		InitialUp:    &Vector.Vector3{0, 1, 0},
	}

	sunLightNode := &LightNode.DirectionalLightNode{
		IChildNode: Scene.NewChildNodeBase(),
		DirectionalLight: Light.DirectionalLight{
			Direction: Vector.Vector3{0.0, 0.0, 0.0},
			Ambient:   Vector.Vector3{0.32, 0.32, 0.32},
			Diffuse:   Vector.Vector3{0.8, 0.8, 0.8},
			Specular:  Vector.Vector3{0.8, 0.8, 0.8},
		},
		InitialDirection: &Vector.Vector3{-1.0, -1.0, -1.0},
	}

	modelRotorNode := &SceneGraph.Rotor{
		IIntermediateNode: Scene.NewIntermediateNodeBase(),
		Speed:             0,
	}

	model, err := Model.ImportModel(modelFile)
	if err != nil {
		panic(err)
	}

	geometryNode := &Node.ModelNode{
		IChildNode: Scene.NewChildNodeBase(),
		Model:      model,
	}
	geometryNode.SetLocalTransformation(Matrix.Scale(0.1))

	scene := Scene.NewScene(openGLRenderTarget)
	scene.SetRoot(Scene.NewParentNodeBase())

	scene.Root.AddChild(controlNode)
	controlNode.AddChild(cameraNode)
	scene.Root.AddChild(sunLightNode)
	scene.Root.AddChild(modelRotorNode)
	modelRotorNode.AddChild(geometryNode)

	RenderTarget.RunRenderLoop(scene)
}
