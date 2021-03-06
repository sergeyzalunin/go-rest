package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) Validate() error {
	minLen, maxLen := 6, 100
	passRule := requiredIf(u.EncryptedPassword == "")

	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(passRule), validation.Length(minLen, maxLen)),
	)
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
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
