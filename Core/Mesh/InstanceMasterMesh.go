package Mesh

import (
	"github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer"
	"github.com/Adi146/goggle-engine/Core/Scene"
	"github.com/Adi146/goggle-engine/Core/Shader"
	"github.com/Adi146/goggle-engine/Utils/Error"

	"github.com/go-gl/gl/v4.3-core/gl"
)

type InstanceMasterMesh struct {
	Mesh
	MasterMatrix GeometryMath.Matrix4x4
	MatrixBuffer Buffer.Buffer
	Instances    []*InstancedMesh
}

func NewInstanceMasterMesh(mesh *Mesh, matrices ...GeometryMath.Matrix4x4) *InstanceMasterMesh {
	allMatrices := append([]GeometryMath.Matrix4x4{mesh.GetModelMatrix()}, matrices...)
	mbo := Buffer.NewStaticArrayBuffer(&allMatrices)
	mesh.VertexArray = NewInstancedVertexArray(mesh.GetVertexArray(), mbo)

	master := InstanceMasterMesh{
		Mesh:         *mesh,
		MasterMatrix: GeometryMath.Identity(),
		MatrixBuffer: mbo,
	}

	instances := make([]*InstancedMesh, len(matrices))
	for i := range matrices {
		instances[i] = &InstancedMesh{
			ModelMatrix:               matrices[i],
			boundingVolume:            mesh.GetBoundingVolume(),
			TransformedBoundingVolume: mesh.GetBoundingVolume().Transform(matrices[i].Mul(master.MasterMatrix)),
			FrustumCulling:            mesh.FrustumCulling,
			Master:                    &master,
		}
	}

	master.Instances = instances
	return &master
}

func (mesh *InstanceMasterMesh) Draw(shader Shader.IShaderProgram, invoker Scene.IDrawable, scene Scene.IScene, camera Camera.ICamera) error {
	var err Error.ErrorCollection

	var matrices []GeometryMath.Matrix4x4
	if !mesh.FrustumCulling || (mesh.FrustumCulling && camera.GetFrustum().Contains(mesh.GetBoundingVolume())) {
		matrices = append(matrices, mesh.GetModelMatrix())
	}

	for _, instance := range mesh.Instances {
		if !instance.FrustumCulling || (instance.FrustumCulling && camera.GetFrustum().Contains(instance.GetBoundingVolume())) {
			matrices = append(matrices, instance.GetModelMatrix())
		}
	}

	if len(matrices) == 0 {
		return nil
	}

	mesh.MatrixBuffer.Set(&matrices)
	mesh.MatrixBuffer.Sync()

	err.Push(shader.BindObject(&mesh.MasterMatrix))
	err.Push(shader.BindObject(mesh.VertexArray))
	mesh.IndexBuffer.Bind()
	gl.DrawElementsInstanced(gl.TRIANGLES, mesh.IndexBuffer.Length, gl.UNSIGNED_INT, nil, int32(len(matrices)))
	mesh.IndexBuffer.Unbind()

	return err.Err()
}

func (mesh *InstanceMasterMesh) SetMasterMatrix(mat GeometryMath.Matrix4x4) {
	mesh.MasterMatrix = mat
	mesh.TransformedBoundingVolume = mesh.boundingVolume.Transform(mesh.GetModelMatrix().Mul(mat))

	for _, instance := range mesh.Instances {
		instance.TransformedBoundingVolume = instance.boundingVolume.Transform(instance.GetModelMatrix().Mul(mat))
	}
}

func (mesh *InstanceMasterMesh) CreateNewInstances(matrices ...GeometryMath.Matrix4x4) []*InstancedMesh {
	instances := make([]*InstancedMesh, len(matrices))
	for i := range matrices {
		instances[i] = &InstancedMesh{
			ModelMatrix:               matrices[i],
			boundingVolume:            mesh.GetBoundingVolume(),
			TransformedBoundingVolume: mesh.GetBoundingVolume().Transform(matrices[i].Mul(mesh.MasterMatrix)),
			FrustumCulling:            mesh.FrustumCulling,
			Master:                    mesh,
		}
	}

	mesh.Instances = append(mesh.Instances, instances...)

	return instances
}
