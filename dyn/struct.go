package dyn

import (
	"fmt"
	"reflect"
)

// Get item.field from structRef
func GetFieldValue(structRef any, field string) any {
	if !IsStructPointer(structRef) {
		return nil
	}
	return reflect.ValueOf(structRef).Elem().FieldByName(field).Interface()
}

// Get item.field from structRef, and type coerce into T
func GetField[T any](structRef any, field string) (T, bool) {
	rawValue := GetFieldValue(structRef, field)
	value, ok := rawValue.(T)
	return value, ok
}

// Get item.field from structRef, return field value as string
func GetFieldString(structRef any, field string) string {
	return fmt.Sprintf("%v", GetFieldValue(structRef, field))
}

// Set item.field = value for structRef
func SetFieldValue(structRef any, field string, value any) {
	if !IsStructPointer(structRef) {
		return
	}
	reflect.ValueOf(structRef).Elem().FieldByName(field).Set(reflect.ValueOf(value))
}
