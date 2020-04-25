package DirectionalLight

import (
	"github.com/Adi146/goggle-engine/Core/BoundingVolume"
	core "github.com/Adi146/goggle-engine/Core/Camera"
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
	"github.com/Adi146/goggle-engine/Core/UniformBuffer/UniformBufferSection"
)

const (
	near_plane = 0.1
	offset     = 15
)

type Camera struct {
	core.Camera
	UniformBufferSection.Matrix4x4
}

func (camera *Camera) Update(sceneCamera core.ICamera, direction GeometryMath.Vector3, distance float32) {
	boundingBox, center := camera.calcCameraFrustumBoundingBox(sceneCamera, direction, distance)

	projectionMatrix := GeometryMath.Orthographic(boundingBox.Min.X(), boundingBox.Max.X(), boundingBox.Min.Y(), boundingBox.Max.Y(), boundingBox.Min.Z(), boundingBox.Max.Z())
	viewMatrix := GeometryMath.LookAt(center.Add(direction.Invert()), center, GeometryMath.Vector3{0, 1, 0})
	camera.Set(projectionMatrix.Mul(viewMatrix))
}

func (camera *Camera) calcCameraFrustumBoundingBox(sceneCamera core.ICamera, direction GeometryMath.Vector3, distance float32) (BoundingVolume.AABB, GeometryMath.Vector3) {
	tmpViewMatrix := GeometryMath.LookAt(direction.Invert(), GeometryMath.Vector3{0, 0, 0}, GeometryMath.Vector3{0, 1, 0})

	frustumPoints := camera.calcCameraFrustumPoints(sceneCamera, distance)
	for i := range frustumPoints {
		frustumPoints[i] = tmpViewMatrix.MulVector(frustumPoints[i])
	}
	boundingBox := BoundingVolume.NewAABB(frustumPoints[:])
	boundingBox.Max[2] += offset

	return boundingBox, tmpViewMatrix.Inverse().MulVector(boundingBox.GetCenter())
}

func (camera *Camera) calcCameraFrustumPoints(sceneCamera core.ICamera, distance float32) [8]GeometryMath.Vector3 {
	position := sceneCamera.GetPosition()

	farWidth, farHeight := sceneCamera.GetFrustum().GetProjectionConfig().GetPlane(distance)
	nearWidth, nearHeight := sceneCamera.GetFrustum().GetProjectionConfig().GetPlane(near_plane)

	front := sceneCamera.GetFront()
	up := sceneCamera.GetUp()
	right := sceneCamera.GetRight()
	down := up.Invert()
	left := right.Invert()

	centerFar := position.Add(front.MulScalar(distance))
	centerNear := position.Add(front.MulScalar(near_plane))

	farTop := centerFar.Add(up.MulScalar(farHeight))
	farBottom := centerFar.Add(down.MulScalar(farHeight))
	nearTop := centerNear.Add(up.MulScalar(nearHeight))
	nearBottom := centerNear.Add(down.MulScalar(nearHeight))

	return [8]GeometryMath.Vector3{
		farTop.Add(right.MulScalar(farWidth)),
		farTop.Add(left.MulScalar(farWidth)),
		farBottom.Add(right.MulScalar(farWidth)),
		farBottom.Add(left.MulScalar(farWidth)),
		nearTop.Add(right.MulScalar(nearWidth)),
		nearTop.Add(left.MulScalar(nearWidth)),
		nearBottom.Add(right.MulScalar(nearWidth)),
		nearBottom.Add(left.MulScalar(nearWidth)),
	}
}
