package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sergeyzalunin/go-rest/internal/app/store"
	"github.com/sirupsen/logrus"
)

type server struct {
	logger *logrus.Logger
	router *mux.Router
	store  store.Storer
}

func newServer(store store.Storer) *server {
	s := &server{
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleFunc()).Methods("GET")
}

func (s *server) handleFunc() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.logger.Info("info")

		_, err := rw.Write([]byte("hello"))
		if err != nil {
			s.logger.Error(err)
		}
	}
}
