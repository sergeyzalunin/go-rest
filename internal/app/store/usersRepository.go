package store

import (
	"github.com/pkg/errors"
	"github.com/sergeyzalunin/go-rest/internal/app/models"
)

type UsersRepository struct {
	store *Store
}

func (ur *UsersRepository) Create(u *models.User) (*models.User, error) {
	if err := ur.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, errors.Wrap(err, "couldn't insert an user")
	}

	return u, nil
}

func (ur *UsersRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User

	if err := ur.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email=$1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, errors.Wrap(err, "couldn't find an user by email")
	}

	return &u, nil
}
