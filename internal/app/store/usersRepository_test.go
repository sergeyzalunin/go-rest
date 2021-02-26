package store_test

import (
	"testing"

	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUsersRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	e := "user@mail.go"

	u, err := s.Users().Create(&models.User{
		Email: e,
	})

	assert.NoError(t, err)

	if assert.NotNil(t, u) {
		assert.Equal(t, e, u.Email)
	}
}

func TestUsersRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	e := "user@mail.go"

	_, err := s.Users().FindByEmail(e)
	assert.Error(t, err)

	u, err := s.Users().Create(&models.User{
		Email: e,
	})

	assert.NoError(t, err)

	if assert.NotNil(t, u) {
		assert.Equal(t, e, u.Email)
	}
}
