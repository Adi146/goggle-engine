package Terrain

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Model"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Core/VertexBuffer"
	"gopkg.in/yaml.v3"
)

type Terrain struct {
	Model.Model
	TileSize  float32
	HeightMap HeightMap

	OffsetX float32
	OffsetZ float32
}

func GenerateTerrain(heightMap HeightMap, tileSize float32) (*Terrain, error) {
	vertices := make([]VertexBuffer.Vertex, heightMap.NumRows*heightMap.NumColumns)
	indices := make([]uint32, 6*(heightMap.NumRows-1)*(heightMap.NumColumns-1))

	offsetX := float32(heightMap.NumColumns)/2 - 0.5
	offsetZ := float32(heightMap.NumRows)/2 - 0.5

	for z := 0; z < heightMap.NumRows; z++ {
		for x := 0; x < heightMap.NumColumns; x++ {
			vertices[z*heightMap.NumColumns+x] = VertexBuffer.Vertex{
				Position: GeometryMath.Vector3{
					(float32(x) - offsetX) * tileSize,
					heightMap.GetHeight(x, z),
					(float32(z) - offsetZ) * tileSize,
				},
				Normal:    heightMap.GetNormal(x, z),
				UV:        GeometryMath.Vector2{float32(x) / float32(heightMap.NumColumns), float32(z) / float32(heightMap.NumRows)},
				Tangent:   heightMap.GetTangent(x, z),
				BiTangent: heightMap.GetBiTangent(x, z),
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

	mesh, err := Model.NewMesh(vertices, indices)
	if err != nil {
		return nil, err
	}

	return &Terrain{
		Model: Model.Model{
			Mesh:     *mesh,
			Material: nil,
		},
		TileSize:  tileSize,
		HeightMap: heightMap,
		OffsetX:   offsetX,
		OffsetZ:   offsetZ,
	}, nil
}

func (terrain *Terrain) GetHeightAt(terrainPos GeometryMath.Vector3) float32 {
	terrainPos[0] += terrain.OffsetX
	terrainPos[2] += terrain.OffsetZ

	gridX := int(GeometryMath.Floor(terrainPos.X() / terrain.TileSize))
	gridZ := int(GeometryMath.Floor(terrainPos.Z() / terrain.TileSize))

	if gridX < 0 || gridX >= terrain.HeightMap.NumColumns || gridZ < 0 || gridZ >= terrain.HeightMap.NumRows {
		return 0
	}

	xCoordOnTile := GeometryMath.Mod(terrainPos.X(), terrain.TileSize) / terrain.TileSize
	zCoordOnTile := GeometryMath.Mod(terrainPos.Z(), terrain.TileSize) / terrain.TileSize

	var answer float32
	if xCoordOnTile <= 1-zCoordOnTile {
		answer = barryCentric(
			GeometryMath.Vector3{0, terrain.HeightMap.GetHeight(gridX, gridZ), 0},
			GeometryMath.Vector3{1, terrain.HeightMap.GetHeight(gridX+1, gridZ), 0},
			GeometryMath.Vector3{0, terrain.HeightMap.GetHeight(gridX, gridZ+1), 1},
			GeometryMath.Vector2{xCoordOnTile, zCoordOnTile})
	} else {
		answer = barryCentric(
			GeometryMath.Vector3{1, terrain.HeightMap.GetHeight(gridX+1, gridZ), 0},
			GeometryMath.Vector3{1, terrain.HeightMap.GetHeight(gridX+1, gridZ+1), 1},
			GeometryMath.Vector3{0, terrain.HeightMap.GetHeight(gridX, gridZ+1), 1},
			GeometryMath.Vector2{xCoordOnTile, zCoordOnTile})
	}

	return answer
}

func barryCentric(p1, p2, p3 GeometryMath.Vector3, pos GeometryMath.Vector2) float32 {
	det := (p2.Z()-p3.Z())*(p1.X()-p3.X()) + (p3.X()-p2.X())*(p1.Z()-p3.Z())
	l1 := ((p2.Z()-p3.Z())*(pos.X()-p3.X()) + (p3.X()-p2.X())*(pos.Y()-p3.Z())) / det
	l2 := ((p3.Z()-p1.Z())*(pos.X()-p3.X()) + (p1.X()-p3.X())*(pos.Y()-p3.Z())) / det
	l3 := 1.0 - l1 - l2
	return l1*p1.Y() + l2*p2.Y() + l3*p3.Y()
}

func (terrain *Terrain) UnmarshalYAML(value *yaml.Node) error {
	var yamlConfig struct {
		HeightMap HeightMap `yaml:",inline"`
		TileSize  float32   `yaml:"tileSize"`
		Material  struct {
			BlendMaterial     Material.BlendMaterial `yaml:",inline"`
			BlendMapGenerator BlendMapGenerator      `yaml:"blendRange"`
		} `yaml:"material"`
	}
	if err := value.Decode(&yamlConfig); err != nil {
		return err
	}

	tmpTerrain, err := GenerateTerrain(yamlConfig.HeightMap, yamlConfig.TileSize)
	if err != nil {
		return err
	}

	if yamlConfig.Material.BlendMaterial.BlendMap == nil {
		blendMap, err := yamlConfig.Material.BlendMapGenerator.GenerateBlendMap(&yamlConfig.HeightMap)
		if err != nil {
			return err
		}
		yamlConfig.Material.BlendMaterial.BlendMap = blendMap
	}

	yamlConfig.Material.BlendMaterial.SetWrapMode(Texture.Repeat)
	yamlConfig.Material.BlendMaterial.GenerateMibMaps(-1)

	*terrain = *tmpTerrain
	terrain.Material = &yamlConfig.Material.BlendMaterial
	terrain.SetModelMatrix(GeometryMath.Identity())

	return nil
}
