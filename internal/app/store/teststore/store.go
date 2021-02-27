package teststore

import (
	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
)

type Store struct {
	userRepo *UserTestRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepoHandler {
	if s.userRepo == nil {
		s.userRepo = &UserTestRepository{
			store: s,
			users: make(map[string]*models.User),
		}
	}

	return s.userRepo
}
