package store

import (
	"database/sql"

	_ "github.com/lib/pq" //...
	"github.com/pkg/errors"
)

type Store struct {
	config   *Config
	db       *sql.DB
	userRepo *UsersRepository
}

func New(config *Config) *Store {
	return &Store{config: config}
}

func (s *Store) Open() error {
	var err error

	s.db, err = sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return errors.Wrap(err, "couldn't get handler to "+s.config.DatabaseURL)
	}

	if err = s.db.Ping(); err != nil {
		return errors.Wrap(err, "couldn't ping database "+s.config.DatabaseURL)
	}

	return nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Users() *UsersRepository {
	if s.userRepo == nil {
		s.userRepo = &UsersRepository{s}
	}

	return s.userRepo
}
