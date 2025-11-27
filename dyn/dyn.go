// Package dyn contains shortcuts for common reflect package steps.
package dyn

import (
	"fmt"
	"reflect"
)

// Check if two any values are equal
func Equal(item1, item2 any) bool {
	// Dereference item1 if pointer and not null
	if IsPointer(item1) && NotNull(item1) {
		return Equal(Deref(item1), item2)
	}
	// Dereference item2 if pointer and not null
	if IsPointer(item2) && NotNull(item2) {
		return Equal(item1, Deref(item2))
	}
	return item1 == item2
}

// Check if two any values are not equal
func NotEqual(item1, item2 any) bool {
	return !Equal(item1, item2)
}

// Dereference the pointer (assumed that item already confirmed to be pointer)
func Deref(x any) any {
	return reflect.ValueOf(x).Elem().Interface()
}

// Return memory address of given item as string
func AddressOf(x any) string {
	return fmt.Sprintf("%p", x)
}

// Check if given item has zero value
func IsZero(x any) bool {
	return reflect.ValueOf(x).IsZero()
}

// Check if given item is null
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

// Check if given item is not null
func NotNull(x any) bool {
	return !IsNull(x)
}

// Check if given item is a pointer
func IsPointer(x any) bool {
	return reflect.TypeOf(x).Kind() == reflect.Pointer
}

// Check if given item is a struct
func IsStruct(x any) bool {
	return reflect.TypeOf(x).Kind() == reflect.Struct
}

// Check if given item is pointer to a struct
func IsStructPointer(x any) bool {
	if !IsPointer(x) {
		return false
	}
	return IsStruct(Deref(x))
}

// Return base type name of given item (dereferences pointers)
func TypeOf(x any) string {
	if IsPointer(x) {
		return TypeOf(Deref(x))
	}
	return reflect.TypeOf(x).Name()
}

// Return full type name of given item (*Type for pointers)
func FullTypeOf(x any) string {
	if IsPointer(x) {
		return "*" + FullTypeOf(Deref(x))
	}
	return reflect.TypeOf(x).Name()
}
