package sqlstore

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (ur *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return errors.Wrap(err, "could not validate an user")
	}

	if err := u.BeforeCreate(); err != nil {
		return errors.Wrap(err, "could not add an user into db")
	}

	if err := ur.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return errors.Wrap(err, "couldn't insert an user")
	}

	return nil
}

func (ur *UserRepository) Find(id int) (*models.User, error) {
	var u models.User

	if err := ur.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id=$1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrItemNotFound
		}

		return nil, errors.Wrap(err, "something went wrong when getting an user by specified id")
	}

	return &u, nil
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User

	if err := ur.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email=$1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrItemNotFound
		}

		return nil, errors.Wrap(err, "something went wrong when getting an user by specified email - "+email)
	}

	return &u, nil
}
