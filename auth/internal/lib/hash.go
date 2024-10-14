package lib

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(secret string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(secret), 8)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func ComparePassword(hashed string, plain string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err
}
