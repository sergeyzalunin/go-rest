package sqlstore_test

import (
	"testing"

	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
	"github.com/sergeyzalunin/go-rest/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUsersRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	tu := models.TestUser(t)

	assert.NoError(t, s.User().Create(tu))
	assert.NotNil(t, tu)
}

func TestUsersRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	e := "user@mail.go"
	s := sqlstore.New(db)

	_, err := s.User().FindByEmail(e)
	assert.EqualValues(t, err, store.ErrItemNotFound)

	tu := models.TestUser(t)
	tu.Email = e

	assert.NoError(t, s.User().Create(tu))
	assert.NotNil(t, tu)
}
