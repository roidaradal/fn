package dyn

import (
	"fmt"
	"reflect"
)

func TypeOf(x any) string {
	if IsPointer(x) {
		return TypeOf(deref(x))
	}
	return reflect.TypeOf(x).Name()
}

func AddressOf(x any) string {
	return fmt.Sprintf("%p", x)
}

func IsZero(x any) bool {
	return reflect.ValueOf(x).IsZero()
}

func IsPointer(x any) bool {
	return reflect.TypeOf(x).Kind() == reflect.Pointer
}

func IsStruct(x any) bool {
	return reflect.TypeOf(x).Kind() == reflect.Struct
}

func IsStructPointer(x any) bool {
	if !IsPointer(x) {
		return false
	}
	return IsStruct(deref(x))
}

func deref(x any) any {
	return reflect.ValueOf(x).Elem().Interface()
}
