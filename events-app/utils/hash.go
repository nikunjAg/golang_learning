package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ComparePassword(hashed_password, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))

	return err == nil
}
