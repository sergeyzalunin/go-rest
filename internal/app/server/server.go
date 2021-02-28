package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sergeyzalunin/go-rest/internal/app/models"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
	"github.com/sirupsen/logrus"
)

const (
	sessionName            = "gosessionname"
	ctxKeyUser  contextKey = iota
	ctxKeyRequestId
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

type contextKey int8

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type server struct {
	logger        *logrus.Logger
	router        *mux.Router
	store         store.Storer
	sessionsStore sessions.Store
}

func newServer(store store.Storer, sessionsStore sessions.Store) *server {
	s := &server{
		logger:        logrus.New(),
		router:        mux.NewRouter(),
		store:         store,
		sessionsStore: sessionsStore,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/ping", s.handlePing()).Methods(http.MethodGet)
	s.router.HandleFunc("/users", s.handleCreateUser()).Methods(http.MethodPost)
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods(http.MethodPost)

	// /private/***
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoAmI()).Methods(http.MethodGet)
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		rw.Header().Set("X-Request-ID", id)
		next.ServeHTTP(rw, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestId, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestId),
		})

		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rwnew := &responseWriter{rw, http.StatusOK}
		next.ServeHTTP(rwnew, r)

		logger.Infof("completed with %d %s in %v",
			rwnew.Code,
			http.StatusText(rwnew.Code),
			time.Since(start),
		)
	})
}

func (s *server) handleWhoAmI() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.respond(rw, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*models.User))
	}
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)

			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(rw, r, http.StatusUnauthorized, errNotAuthenticated)

			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(rw, r, http.StatusUnauthorized, errNotAuthenticated)

			return
		}

		next.ServeHTTP(rw, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) handlePing() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.respond(rw, r, http.StatusOK, "pong")
	}
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

		session, err := s.sessionsStore.Get(r, sessionName)
		if err != nil {
			s.logger.Error(err)
			s.error(rw, r, http.StatusInternalServerError, err)

			return
		}

		session.Values["user_id"] = u.ID

		if err := s.sessionsStore.Save(r, rw, session); err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)

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
