package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// @TODO : make function hash password

// UnsafeHashPassword hashes the password without checking for an error
func UnsafeHashPassword(pw string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pw), 12)
	return string(hash)
}

// UnsafeHashPassword hashes the password
func HashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ValidatePassword checks whether the password matches the hashed password
func ValidatePassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
