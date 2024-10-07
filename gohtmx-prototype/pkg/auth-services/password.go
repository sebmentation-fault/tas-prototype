package authservices

import (
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

// Generate the hash from the password
func GenerateHashFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

// Compare the hashed password and a newly inputted password
func ComparePasswordWithHash(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	// theres a problem with the checking --> maybe password was supposed to match
	// but somehow an error is being thrown
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		slog.Warn("Could not compare the passwords: %s", err)
	}

	return err == nil
}
