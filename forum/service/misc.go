package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func CheckPassword(plain, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plain))
}
