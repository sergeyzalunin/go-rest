package models

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	id := 100

	return &User{
		ID:                id,
		Email:             "user@example.org",
		Password:          "password",
		EncryptedPassword: "",
	}
}
