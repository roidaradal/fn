// Package fail contains common error messages
package fail

import (
	"fmt"
	"strings"
)

// FromErrors produces an error message from list of errors
func FromErrors(label string, errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return fmt.Errorf("%s: %d errors encountered: %w", label, len(errs), errs[0])
}

// EnvLoad produces an error message from loading the env file
func LoadEnv(err error) error {
	return fmt.Errorf("envLoadFail: %w", err)
}

// PublicMessage gets the public error message from "public: <message>" error
func PublicMessage(err error) (string, bool) {
	msg := err.Error()
	if strings.HasPrefix(msg, "public:") {
		parts := strings.Split(msg, ":")
		return strings.TrimSpace(parts[1]), true
	}
	return "Unexpected error occured", false
}
