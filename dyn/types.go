package dyn

import (
	"fmt"
	"reflect"
)

// Return the type name of given item
func TypeOf(x any) string {
	if IsPointer(x) {
		return "*" + TypeOf(deref(x))
	}
	return reflect.TypeOf(x).Name()
}

// Check if given item has zero value
func IsZero(x any) bool {
	return reflect.ValueOf(x).IsZero()
}

// Check if given item is a pointer
func IsPointer(x any) bool {
	return reflect.TypeOf(x).Kind() == reflect.Pointer
}

// Check if given item is a struct
func IsStruct(x any) bool {
	return reflect.TypeOf(x).Kind() == reflect.Struct
}

// Check if given item is a pointer to a struct
func IsStructPointer(x any) bool {
	if !IsPointer(x) {
		return false
	}
	return IsStruct(deref(x))
}

// Returns the memory address of given item as string
func AddressOf(x any) string {
	return fmt.Sprintf("%p", x)
}

// Internal: De-references the pointer,
// Assumed that given item is already confirmed to be a pointer
func deref(x any) any {
	return reflect.ValueOf(x).Elem().Interface()
}
