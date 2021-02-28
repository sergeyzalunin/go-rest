package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"github.com/sergeyzalunin/go-rest/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	srv.logger.Infof("starting server at http://localhost%s ...", config.BindAddr)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "could not establish connection to db"+dbURL)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "could not ping database")
	}

	return db, nil
}
