package hash

import "golang.org/x/crypto/bcrypt"

const hashCost int = 10

// Generate password hash using bcrypt
func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

// Check if password matches the hash password
func MatchPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
