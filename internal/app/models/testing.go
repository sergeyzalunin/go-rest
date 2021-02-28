package models

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		ID:                100,
		Email:             "user@example.org",
		Password:          "password",
		EncryptedPassword: "",
	}
}