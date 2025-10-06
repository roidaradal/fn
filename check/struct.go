package check

import "github.com/go-playground/validator/v10"

var validate = validator.New(validator.WithRequiredStructEnabled())

// Validate struct based on Go JSON validators
func IsValidStruct[T any](item *T) bool {
	if item == nil {
		return false
	}
	err := validate.Struct(item)
	return err == nil
}

// Check if pointer is nil
func IsNull[T any](item *T) bool {
	return item == nil
}

// Check if pointer is not nil
func NotNull[T any](item *T) bool {
	return item != nil
}
