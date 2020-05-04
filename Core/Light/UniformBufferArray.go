package Light

import (
	"github.com/Adi146/goggle-engine/Core/OpenGL/Buffer"
	"reflect"
)

type UniformBufferArray struct {
	array []Buffer.IDynamicBufferData
}

func (array *UniformBufferArray) Add(data Buffer.IDynamicBufferData) int {
	array.array = append(array.array, data)

	return len(array.array) - 1
}

func (array *UniformBufferArray) GetBufferData() interface{} {
	if len(array.array) == 0 {
		return nil
	}

	t := reflect.TypeOf(array.array[0].GetBufferData())
	structType := reflect.StructOf([]reflect.StructField{
		{
			Name: "Len",
			Type: reflect.TypeOf(int32(0)),
		},
		{
			Name: "Padding",
			Type: reflect.ArrayOf(3, reflect.TypeOf(int32(0))),
		},
		{
			Name: "Array",
			Type: reflect.ArrayOf(len(array.array), t),
		},
	})

	val := reflect.New(structType).Elem()
	val.Field(0).SetInt(int64(len(array.array)))

	for i, element := range array.array {
		val.Field(2).Index(i).Set(reflect.ValueOf(element.GetBufferData()))
	}

	return val
}

func (array *UniformBufferArray) IsSync() bool {
	for _, element := range array.array {
		if !element.IsSync() {
			return false
		}
	}

	return true
}

func (array *UniformBufferArray) SetIsSync(val bool) {
	for _, element := range array.array {
		element.SetIsSync(val)
	}
}
