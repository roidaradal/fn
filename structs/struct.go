package structs

import (
	"encoding/json"
	"reflect"

	"github.com/roidaradal/fn/check"
)

func GetFieldValue(x any, field string) any {
	if !check.IsStructPointer(x) {
		return nil
	}
	return reflect.ValueOf(x).Elem().FieldByName(field).Interface()
}

func SetFieldValue(x any, field string, value any) {
	if !check.IsStructPointer(x) {
		return
	}
	reflect.ValueOf(x).Elem().FieldByName(field).Set(reflect.ValueOf(value))
}

func ToMap[T any, V any](item *T) (map[string]V, error) {
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
