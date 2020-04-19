package Terrain

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Model/Material"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Texture"
	"github.com/Adi146/goggle-engine/Utils/Error"
	"gopkg.in/yaml.v3"
)

const (
	maxChunkRows    = 64
	maxChunkColumns = 64
)

type Terrain struct {
	Chunks   []Chunk
	Material *Material.BlendMaterial

	TileSize  float32
	HeightMap HeightMap

	OffsetX float32
	OffsetZ float32
}

func GenerateTerrain(heightMap HeightMap, tileSize float32) (*Terrain, error) {
	offsetX := float32(heightMap.NumColumns)/2 - 0.5
	offsetZ := float32(heightMap.NumRows)/2 - 0.5

	var chunks []Chunk
	for z := 0; z < heightMap.NumRows; z += maxChunkRows - 1 {
		chunkRows := maxChunkRows
		if z+chunkRows > heightMap.NumRows {
			chunkRows = heightMap.NumRows - z
		}
		for x := 0; x < heightMap.NumColumns; x += maxChunkColumns - 1 {
			chunkColumns := maxChunkColumns
			if x+chunkColumns > heightMap.NumColumns {
				chunkColumns = heightMap.NumColumns - x
			}

			chunks = append(chunks, generateChunk(&heightMap, x, z, chunkRows, chunkColumns, tileSize))
		}
	}

	return &Terrain{
		Chunks:    chunks,
		TileSize:  tileSize,
		HeightMap: heightMap,
		OffsetX:   offsetX,
		OffsetZ:   offsetZ,
	}, nil
}

func (terrain *Terrain) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	var err Error.ErrorCollection

	err.Push(shader.BindObject(terrain.Material))
	for _, chunk := range terrain.Chunks {
		err.Push(chunk.Draw(shader, invoker, scene))
	}
	terrain.Material.Unbind()

	return err.Err()
}

func (terrain *Terrain) SetModelMatrix(mat GeometryMath.Matrix4x4) {
	for _, chunk := range terrain.Chunks {
		chunk.SetModelMatrix(mat)
	}
}

func (terrain *Terrain) GetHeightAt(terrainPos GeometryMath.Vector3) float32 {
	terrainPos[0] += terrain.OffsetX * terrain.TileSize
	terrainPos[2] += terrain.OffsetZ * terrain.TileSize

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
			GeometryMath.Vector3{0, terrain.HeightMap.GetHeightScaled(gridX, gridZ), 0},
			GeometryMath.Vector3{1, terrain.HeightMap.GetHeightScaled(gridX+1, gridZ), 0},
			GeometryMath.Vector3{0, terrain.HeightMap.GetHeightScaled(gridX, gridZ+1), 1},
			GeometryMath.Vector2{xCoordOnTile, zCoordOnTile})
	} else {
		answer = barryCentric(
			GeometryMath.Vector3{1, terrain.HeightMap.GetHeightScaled(gridX+1, gridZ), 0},
			GeometryMath.Vector3{1, terrain.HeightMap.GetHeightScaled(gridX+1, gridZ+1), 1},
			GeometryMath.Vector3{0, terrain.HeightMap.GetHeightScaled(gridX, gridZ+1), 1},
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

	*terrain = *tmpTerrain
	terrain.Material = &yamlConfig.Material.BlendMaterial

	return nil
}
