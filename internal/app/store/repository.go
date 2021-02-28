package store

import "github.com/sergeyzalunin/go-rest/internal/app/models"

type UserRepoHandler interface {
	Create(*models.User) error
	Find(int) (*models.User, error)
	FindByEmail(string) (*models.User, error)
}
