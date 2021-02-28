package teststore_test

import (
	"testing"

	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
	"github.com/sergeyzalunin/go-rest/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUsersRepository_Create(t *testing.T) {
	t.Parallel()

	s := teststore.New()

	tu := models.TestUser(t)

	assert.NoError(t, s.User().Create(tu))
	assert.NotNil(t, tu)
}

func TestUsersRepository_FindByEmail(t *testing.T) {
	t.Parallel()

	e := "user@mail.go"
	s := teststore.New()

	_, err := s.User().FindByEmail(e)
	assert.EqualError(t, err, store.ErrItemNotFound.Error())

	tu := models.TestUser(t)
	tu.Email = e

	assert.NoError(t, s.User().Create(tu))

	u2, err := s.User().FindByEmail(tu.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUsersRepository_Find(t *testing.T) {
	t.Parallel()

	s := teststore.New()

	_, err := s.User().Find(100)
	assert.EqualError(t, err, store.ErrItemNotFound.Error())

	tu := models.TestUser(t)
	tu.Email = "user@mail.go"

	assert.NoError(t, s.User().Create(tu))

	u2, err := s.User().Find(tu.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
