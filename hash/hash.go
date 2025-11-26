// Package hash contains password hashing function using bcrypt.
package hash

import "golang.org/x/crypto/bcrypt"

const hashCost int = 10

// Generate password hash using bcrypt
func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

// Check if password matches the hash password
func MatchPassword(rawPassword, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))
	return err == nil
}
