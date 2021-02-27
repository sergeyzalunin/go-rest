package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" //...
	"github.com/sergeyzalunin/go-rest/internal/app/store"
)

type Store struct {
	db       *sql.DB
	userRepo *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) User() store.UserRepoHandler {
	if s.userRepo == nil {
		s.userRepo = &UserRepository{s}
	}

	return s.userRepo
}
