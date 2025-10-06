package dyn

import (
	"fmt"
	"reflect"
)

// Get item.field, where item is a *struct
func GetFieldValue(item any, field string) any {
	if !IsStructPointer(item) {
		return nil
	}
	return reflect.ValueOf(item).Elem().FieldByName(field).Interface()
}

// Get item.field, where item is a *struct,
// Returns the field as the given type, if valid
func GetField[T any](item any, field string) (T, bool) {
	rawValue := GetFieldValue(item, field)
	value, ok := rawValue.(T)
	return value, ok
}

// Get item.field, where item is a *struct, and field value is a string
func GetFieldString(item any, field string) string {
	value := GetFieldValue(item, field)
	return fmt.Sprintf("%v", value)
}

// Set item.field, where item is a *struct
func SetFieldValue(item any, field string, value any) {
	if !IsStructPointer(item) {
		return
	}
	reflect.ValueOf(item).Elem().FieldByName(field).Set(reflect.ValueOf(value))
}
