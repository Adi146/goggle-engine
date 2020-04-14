package Skybox

import (
	"github.com/Adi146/goggle-engine/Core/Function"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Mesh"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
)

var (
	vertices = []Mesh.Vertex{
		{Position: GeometryMath.Vector3{-1.0, 1.0, 1.0}},
		{Position: GeometryMath.Vector3{-1.0, 1.0, -1.0}},
		{Position: GeometryMath.Vector3{-1.0, -1.0, 1.0}},
		{Position: GeometryMath.Vector3{-1.0, -1.0, -1.0}},

		{Position: GeometryMath.Vector3{1.0, 1.0, 1.0}},
		{Position: GeometryMath.Vector3{1.0, 1.0, -1.0}},
		{Position: GeometryMath.Vector3{1.0, -1.0, 1.0}},
		{Position: GeometryMath.Vector3{1.0, -1.0, -1.0}},
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
	*Mesh.Mesh
	*Texture.CubeMap
}

func NewSkybox(texture *Texture.CubeMap) (*Skybox, error) {
	return &Skybox{
		Mesh:    Mesh.NewMesh(vertices, indices),
		CubeMap: texture,
	}, nil
}

func (skybox *Skybox) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	defer Function.GetCurrentDepthFunction().Set()
	Function.LessEqual.Set()

	err := shader.BindObject(skybox.CubeMap)
	skybox.Mesh.Draw(shader, nil, nil)
	skybox.CubeMap.Unbind()

	return err
}
