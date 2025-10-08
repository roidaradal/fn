package dyn

import (
	"fmt"
	"reflect"
)

// Return the full type name of given item (*Type for pointers)
func FullTypeOf(x any) string {
	if IsPointer(x) {
		return "*" + TypeOf(Deref(x))
	}
	return reflect.TypeOf(x).Name()
}

// Return the base type name of given item (dereferences pointers)
func TypeOf(x any) string {
	if IsPointer(x) {
		return TypeOf(Deref(x))
	}
	return reflect.TypeOf(x).Name()
}

// Check if given item has zero value
func IsZero(x any) bool {
	return reflect.ValueOf(x).IsZero()
}

// Check if given item is nil
func IsNull(x any) bool {
	if x == nil {
		return true
	}
	switch reflect.TypeOf(x).Kind() {
	case reflect.Pointer, reflect.Map, reflect.Array, reflect.Slice, reflect.Chan:
		return reflect.ValueOf(x).IsNil()
	}
	return false
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
	return IsStruct(Deref(x))
}

// Returns the memory address of given item as string
func AddressOf(x any) string {
	return fmt.Sprintf("%p", x)
}

// Internal: De-references the pointer,
// Assumed that given item is already confirmed to be a pointer
func Deref(x any) any {
	return reflect.ValueOf(x).Elem().Interface()
}
