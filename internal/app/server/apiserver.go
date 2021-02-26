package server

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *APIServer) Start() error {
	if err := a.configureLogger(); err != nil {
		return errors.Wrap(err, "could not start the server")
	}

	a.configRouter()
	a.logger.Info("starting server http://localhost:8080 ...")

	return http.ListenAndServe(a.config.BindAddr, a.router)
}

func (a *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return errors.Wrap(err, "could not parse log level")
	}

	a.logger.SetOutput(os.Stdout)
	a.logger.SetLevel(level)

	return nil
}

func (a *APIServer) configRouter() {
	a.router.HandleFunc("/hello", a.handleFunc()).Methods("GET")
}

func (a *APIServer) handleFunc() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		a.logger.Info("info")

		_, err := rw.Write([]byte("hello"))
		if err != nil {
			a.logger.Error(err)
		}
	}
}
