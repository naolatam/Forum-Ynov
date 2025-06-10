package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func IsSecurePassword(password string) bool {

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case char >= '0' && char <= '9':
			hasDigit = true
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '!' && char <= '/':
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial && len(password) >= 8
}

func CheckForNewPassword(password, confirmPassword string) (string, error) {
	if confirmPassword != password {
		return "", errors.New("passwords do not match")
	}
	if len(password) < 8 {
		return "", errors.New("password must be at least 8 characters long")
	}
	if !IsSecurePassword(password) {
		return "", errors.New("password must contain at least one number, one uppercase letter, one lowercase letter, and one special character")
	}

	passwordBytes := []byte(password)
	pass, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to compute password")
	}
	return string(pass), nil

}
