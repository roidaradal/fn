// Package fail contains common error messages
package fail

import "fmt"

// FromErrors produces an error message from list of errors
func FromErrors(label string, errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return fmt.Errorf("%s: %d errors encountered: %w", label, len(errs), errs[0])
}

// Initialization produces an error message for failing initialization process
func Initialization(label string, err error) error {
	return fmt.Errorf("%s: failed to initialize: %w", label, err)
}
