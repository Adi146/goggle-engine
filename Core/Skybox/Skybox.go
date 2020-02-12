package Skybox

import (
	"github.com/Adi146/goggle-engine/Core/Buffer"
	"github.com/Adi146/goggle-engine/Core/GeometryMath/Vector"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
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
		0, 1, 2, //Left
		2, 3, 1,

		0, 2, 6, //Front
		6, 4, 0,

		2, 6, 7, //Bottom
		7, 2, 3,

		6, 4, 7, //Right
		7, 5, 4,

		0, 4, 1, //Top
		1, 5, 4,

		3, 1, 7, //Back
		7, 1, 5,
	}
)

type Skybox struct {
	*Model.Mesh
	*Texture.CubeMap
}

func NewSkybox(texture *Texture.CubeMap) (*Skybox, error) {
	mesh, err := Model.NewMesh(vertieces, Buffer.RegisterVertexBufferAttributes, indices)
	if err != nil {
		return nil, err
	}

	return &Skybox{
		Mesh:    mesh,
		CubeMap: texture,
	}, nil
}

func (skybox *Skybox) Draw(shader Shader.IShaderProgram) error {
	err := shader.BindObject(skybox.CubeMap)
	skybox.Mesh.Draw()

	return err
}

func (skybox *Skybox) GetPosition() *Vector.Vector3 {
	return &Vector.Vector3{0, 0, 0}
}
