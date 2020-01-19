package main

import (
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

	_ "image/jpeg"
)

const (
	width  = 500
	height = 500

	modelFile   = "Models/wall.fbx"
	diffuseFile = "Models/brickwall.jpg"
	normalFile  = "Models/brickwall_normal.jpg"

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

	pointLightRotor := &SceneGraph.Rotor{
		IIntermediateNode: Scene.NewIntermediateNodeBase(),
		Speed:             -1,
	}

	pointLightNode1 := &LightNode.PointLightNode{
		IChildNode: Scene.NewChildNodeBase(),
		PointLight: Light.PointLight{
			Position:  Vector.Vector3{0.0, 0.0, 0.0},
			Ambient:   Vector.Vector3{0.32, 0.32, 0.32},
			Diffuse:   Vector.Vector3{0.8, 0.8, 0.8},
			Specular:  Vector.Vector3{0.8, 0.8, 0.8},
			Linear:    0.027,
			Quadratic: 0.0028,
		},
	}
	pointLightNode1.SetLocalTransformation(Matrix.Translate(&Vector.Vector3{0.0, 0.0, 30.0}))

	modelNode := &Node.ModelNode{
		IChildNode: Scene.NewChildNodeBase(),
		File:       modelFile,
		Textures: Node.TextureConfiguration{
			Diffuse: []string{
				diffuseFile,
			},
			Normals: []string{
				normalFile,
			},
		},
	}
	if err := modelNode.Init(); err != nil {
		panic(err)
	}
	modelNode.SetLocalTransformation(Matrix.Scale(0.1))

	scene := Scene.NewScene(openGLRenderTarget)
	scene.SetRoot(Scene.NewParentNodeBase())

	scene.Root.AddChild(controlNode)
	controlNode.AddChild(cameraNode)
	scene.Root.AddChild(pointLightRotor)
	pointLightRotor.AddChild(pointLightNode1)
	scene.Root.AddChild(modelNode)

	RenderTarget.RunRenderLoop(scene)
}
