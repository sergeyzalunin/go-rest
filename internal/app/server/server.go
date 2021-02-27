package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
	"github.com/sirupsen/logrus"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type server struct {
	logger *logrus.Logger
	router *mux.Router
	store  store.Storer
}

func newServer(store store.Storer) *server {
	s := &server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleCreateUser()).Methods(http.MethodPost)
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods(http.MethodPost)
}

func (s *server) handleCreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var req request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(rw, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(rw, r, http.StatusUnprocessableEntity, err)

			return
		}

		u.Sanitize()

		s.respond(rw, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var req request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(rw, r, http.StatusBadRequest, err)

			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil {
			s.logger.Error(err)
			s.error(rw, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)

			return
		}

		if !u.ComparePassword(req.Password) {
			s.error(rw, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)

			return
		}

		s.respond(rw, r, http.StatusOK, nil)
	}
}

func (s *server) error(rw http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(rw, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(rw http.ResponseWriter, _ *http.Request, code int, data interface{}) {
	rw.WriteHeader(code)

	if data != nil {
		if err := json.NewEncoder(rw).Encode(data); err != nil {
			s.logger.Error(err)
		}
	}
}
