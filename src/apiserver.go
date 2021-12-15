package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type ApiServer struct {
	logger *logrus.Logger
	router *mux.Router
}

func NewServer() *ApiServer {
	return &ApiServer{
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {
	err := s.configureLogger()
	if err != nil {
		return fmt.Errorf("can't configure logger: %s", err)
	}

	bindAddr := os.Getenv("PORT")
	if bindAddr == "" {
		bindAddr = "8080"
	}

	s.configureRouting()

	s.logger.Info("starting server")
	return http.ListenAndServe(":"+bindAddr, s.router)
}

func (s *ApiServer) configureLogger() error {
	Loglevel := os.Getenv("LOGLEVEL")
	if Loglevel == "" {
		Loglevel = "debug"
	}

	level, err := logrus.ParseLevel(Loglevel)
	if err != nil {
		return fmt.Errorf("can't parse level logger for string %s : %s", Loglevel, err)
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) configureRouting() {
	s.router.HandleFunc("/", s.handleHome())
}

func (s *ApiServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to homePage!")
	}
}
