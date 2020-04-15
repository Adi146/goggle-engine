package Mesh

import (
	"fmt"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Core/Utils"
	"github.com/Adi146/goggle-engine/Utils/Error"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type InstancedMesh struct {
	Mesh
	InstanceID   int32
	MatrixBuffer ArrayBuffer
	NumInstances int32
}

func NewInstancedMeshes(mesh IMesh, matrices ...GeometryMath.Matrix4x4) ([]*InstancedMesh, error) {
	if instancedMesh, isInstancedMesh := mesh.(*InstancedMesh); isInstancedMesh {
		newInstanced, err := instancedMesh.CreateNewInstances(matrices)
		if err != nil {
			return nil, err
		}

		return append([]*InstancedMesh{instancedMesh}, newInstanced...), nil
	} else {
		matrices = append([]GeometryMath.Matrix4x4{mesh.GetModelMatrix()}, matrices ...)
		mbo := NewMatrixBuffer(matrices)

		instances := make([]*InstancedMesh, len(matrices))
		for i := range matrices {
			instances[i] = &InstancedMesh{
				Mesh: Mesh{
					VertexBuffer: mesh.GetVertexBuffer(),
					VertexArray:  NewInstancedVertexArray(mesh.GetVertexArray(), mbo),
					IndexBuffer:  mesh.GetIndexBuffer(),
				},
				MatrixBuffer: mbo,
				NumInstances: int32(len(matrices)),
				InstanceID:   int32(i),
			}
		}

		return instances, nil
	}
}

func (mesh *InstancedMesh) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene) error {
	if mesh.InstanceID != 0 {
		return fmt.Errorf("draw can not be called from subinstances (instancedID %d)", mesh.InstanceID)
	}

	var err Error.ErrorCollection

	err.Push(shader.BindObject(GeometryMath.Identity()))
	err.Push(shader.BindObject(mesh.VertexArray))
	mesh.IndexBuffer.Bind()
	gl.DrawElementsInstanced(gl.TRIANGLES, mesh.IndexBuffer.Length, gl.UNSIGNED_INT, nil, mesh.NumInstances)
	mesh.IndexBuffer.Unbind()
	mesh.VertexBuffer.Unbind()

	return err.Err()
}

func (mesh *InstancedMesh) SetModelMatrix(mat GeometryMath.Matrix4x4) {
	mesh.Mesh.SetModelMatrix(mat)
	mesh.MatrixBuffer.UpdateData(&mat, int(mesh.InstanceID)*Utils.SizeOf(mat))
}

func (mesh *InstancedMesh) CreateNewInstances(matrices... []GeometryMath.Matrix4x4) ([]*InstancedMesh, error) {
	if mesh.InstanceID != 0 {
		return nil, fmt.Errorf("only master instance can create subinstances (instancedID %d)", mesh.InstanceID)
	}

	firstIndex := mesh.NumInstances
	mesh.MatrixBuffer.IncreaseSize(len(matrices) * Utils.SizeOf(GeometryMath.Matrix4x4{}))
	mesh.NumInstances += int32(len(matrices))

	instances := make([]*InstancedMesh, len(matrices))
	for i := range matrices {
		instances[i] = &InstancedMesh{
			Mesh: Mesh{
				VertexBuffer: mesh.GetVertexBuffer(),
				VertexArray:  mesh.GetVertexArray(),
				IndexBuffer:  mesh.GetIndexBuffer(),
			},
			MatrixBuffer: mesh.MatrixBuffer,
			NumInstances: mesh.NumInstances,
			InstanceID:  firstIndex + int32(i),
		}
	}

	return instances, nil
}
