package check

import "reflect"

func SetStructFieldValue(x any, field string, value any) {
	if !IsStructPointer(x) {
		return
	}
	reflect.ValueOf(x).Elem().FieldByName(field).Set(reflect.ValueOf(value))
}
