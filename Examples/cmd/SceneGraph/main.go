package main

import (
	"github.com/Adi146/goggle-engine/Core/AssetImporter"
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Angle"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Matrix"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Light"
	"github.com/Adi146/goggle-engine/Core/RenderTarget"
	"github.com/Adi146/goggle-engine/Core/Shader/PhongShader"
	"github.com/Adi146/goggle-engine/Core/Window"
	"github.com/Adi146/goggle-engine/Examples/SceneGraph"
	"github.com/Adi146/goggle-engine/SceneGraph/Node"
	"github.com/Adi146/goggle-engine/SceneGraph/Node/LightNode"
	"github.com/Adi146/goggle-engine/SceneGraph/Scene"
	"github.com/Adi146/goggle-engine/UI/Control"
	"runtime"
)

const (
	width  = 500
	height = 500

	modelFile      = "Models/suzanne.obj"
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
		KeyboardSensitivity: 1,
		MouseSensitivity:    0.5,
	}
	controlNode.Init("CameraControl")
	controlNode.Translate(&Vector.Vector3{0, 0, 5})

	cameraNode := &Node.CameraNode{
		Camera:       Camera.NewCameraPerspective(Angle.Radians(90), float32(width), float32(height)),
		InitialFront: &Vector.Vector3{0, 0, -1},
		InitialUp:    &Vector.Vector3{0, 1, 0},
	}
	cameraNode.Init("Camera")

	sunLightNode := &LightNode.DirectionalLightNode{
		DirectionalLight: Light.DirectionalLight{
			Direction: Vector.Vector3{0.0, 0.0, 0.0},
			Ambient:   Vector.Vector3{0.32, 0.32, 0.32},
			Diffuse:   Vector.Vector3{0.8, 0.8, 0.8},
			Specular:  Vector.Vector3{0.8, 0.8, 0.8},
		},
		InitialDirection: &Vector.Vector3{-1.0, -1.0, -1.0},
	}
	sunLightNode.Init("SunLight")

	pointLightRotor := &SceneGraph.Rotor{
		Speed:             -1,
	}
	pointLightRotor.Init("PointLightRotor")

	pointLightNode1 := &LightNode.PointLightNode{
		PointLight: Light.PointLight{
			Position:  Vector.Vector3{0.0, 0.0, 0.0},
			Ambient:   Vector.Vector3{0.0, 0.0, 0.2},
			Diffuse:   Vector.Vector3{0.0, 0.0, 1.0},
			Specular:  Vector.Vector3{0.0, 0.0, 1.0},
			Linear:    0.027,
			Quadratic: 0.0028,
		},
	}
	pointLightNode1.Init("PointLightBlue")
	pointLightNode1.SetLocalTransformation(Matrix.Translate(&Vector.Vector3{0.0, 0.0, 30.0}))

	pointLightNode2 := &LightNode.PointLightNode{
		PointLight: Light.PointLight{
			Position:  Vector.Vector3{0.0, 0.0, 0.0},
			Ambient:   Vector.Vector3{0.2, 0.0, 0.0},
			Diffuse:   Vector.Vector3{1.0, 0.0, 0.0},
			Specular:  Vector.Vector3{1.0, 0.0, 0.0},
			Linear:    0.027,
			Quadratic: 0.0028,
		},
	}
	pointLightNode2.Init("PointLightRed")
	pointLightNode2.SetLocalTransformation(Matrix.Translate(&Vector.Vector3{0.0, 0.0, -30.0}))

	modelRotorNode := &SceneGraph.Rotor{
		Speed:             1,
	}
	modelRotorNode.Init("ModelRotor")

	model, result := AssetImporter.ImportModel(modelFile)
	if !result.Success() {
		panic(err)
	}

	modelNode := &Node.ModelNode{
		Model:      model,
	}
	modelNode.Init("Model")

	rootNode := &Scene.ParentNodeBase{}
	rootNode.Init("Root")

	scene := Scene.NewScene(openGLRenderTarget)
	scene.SetRoot(rootNode)

	scene.Root.AddChild(controlNode)
	controlNode.AddChild(cameraNode)
	scene.Root.AddChild(sunLightNode)
	scene.Root.AddChild(pointLightRotor)
	pointLightRotor.AddChild(pointLightNode1)
	pointLightRotor.AddChild(pointLightNode2)
	scene.Root.AddChild(modelRotorNode)
	modelRotorNode.AddChild(modelNode)

	RenderTarget.RunRenderLoop(scene)
}
