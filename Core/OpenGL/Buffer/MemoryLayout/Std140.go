package MemoryLayout

import (
	"github.com/Adi146/goggle-engine/Core/GeometryMath"
)

type Padding float32

type Std140Vector3 struct {
	GeometryMath.Vector3
	padding [1]Padding
}

type Std140Matrix4x4 struct {
	GeometryMath.Matrix4x4
}

type Std140Float32 struct {
	Float32 float32
	padding [3]Padding
}
