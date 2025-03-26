package check

import (
	"fmt"
	"reflect"
)

func TypeOf(x any) string {
	if IsPointer(x) {
		return TypeOf(Deref(x))
	}
	return reflect.TypeOf(x).Name()
}

func Deref(x any) any {
	return reflect.ValueOf(x).Elem().Interface()
}

func AddressOf(x any) string {
	return fmt.Sprintf("%p", x)
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
	return IsStruct(Deref(x))
}
