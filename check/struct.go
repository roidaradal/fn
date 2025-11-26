package check

import "github.com/go-playground/validator/v10"

var validate = validator.New(validator.WithRequiredStructEnabled())

// Validate struct using Go struct validator
func IsValidStruct[T any](item *T) bool {
	if item == nil {
		return false
	}
	err := validate.Struct(item)
	return err == nil
}

// Check if struct is invalid, using Go struct validator
func NotValidStruct[T any](item *T) bool {
	return !IsValidStruct(item)
}

// Registers new custom string field validator
func RegisterStringValidator(name string, stringValidator func(string) bool) {
	validatorFn := func(fl validator.FieldLevel) bool {
		return stringValidator(fl.Field().String())
	}
	validate.RegisterValidation(name, validatorFn)
}

// Registers new custom uint field validator
func RegisterUintValidator(name string, uintValidator func(uint) bool) {
	validatorFn := func(fl validator.FieldLevel) bool {
		return uintValidator(uint(fl.Field().Uint()))
	}
	validate.RegisterValidation(name, validatorFn)
}
