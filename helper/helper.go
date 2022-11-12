package helper

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var ErrUnauthorized = errors.New("User unauthorized")

func GeneratePasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return string(passwordHash), err
	}

	return string(passwordHash), nil
}

func ComparePasswordHash(passwordHashed, passwrod string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(passwrod))
	if err != nil {
		return err
	}

	return nil
}

func FormatValidationError(err error) []string {
	errors := make([]string, 0)

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
