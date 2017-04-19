package goutil

import (
	"reflect"
)

func typeIsStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func typeIsFunc(t reflect.Type) bool {
	return t.Kind() == reflect.Func
}

func typeIsSlice(t reflect.Type) bool {
	return t.Kind() == reflect.Slice
}

func typeIsMap(t reflect.Type) bool {
	return t.Kind() == reflect.Map
}

func isZero(v reflect.Value) bool {
	return v.Kind() == reflect.Invalid
}

func isStruct(v reflect.Value) bool {
	return v.Kind() == reflect.Struct
}

func isFunc(v reflect.Value) bool {
	return v.Kind() == reflect.Func
}

func isSlice(v reflect.Value) bool {
	return v.Kind() == reflect.Slice
}

func isMap(v reflect.Value) bool {
	return v.Kind() == reflect.Map
}

func isPtr(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr
}
