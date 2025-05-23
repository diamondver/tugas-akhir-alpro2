package helper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword creates a cryptographically secure hash of the provided password using bcrypt.
// The bcrypt algorithm automatically incorporates a random salt during the hashing process.
//
// Parameters:
//   - password: The plain text password to be hashed
//
// Returns:
//   - string: The bcrypt hashed password (includes the salt)
//   - error: Any error encountered during the hashing process
//
// Note: There's a small bug in the return statement - it should return nil for the error on success.
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to generate hash from password: %w", err)
	}

	return string(hashedBytes), err
}

// CheckPasswordHash verifies if a plain text password matches a previously hashed password.
// It uses bcrypt's comparison function which is resistant to timing attacks.
//
// Parameters:
//   - password: The plain text password to check
//   - hash: The bcrypt hashed password to compare against
//
// Returns:
//   - bool: True if the password matches the hash, false otherwise
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}

	return true
}
