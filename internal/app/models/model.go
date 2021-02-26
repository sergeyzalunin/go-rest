package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func (u *User) Validate() error {
	passRule := requiredIf(u.EncryptedPassword == "")

	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(passRule), validation.Length(6, 100)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptedString(u.Password)
		if err != nil {
			return errors.Wrap(err, "could not encrypt password")
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func encryptedString(s string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(s), 11)

	return string(res), errors.Wrap(err, "something went wrong on encrypting password")
}
