package check

import "github.com/go-playground/validator/v10"

type CustomValidatorFn = func(validator.FieldLevel) bool

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

// Create custom string validator function
func NewStringValidator(validatorFn func(string) bool) CustomValidatorFn {
	return func(fl validator.FieldLevel) bool {
		return validatorFn(fl.Field().String())
	}
}

// Create custom uint validator function
func NewUintValidator(validatorFn func(uint) bool) CustomValidatorFn {
	return func(fl validator.FieldLevel) bool {
		return validatorFn(uint(fl.Field().Uint()))
	}
}

// Registers new custom string validator function
func RegisterStringValidator(name string, validatorFn func(string) bool) {
	validate.RegisterValidation(name, NewStringValidator(validatorFn))
}

// Registers new custom uint validator function
func RegisterUintValidator(name string, validatorFn func(uint) bool) {
	validate.RegisterValidation(name, NewUintValidator(validatorFn))
}
