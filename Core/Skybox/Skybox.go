package Skybox

import (
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/go-gl/gl/v4.1-core/gl"
)

var (
	vertieces = []Buffer.Vertex{
		{Position: Vector.Vector3{-1.0, 1.0, 1.0}},
		{Position: Vector.Vector3{-1.0, 1.0, -1.0}},
		{Position: Vector.Vector3{-1.0, -1.0, 1.0}},
		{Position: Vector.Vector3{-1.0, -1.0, -1.0}},

		{Position: Vector.Vector3{1.0, 1.0, 1.0}},
		{Position: Vector.Vector3{1.0, 1.0, -1.0}},
		{Position: Vector.Vector3{1.0, -1.0, 1.0}},
		{Position: Vector.Vector3{1.0, -1.0, -1.0}},
	}
	indices = []uint32{
		0, 2, 1, //Left
		2, 3, 1,

		2, 0, 6, //Front
		0, 4, 6,

		6, 7, 2, //Bottom
		7, 3, 2,

		6, 4, 7, //Right
		4, 5, 7,

		0, 1, 4, //Top
		1, 5, 4,

		3, 7, 1, //Back
		7, 5, 1,
	}
)

type Skybox struct {
	*Model.Mesh
	*Texture.Texture
	Shader Shader.IShaderProgram
}

func NewSkybox(texture *Texture.Texture, shader Shader.IShaderProgram) (*Skybox, error) {
	mesh, err := Model.NewMesh(vertieces, Buffer.RegisterVertexBufferAttributes, indices)
	if err != nil {
		return nil, err
	}

	return &Skybox{
		Mesh:    mesh,
		Texture: texture,
		Shader:  shader,
	}, nil
}

func (skybox *Skybox) Draw(step *Scene.ProcessingPipelineStep) error {
	var oldDepthFunc int32
	gl.GetIntegerv(gl.DEPTH_FUNC, &oldDepthFunc)
	gl.DepthFunc(gl.LEQUAL)

	var shader Shader.IShaderProgram
	if step.EnforcedShader == nil {
		shader = skybox.Shader
	} else {
		shader = step.EnforcedShader
	}

	shader.Bind()
	err := shader.BindObject(skybox.Texture)
	skybox.Mesh.Draw()
	shader.Unbind()

	gl.DepthFunc(uint32(oldDepthFunc))

	return err
}

func (skybox *Skybox) GetPosition() *Vector.Vector3 {
	return &Vector.Vector3{0, 0, 0}
}
