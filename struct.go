package fn

import "encoding/json"

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
