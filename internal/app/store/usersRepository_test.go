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

	tu := models.TestUser(t)
	u, err := s.Users().Create(tu)

	assert.NoError(t, err)

	if assert.NotNil(t, u) {
		assert.Equal(t, tu.Email, u.Email)
	}
}

func TestUsersRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	e := "user@mail.go"

	_, err := s.Users().FindByEmail(e)
	assert.Error(t, err)

	tu := models.TestUser(t)
	tu.Email = e

	u, err := s.Users().Create(tu)

	assert.NoError(t, err)

	if assert.NotNil(t, u) {
		assert.Equal(t, e, u.Email)
	}
}
