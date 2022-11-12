package storage

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash_Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func check_password(password, e_password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e_password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
