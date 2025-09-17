package check

import "github.com/go-playground/validator/v10"

var validate = validator.New(validator.WithRequiredStructEnabled())

func IsValidStruct[T any](item *T) bool {
	if item == nil {
		return false
	}
	err := validate.Struct(item)
	return err == nil
}
