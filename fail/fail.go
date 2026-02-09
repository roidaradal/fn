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
