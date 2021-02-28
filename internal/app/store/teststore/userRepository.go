package teststore

import (
	"github.com/pkg/errors"
	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
)

type UserTestRepository struct {
	store *Store
	users map[int]*models.User
}

func (ur *UserTestRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return errors.Wrap(err, "could not validate an user")
	}

	if err := u.BeforeCreate(); err != nil {
		return errors.Wrap(err, "could not add an user into db")
	}

	u.ID = len(ur.users) + 1
	ur.users[u.ID] = u

	return nil
}

func (ur *UserTestRepository) Find(id int) (*models.User, error) {
	u, ok := ur.users[id]

	if !ok {
		return nil, store.ErrItemNotFound
	}

	return u, nil
}

func (ur *UserTestRepository) FindByEmail(email string) (*models.User, error) {
	for _, u := range ur.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrItemNotFound
}
