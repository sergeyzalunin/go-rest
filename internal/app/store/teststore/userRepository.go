package teststore

import (
	"github.com/pkg/errors"
	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
)

type UserTestRepository struct {
	store *Store
	users map[string]*models.User
}

func (ur *UserTestRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return errors.Wrap(err, "could not validate an user")
	}

	if err := u.BeforeCreate(); err != nil {
		return errors.Wrap(err, "could not add an user into db")
	}

	ur.users[u.Email] = u

	return nil
}

func (ur *UserTestRepository) FindByEmail(email string) (*models.User, error) {
	u, ok := ur.users[email]

	if !ok {
		return nil, store.ErrItemNotFound
	}

	return u, nil
}
