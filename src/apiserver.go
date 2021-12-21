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
	store  *Store
}

func NewServer() *ApiServer {
	return &ApiServer{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  NewStore(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return fmt.Errorf("can't configure logger: %s", err)
	}

	bindAddr := os.Getenv("PORT")
	if bindAddr == "" {
		bindAddr = "8080"
	}

	s.configureRouting()

	if err := s.configureStore(); err != nil {
		return fmt.Errorf("can't configure store: %s", err)
	}
	defer s.store.Close()

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
	s.router.HandleFunc("/registration", s.handleRegistration())
}

func (s *ApiServer) configureStore() error {
	if err := s.store.Open(); err != nil {
		fmt.Errorf("can't open store: %s", err)
	}
	return nil
}

func (s *ApiServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, user := s.store.all()
		fmt.Fprintf(w, "Welcome to homePage!"+user)
	}
}

func (s *ApiServer) handleRegistration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
