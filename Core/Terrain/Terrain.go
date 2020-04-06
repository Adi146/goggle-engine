package Terrain

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Core/VertexBuffer"
	"gopkg.in/yaml.v3"
)

type Terrain Model.Model

func GenerateTerrain(heightMap HeightMap, tileSize float32, uvScale float32) (*Model.Mesh, error) {
	vertices := make([]VertexBuffer.Vertex, heightMap.NumRows*heightMap.NumColumns)
	indices := make([]uint32, 6*(heightMap.NumRows-1)*(heightMap.NumColumns-1))

	offsetZ := float32(heightMap.NumRows)/2 - 0.5
	offsetX := float32(heightMap.NumColumns)/2 - 0.5

	for z := 0; z < heightMap.NumRows; z++ {
		for x := 0; x < heightMap.NumColumns; x++ {
			vertices[z*heightMap.NumColumns+x] = VertexBuffer.Vertex{
				Position: GeometryMath.Vector3{
					(float32(x) - offsetX) * tileSize,
					heightMap.GetHeight(x, z),
					(float32(z) - offsetZ) * tileSize,
				},
				Normal:  heightMap.GetNormal(x, z),
				UV:      GeometryMath.Vector2{float32(x) / uvScale, float32(z) / uvScale},
				Tangent: GeometryMath.Vector3{},
			}
		}
	}

	for z := 0; z < heightMap.NumRows-1; z++ {
		for x := 0; x < heightMap.NumColumns-1; x++ {
			topLeft := uint32(z*heightMap.NumColumns + x)
			topRight := uint32(z*heightMap.NumColumns + x + 1)
			bottomLeft := uint32((z+1)*heightMap.NumColumns + x)
			bottomRight := uint32((z+1)*heightMap.NumColumns + x + 1)

			indices[(z*(heightMap.NumColumns-1)+x)*6+0] = topLeft
			indices[(z*(heightMap.NumColumns-1)+x)*6+1] = bottomLeft
			indices[(z*(heightMap.NumColumns-1)+x)*6+2] = topRight
			indices[(z*(heightMap.NumColumns-1)+x)*6+3] = topRight
			indices[(z*(heightMap.NumColumns-1)+x)*6+4] = bottomLeft
			indices[(z*(heightMap.NumColumns-1)+x)*6+5] = bottomRight
		}
	}

	return Model.NewMesh(vertices, VertexBuffer.RegisterVertexBufferAttributes, indices)
}

func (terrain *Terrain) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	return (*Model.Model)(terrain).Draw(shader, invoker, scene)
}

func (terrain *Terrain) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		HeightMap HeightMap         `yaml:",inline"`
		TileSize  float32           `yaml:"tileSize"`
		UvScale   float32           `yaml:"uvScale"`
		Material  Material.Material `yaml:"material"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	mesh, err := GenerateTerrain(yamlConfig.HeightMap, yamlConfig.TileSize, yamlConfig.UvScale)
	if err != nil {
		return err
	}

	yamlConfig.Material.DiffuseTexture.SetWrapMode(Texture.Repeat)
	yamlConfig.Material.DiffuseTexture.GenerateMipMap(-1)

	*terrain = Terrain{
		Meshes: []Model.MeshesWithMaterial{
			{
				Mesh:     mesh,
				Material: &yamlConfig.Material,
			},
		},
		ModelMatrix: GeometryMath.Identity(),
	}

	return nil
}
