package teststore

import (
	"testing"

	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUsersRepository_Create(t *testing.T) {
	s := New()

	tu := models.TestUser(t)

	assert.NoError(t, s.User().Create(tu))
	assert.NotNil(t, tu)
}

func TestUsersRepository_FindByEmail(t *testing.T) {
	e := "user@mail.go"
	s := New()

	_, err := s.User().FindByEmail(e)
	assert.EqualError(t, err, store.ErrItemNotFound.Error())

	tu := models.TestUser(t)
	tu.Email = e

	assert.NoError(t, s.User().Create(tu))
	assert.NotNil(t, tu)
}
