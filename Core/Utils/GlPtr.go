package Utils

import (
	"fmt"
	"reflect"
	"unsafe"
)

func GlPtr(data interface{}) unsafe.Pointer {
	if data == nil {
		return unsafe.Pointer(nil)
	}

	switch v := data.(type) {
	case reflect.Value:
		return glPtr(v)
	default:
		return glPtr(reflect.ValueOf(data))
	}
}

func glPtr(v reflect.Value) unsafe.Pointer {
	var addr unsafe.Pointer

	switch v.Type().Kind() {
	case reflect.Ptr:
		e := v.Elem()
		switch e.Kind() {
		case
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			addr = unsafe.Pointer(e.UnsafeAddr())
		case reflect.Array, reflect.Slice:
			addr = unsafe.Pointer(e.Index(0).UnsafeAddr())
		case reflect.Struct:
			addr = unsafe.Pointer(e.Field(0).UnsafeAddr())
		default:
			panic(fmt.Errorf("unsupported pointer to type %s; must be a slice or pointer to a singular scalar value or the first element of an array or slice", e.Kind()))
		}
	case reflect.Uintptr:
		addr = unsafe.Pointer(v.Pointer())
	case reflect.Array, reflect.Slice:
		addr = unsafe.Pointer(v.Index(0).UnsafeAddr())
	case reflect.Struct:
		addr = unsafe.Pointer(v.Field(0).UnsafeAddr())
	default:
		panic(fmt.Errorf("unsupported type %s; must be a slice or pointer to a singular scalar value or the first element of an array or slice", v.Type()))
	}

	return addr
}
