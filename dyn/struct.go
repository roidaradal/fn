package dyn

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func GetFieldValue(x any, field string) any {
	if !IsStructPointer(x) {
		return nil
	}
	return reflect.ValueOf(x).Elem().FieldByName(field).Interface()
}

func GetFieldString(x any, field string) string {
	value := GetFieldValue(x, field)
	return fmt.Sprintf("%v", value)
}

func SetFieldValue(x any, field string, value any) {
	if !IsStructPointer(x) {
		return
	}
	reflect.ValueOf(x).Elem().FieldByName(field).Set(reflect.ValueOf(value))
}

func StructToMap[T any, V any](item *T) (map[string]V, error) {
	var output map[string]V
	data, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
