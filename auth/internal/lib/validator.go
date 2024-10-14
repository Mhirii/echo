package lib

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func ValidateUsername(username string) error {
	if !govalidator.IsAlphanumeric(username) {
		return errors.New("invalid username: must be alphanumeric")
	}
	if !govalidator.MinStringLength(username, "3") {
		return errors.New("invalid username: must be longer than 3 characters")
	}
	if !govalidator.MaxStringLength(username, "30") {
		return errors.New("invalid username: must be shorter than 30 characters")
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return nil
	}
	if !govalidator.IsEmail(email) {
		return errors.New("invalid email")
	}
	return nil
}

func ValidatePassword(password string) error {
	if !govalidator.MinStringLength(password, "8") {
		return errors.New("invalid password: must be longer than 8 characters")
	}
	if !govalidator.MaxStringLength(password, "30") {
		return errors.New("invalid password: must be shorter than 30 characters")
	}
	return nil
}
