package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	generatedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(generatedPassword), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
