package Utils

import (
	"fmt"
	"reflect"
)

func SizeOf(data interface{}) int {
	if data == nil {
		return 0
	}

	v := reflect.ValueOf(data)
	return sizeOf(v)
}

func sizeOf(v reflect.Value) int {
	var size int

	switch t := v.Type(); t.Kind() {
	case reflect.Array, reflect.Slice:
		size = int(t.Elem().Size()) * v.Len()
	case reflect.Struct:
		for i, n := 0, t.NumField(); i < n; i++ {
			size += sizeOf(v.Field(i))
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		size = int(t.Size())
	case reflect.Ptr:
		size = sizeOf(v.Elem())
	default:
		panic(fmt.Errorf("unsupported type %s", t))
	}

	return size
}
